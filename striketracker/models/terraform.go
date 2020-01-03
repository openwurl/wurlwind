package models

import (
	"reflect"

	"github.com/openwurl/wurlwind/pkg/debug"
)

const terraformTag = "tf"

// MapFromStruct extracts a map[string]interface from a HW API model if it uses the tf struct tag
// should only be used on structs that contain generics
func MapFromStruct(s interface{}) map[string]interface{} {
	if s != nil {
		ret := make(map[string]interface{})

		reflection := reflect.ValueOf(s).Elem()

		for i := 0; i < reflection.NumField(); i++ {
			thisField := reflection.Field(i)
			thisType := reflection.Type().Field(i)
			tag := thisType.Tag

			if val, ok := tag.Lookup(terraformTag); ok {
				// Dereference pointers within the struct to their types
				var thisFieldDeref reflect.Value
				if thisField.Kind() == reflect.Ptr {
					thisFieldDeref = thisField.Elem()
				} else {
					thisFieldDeref = thisField
				}

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

		for i := 0; i < rv.NumField(); i++ {
			thisField := rv.Field(i)
			thisType := rv.Type().Field(i)
			tag := thisType.Tag

			if val, ok := tag.Lookup(terraformTag); ok {

				// Dereference pointers within the struct to their types
				var thisFieldDeref reflect.Value
				if thisField.Kind() == reflect.Ptr {
					thisField.Set(reflect.New(thisField.Type().Elem()))
					thisFieldDeref = thisField.Elem()
				} else {
					thisFieldDeref = thisField
				}

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
	return nil

}
