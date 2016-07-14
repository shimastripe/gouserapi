package models

import (
	"reflect"
	"strings"
)

func contains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

func FieldToMap(model interface{}, fields string) map[string]interface{} {
	u := make(map[string]interface{})
	ts, vs := reflect.TypeOf(model), reflect.ValueOf(model)

	for i := 0; i < ts.NumField(); i++ {
		var jsonKey string

		field := ts.Field(i)
		jsonTag := field.Tag.Get("json")

		if jsonTag == "" {
			jsonKey = field.Name
		} else {
			jsonKey = strings.Split(jsonTag, ",")[0]
		}

		if fields == "*" || contains(strings.Split(fields, ","), jsonKey) {
			u[jsonKey] = vs.Field(i).Interface()
		}
	}
	return u
}
