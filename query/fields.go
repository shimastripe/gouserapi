package query

import (
	"reflect"

	"github.com/serenize/snaker"
)

func FilterField(fields []string, model interface{}) map[string]interface{} {
	filterdata := make(map[string]interface{})
	modelmap := StructToMap(model)
	for _, v := range fields {
		filterdata[snaker.CamelToSnake(v)] = modelmap[snaker.SnakeToCamel(v)]
	}
	return filterdata
}

func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		result[field] = value
	}
	return result
}
