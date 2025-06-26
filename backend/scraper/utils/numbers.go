package utils

import "reflect"

func ConvertZeroToMax(input interface{}) {
	val := reflect.ValueOf(input).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.CanSet() {
			switch field.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if field.Int() == 0 {
					field.SetInt(999999)
				}
			case reflect.Float32, reflect.Float64:
				if field.Float() == 0.0 {
					field.SetFloat(999999)
				}
			}
		}
	}
}
