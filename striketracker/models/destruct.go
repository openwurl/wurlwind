package models

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/openwurl/wurlwind/pkg/debug"
)

const (
	terraformTag  = "tf"
	parentTag     = "parent"
	nameTag       = "name"
	modTag        = "modify"
	weightedValue = "weighted"
)

/*
	This work is tailor made to this implementation. Since we know our top-level
	struct only contains pointers to other structs and slices of structs, it is likely
	safe to use. However if the struct were to deviate and contain primatives at the top
	level this would fall apart (and other available methods here are more suited)
*/

// ExtractSchema extracts data from *Configuration
// and returns a schema.Schema ingestible payload
func (c *Configuration) ExtractSchema() map[string]interface{} {
	p := &Parents{
		Queue: make([]*Parent, 0),
	}

	reflection := reflect.ValueOf(*c)

	// Iterate configuration fields
	for i := 0; i < reflection.NumField(); i++ {
		r := &Reflected{
			Reflection: reflection,
		}
		r.Configure(i)

		// check for parent tag
		if parent, pok := r.Tag.Lookup(parentTag); pok {
			// check that it is named
			if name, nok := r.Tag.Lookup(nameTag); nok {
				debug.Log("ExtractSchema", "=========== %v.%v\n", parent, name)
				debug.Log("ExtractSchema", "%v.%v - processing...\n", parent, name)

				if !r.Valid() {
					debug.Log("ExtractSchema", "%v.%v - Invalid or empty, skipping...\n", parent, name)
					continue
				}

				p.AddParent(parent)

				switch r.Field.Kind() {
				case reflect.Slice:
					debug.Log("ExtractSchema", "%v.%v - slice\n", parent, name)
					if mod, mok := r.Tag.Lookup(modTag); mok {
						if mod == weightedValue {
							debug.Log("Modifiers", "%v.%v - modifier = %v\n", parent, name, mod)
							thisParent := p.Get(parent)
							f := thisParent.Add(name)
							f.Content = UnpackWeightedReflectedSlice(r.Field)
							continue
						}
					}
					debug.Log("ExtractSchema", "%v.%v - slice untagged, skipping...\n", parent, name)
				case reflect.Ptr:
					debug.Log("ExtractSchema", "%v.%v - pointer\n", parent, name)
					thisParent := p.Get(parent)
					debug.Log("Parent adding", "name: %s | parent: %v\n", name, spew.Sprintf("%v", thisParent))
					f := thisParent.Add(name)
					f.Content = MapFromReflectValue(r.Field.Elem())
				default:
					debug.Log("ExtractSchema", "%v.%v - Ended up in broken default state, skipping...\n", parent, name)
					continue
				}

			}
		}

	}
	return p.Dump()
}

// IngestSchema ingests a terraform *schema.Schema into the top level config struct and it's subs
func (c *Configuration) IngestSchema(schemaSlice map[string]interface{}) error {
	//schemaSlice := schema.([]interface{})[0].(map[string]interface{})

	//dReflection := reflect.ValueOf(*c)
	reflection := reflect.ValueOf(c)

	// For each field in *Configuration
	for i := 0; i < reflection.Elem().NumField(); i++ {
		r := &Reflected{
			Reflection: reflection.Elem(),
		}
		r.Configure(i)

		if name, ok := r.Tag.Lookup(nameTag); ok {
			if !r.Field.CanAddr() {
				debug.Log("%v is not valid??\n", name)
				continue
			}
			for k, v := range schemaSlice {

				if k == name {

					debug.Log("ITER LOOP", "%s | %v | %v | %v", name, reflect.ValueOf(v).Type(), reflect.ValueOf(v).Kind(), v)

					switch r.Field.Kind() {
					case reflect.Ptr:
						if IsSchemaSet(v) {
							v = DesetSchemaMap(v)
						}
						// Is a pointer to a struct
						StructFromValue(&r.Field, v.(map[string]interface{}))
					case reflect.Slice:
						if IsSchemaSet(v) {
							v = DesetSchemaSlice(v)
						}
						// Is a slice of pointers
						SliceFromValue(&r.Field, r.Type, v.([]interface{}))
					default:
						debug.Log("%v - not implemented\n", name)
					}

				}

			}

		}

	}

	return nil
}

/*
	Helpers
*/

// expandSetOfMaps expands a tf set of maps into its first level map
func expandSetOfMaps(raw interface{}) map[string]interface{} {
	if deliverySet, ok := raw.(*schema.Set); ok {
		set := deliverySet.List()[0]
		if deliverySlice, ok := set.(map[string]interface{}); ok {
			return deliverySlice
		}
	}
	return nil
}

