package models

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/openwurl/wurlwind/pkg/debug"
)

const (
	terraformTag  = "tf"
	parentTag     = "parent"
	nameTag       = "name"
	modTag        = "modify"
	weightedValue = "weighted"
)

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

// StructFromMap attempts to return a given struct packed with the given map
// should only be used on structs that contain generics int / bool /string
func StructFromMap(model interface{}, m map[string]interface{}) interface{} {
	if m != nil {
		rv := reflect.ValueOf(model).Elem()

		// iterate fields in our interface/struct
		for i := 0; i < rv.NumField(); i++ {
			thisField := rv.Field(i)
			thisType := rv.Type().Field(i)
			tag := thisType.Tag

			// check to make sure it is a tagged field
			if val, ok := tag.Lookup(terraformTag); ok {

				// Dereference pointers within the struct to their types
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
					debug.Log("Model Generation", "Something went wrong packing: %v\n", m)
					return nil
				}
			}
		}
		return model
	}
	// We don't want to make empty things where there is nothing
	return nil

}

// ExpandParentFields digs out all of the subfields in a map[string]interface{}
// and attemps to assign a struct
func (c *Configuration) ExpandParentFields(s map[string]interface{}) error {
	reflection := reflect.ValueOf(*c)

	for i := 0; i < reflection.NumField(); i++ {
		thisField := reflection.Field(i)
		thisType := reflection.Type().Field(i)
		tag := thisType.Tag

		// check for our packing tag and parent tag
		if name, nok := tag.Lookup(nameTag); nok {

			for k, v := range s {
				if k == name {
					switch thisField.Kind() {
					case reflect.Slice:
						debug.Log("TESTING", "[SLICE] Found match for %v: [%v]", k, v)
					case reflect.Ptr:
						debug.Log("TESTING", "[POINTER] Found match for %v: [%v]", k, v)
						//packReflectedPointer(thisField, v)
					default:
						debug.Log("TESTING", "[UNKNOWN] Found match for %v: [%v]", k, v)
					}

				}
			}

		}

	}

	return fmt.Errorf("Exiting here on purpose")

}

// CompressParentFields digs out all fields with parent tag and generates a map
func (c *Configuration) CompressParentFields() []interface{} {
	ret := make([]interface{}, 0)
	// map[parentName]map[parentTag]interface{}
	tmp := make(map[string]map[string]interface{}, 0) // holds parents and subordinates

	//var reflection reflect.Value
	reflection := reflect.ValueOf(*c) //.Elem()
	//if reflection.Kind() == reflect.Ptr {
	//	reflection = reflect.ValueOf(c).Elem()
	//}

	// Iterate all fields in configuration
	for i := 0; i < reflection.NumField(); i++ {
		thisField := reflection.Field(i)
		thisType := reflection.Type().Field(i)
		tag := thisType.Tag

		// check for our packing tag and parent tag
		if parent, pok := tag.Lookup(parentTag); pok {
			if name, nok := tag.Lookup(nameTag); nok {

				debug.Log(fmt.Sprintf("REFLECT - %s", name), "processing")

				if !thisField.IsValid() || thisField.IsNil() || thisField.IsZero() {
					debug.Log(fmt.Sprintf("REFLECT - %s", name), "empty or invalid, skipping")
					continue
				}

				if tmp[parent] == nil {
					tmp[parent] = make(map[string]interface{})
				}

				if tmp[parent][name] == nil {
					tmp[parent][name] = make([]interface{}, 0)
				}

				this := tmp[parent][name].([]interface{})

				switch thisField.Kind() {
				case reflect.Slice:
					if wname, wok := tag.Lookup(modTag); wok {
						if wname == weightedValue {
							this = UnpackWeightedReflectedSlice(thisField)
							debug.Log(fmt.Sprintf("REFLECT - %s", name), "%v", spew.Sprintf("%v", this))
						}
					} else {
						this = append(this, UnpackReflectedSlice(thisField))
					}

				case reflect.Ptr:
					this = append(this, MapFromReflectValue(thisField.Elem()))
				default:
					debug.Log(fmt.Sprintf("REFLECT - %s", name), "You ended up somewhere you shouldn't be: %v | %v", parent, name)
					continue
				}

				tmp[parent][name] = this

			}
		}
	}
	ret = append(ret, tmp)
	return ret
}

/*
	Helpers
*/

// UnpackReflectedSlice iterates a reflected slice and unpacks into an []interface{} of maps
func UnpackReflectedSlice(s reflect.Value) []interface{} {
	ret := make([]interface{}, 0)
	for i := 0; i < s.Len(); i++ {
		ret = append(ret, MapFromReflectValue(s.Index(i)))
	}
	return ret
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

// IterateSliceValue iterates a reflected slice
func IterateSliceValue(s reflect.Value) []interface{} {
	ret := make([]interface{}, 0)
	for i := 0; i < s.Len(); i++ {
		ret = append(ret, MapFromReflectValue(s.Index(i)))
	}
	return ret
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
