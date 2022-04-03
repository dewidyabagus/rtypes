package rtypes

import (
	"fmt"
	"reflect"
	"strings"
)

func ConvertStructToMapInterface(data interface{}) (mapInterface map[string]interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			if strings.Contains(r.(string), "NumField of non-struct type") {
				err = fmt.Errorf("data type cannot be processed")
			} else {
				err = fmt.Errorf(r.(string))
			}
		}
	}()

	response := map[string]interface{}{}
	if data == nil {
		return response, nil
	}

	dataTypeOf := reflect.TypeOf(data)
	dataValueOf := reflect.ValueOf(data)

	if dataTypeOf.Kind() == reflect.Ptr {
		dataTypeOf = dataTypeOf.Elem()
		dataValueOf = reflect.Indirect(dataValueOf)
	}

	for i := 0; i < dataTypeOf.NumField(); i++ {
		tag := dataTypeOf.Field(i).Tag.Get("map")

		if tag != "" && tag != "-" {
			valueInterface := dataValueOf.Field(i).Interface()

			if dataTypeOf.Field(i).Type.Kind() == reflect.Struct || dataTypeOf.Field(i).Type.Kind() == reflect.Ptr {
				response[tag], err = ConvertStructToMapInterface(valueInterface)
				if err != nil {
					return nil, err
				}

			} else {
				response[tag] = valueInterface
			}
		}
	}

	return response, nil
}