// MapFromReflectValue returns a map[string]interface
// representing the underlying struct
func MapFromReflectValue(v reflect.Value) map[string]interface{} {
	m := make(map[string]interface{})

	deref := v
	if v.Kind() == reflect.Ptr {
		deref = v.Elem()
	}

	for i := 0; i < deref.NumField(); i++ {
		thisField := deref.Field(i)
		thisType := deref.Type().Field(i)
		tag := thisType.Tag

		if val, ok := tag.Lookup(terraformTag); ok {
			if thisField.Kind() == reflect.Ptr {
				m[val] = thisField.Elem().Interface()
			} else {
				m[val] = thisField.Interface()
			}
		}
	}

	return m
}

// UnpackWeightedReflectedSlice iterates a reflected slice and unpacks into an []interface{} of maps
func UnpackWeightedReflectedSlice(s reflect.Value) []interface{} {
	ret := make([]interface{}, 0)
	for i := 0; i < s.Len(); i++ {
		m := MapFromReflectValue(s.Index(i))
		m["weight"] = i
		ret = append(ret, m)
	}
	return ret
}

// StructFromValue attempst to return a given struct inside a reflect.Value packed
// with the given map
func StructFromValue(model *reflect.Value, m map[string]interface{}) *reflect.Value {
	if m != nil {
		// Dereference the model
		model.Set(reflect.New(model.Type().Elem()))
		deref := model.Elem()

		// iterate fields
		for i := 0; i < deref.NumField(); i++ {
			thisField := deref.Field(i)
			thisType := deref.Type().Field(i)
			tag := thisType.Tag

			// check for tf tag
			if val, ok := tag.Lookup(terraformTag); ok {

				// dereference field
				var thisFieldDeref reflect.Value
				if thisField.Kind() == reflect.Ptr {
					thisField.Set(reflect.New(thisField.Type().Elem()))
					thisFieldDeref = thisField.Elem()
				} else {
					thisFieldDeref = thisField
				}

				// detect the type and cast our map value to that type in that field
				switch thisFieldDeref.Kind() {
				case reflect.Int:
					if v, ok := m[val]; ok {
						thisFieldDeref.SetInt(int64(v.(int)))
					}
				case reflect.String:
					if v, ok := m[val]; ok {
						thisFieldDeref.SetString(v.(string))
					}
				case reflect.Bool:
					if v, ok := m[val]; ok {
						thisFieldDeref.SetBool(v.(bool))
					}
				default:
					if m == nil {
						debug.Log("StructFromValue", "Something went wrong, m is nil: %v", spew.Sprintf("%v", thisFieldDeref.Kind()))
					} else {
						debug.Log("StructFromValue", "Something went wrong packing: %v\n", fmt.Sprintf("%v", m))
					}

					return nil
				}

			}

		}

	}
	return model
}

// SliceFromValue attempts to return a packed slice of structs from a given reflect.Value
// and slice of maps
func SliceFromValue(model *reflect.Value, modelType reflect.StructField, m []interface{}) *reflect.Value {

	// iterate the slice of maps and append them into the model
	for _, sliceItem := range m {
		item := sliceItem.(map[string]interface{})
		r := reflect.New(model.Type().Elem()).Elem()
		StructFromValue(&r, item)
		model.Set(reflect.Append(*model, r))
	}

	return nil
}

// StructFromMap attempts to return a given struct packed with the given
// map and should only be used on structs full of generics
func StructFromMap(model interface{}, m map[string]interface{}) interface{} {
	if m != nil {

		rv := reflect.ValueOf(model)

		// dereference if applicable
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}

		for i := 0; i < rv.NumField(); i++ {
			r := Reflected{
				Reflection: rv,
			}
			r.Configure(i)

			if name, ok := r.Tag.Lookup(terraformTag); ok {
				if !r.Field.CanAddr() {
					debug.Log("StructFromMap", "%v is not valid??\n", name)
					continue
				}

				debug.Log("REFLECTION ITERATION", "[%d] %s (%s)", i, name, r.Type.Type.String())

				// dereference field
				var thisFieldDeref reflect.Value
				if r.Field.Kind() == reflect.Ptr {
					debug.Log("StructFromMap", "Dereferencing: %v [%v]", name, r.Type.Type.String())
					r.Field.Set(reflect.New(r.Field.Type().Elem()))
					thisFieldDeref = r.Field.Elem()
				} else {
					thisFieldDeref = r.Field
				}

				debug.Log("StructFromMap", "Dereferenced: %v: %v => %v", name, r.Type.Type.String(), spew.Sprintf("%v", thisFieldDeref.Type().String()))

				for k, v := range m {
					if k == name {

						switch thisFieldDeref.Kind() {
						case reflect.Int:
							//r.Field.SetInt(int64(v.(int)))
							thisFieldDeref.SetInt(int64(v.(int)))
						case reflect.String:
							//r.Field.SetString(v.(string))
							thisFieldDeref.SetString(v.(string))
						case reflect.Bool:
							//r.Field.SetBool(v.(bool))
							thisFieldDeref.SetBool(v.(bool))
						default:
							debug.Log("StructFromMap", "Something went wrong packing: %v - %v [%v]\n", name, fmt.Sprintf("%v", v), thisFieldDeref.Type().String())
							return nil
						}

					}
				}
			}
		}
		return model

	}
	return nil
}

// StructFromMapOld attempts to return a given struct packed with the given map
// should only be used on structs that contain generics int / bool /string
func StructFromMapOld(model interface{}, m map[string]interface{}) interface{} {
	if m != nil {

		//model.Set(reflect.New(model.Type().Elem()))
		//deref := model.Elem()

		// NEED TO FIX THIS
		rv := reflect.ValueOf(model)

		spew.Dump(rv)
		//		rv := reflect.ValueOf(&model)
		//		rv.Set(reflect.New(rv.Type().Elem()))
		//
		//		if rv.Kind() == reflect.Ptr {
		//			rv = rv.Elem()
		//		} else {
		//			debug.Log("STRUCT GENERATION", "%v is not a pointer", rv.Kind().String())
		//			return nil
		//		}

		debug.Log("BEFORE ITER", "%v", spew.Sprintf("%v", rv))

		debug.Log("BREAKDOWN", "Fields: [%d]", rv.NumField())

		// iterate fields in our interface/struct
		for i := 0; i < rv.NumField(); i++ {
			thisField := rv.Field(i)
			thisType := rv.Type().Field(i)
			tag := thisType.Tag

			debug.Log("FIELD", "[%d] - [%s]", i, thisType.Type.String())

			// check to make sure it is a tagged field
			if val, ok := tag.Lookup(terraformTag); ok {

				debug.Log("BEFORE DEREF", "%v", spew.Sprintf("%v", thisField))

				// Dereference pointers within the struct to their types
				var thisFieldDeref reflect.Value
				if thisField.Kind() == reflect.Ptr {
					debug.Log("DEREFERENCING", "[%s] - %s", thisType.Type.String(), val)
					thisField.Set(reflect.New(thisField.Type().Elem()))
					thisFieldDeref = thisField.Elem()
				} else {
					debug.Log("!!NOT!! DEREFERENCING", "[%s] - %s", thisType.Type.String(), val)
					thisFieldDeref = thisField
				}

				debug.Log("AFTER DEREF", "%v", spew.Sprintf("%v", thisFieldDeref))

				// detect the type and cast our map value to that type in that field
				switch thisFieldDeref.Kind() {
				case reflect.Int:
					if v, ok := m[val]; ok {
						thisFieldDeref.SetInt(int64(v.(int)))
					}
				case reflect.String:
					if v, ok := m[val]; ok {
						thisFieldDeref.SetString(v.(string))
					}
				case reflect.Bool:
					if v, ok := m[val]; ok {
						thisFieldDeref.SetBool(v.(bool))
					}
				default:
					debug.Log("StructFromMap", "Something went wrong packing: %v\n", m)
					return nil
				}
			}
		}
		return model
	}
	// We don't want to make empty things where there is nothing
	return nil
}

// MapFromStruct extracts a map[string]interface from a HW API model if it uses the tf struct tag
// should only be used on structs that contain generics
func MapFromStruct(s interface{}) map[string]interface{} {
	// make sure our interface isn't empty
	if s != nil {
		ret := make(map[string]interface{})

		reflection := reflect.ValueOf(s).Elem()

		// iterate the fields in the interface (struct)
		for i := 0; i < reflection.NumField(); i++ {
			thisField := reflection.Field(i)
			thisType := reflection.Type().Field(i)
			tag := thisType.Tag

			// check for our tag on this field
			if val, ok := tag.Lookup(terraformTag); ok {
				// Dereference pointers within the struct to their types
				var thisFieldDeref reflect.Value
				if thisField.Kind() == reflect.Ptr {
					thisFieldDeref = thisField.Elem()
				} else {
					thisFieldDeref = thisField
				}

				// assign map from this struct field
				ret[val] = thisFieldDeref.Interface()
			}
		}
		return ret
	}
	// We don't want to make empty things where there is nothing
	return nil
}

// IsSchemaSet returns true of the given interface{} is a *schema.Set
func IsSchemaSet(val interface{}) bool {
	trv := reflect.ValueOf(val)
	var ref *schema.Set
	if trv.Type() == reflect.ValueOf(ref).Type() {
		return true
	}
	return false
}

// DesetSchemaMap returns a de-setted schema in interface format
func DesetSchemaMap(val interface{}) map[string]interface{} {
	if deliverySet, ok := val.(*schema.Set); ok {
		if deliverySet.Len() > 0 {
			set := deliverySet.List()[0]
			if deliverySlice, ok := set.(map[string]interface{}); ok {
				return deliverySlice
			}
		} else {
			debug.Log("Schema Parsing", "SCHEMA SET HAS NO LENGTH - %v", val)
		}
	}
	return nil
}

// DesetSchemaSlice returns a de-setted schema in interface format
func DesetSchemaSlice(val interface{}) []interface{} {
	if deliverySet, ok := val.(*schema.Set); ok {
		set := deliverySet.List()[0]
		if deliverySlice, ok := set.([]interface{}); ok {
			return deliverySlice
		}
	}
	return nil
}
