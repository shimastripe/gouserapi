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

func ParseFields(fields string) ([]string, map[string][]string) {
	fieldsParse := strings.Split(fields, ",")
	roop := make([]string, len(fieldsParse))
	copy(roop, fieldsParse)
	nestFields := make(map[string][]string)
	offset := 0
	for k, v := range roop {
		l := strings.Split(v, ".")
		ok := false
		if len(l) > 1 {
			_, ok = nestFields[l[0]]
			nestFields[l[0]] = append(nestFields[l[0]], l[1])
		}
		if ok {
			fieldsParse = append(fieldsParse[:(k-offset)], fieldsParse[(k+1-offset):]...)
			offset += 1
		} else {
			fieldsParse[k-offset] = l[0]
		}
	}
	return fieldsParse, nestFields
}

func FieldToMap(model interface{}, fields []string, nestFields map[string][]string) map[string]interface{} {
	u := make(map[string]interface{})
	ts, vs := reflect.TypeOf(model), reflect.ValueOf(model)
	if len(nestFields) == 0 {
		for i := 0; i < ts.NumField(); i++ {
			var jsonKey string
			field := ts.Field(i)
			jsonTag := field.Tag.Get("json")

			if jsonTag == "" {
				jsonKey = field.Name
			} else {
				jsonKey = strings.Split(jsonTag, ",")[0]
			}

			if fields[0] == "*" || contains(fields, jsonKey) {
				u[jsonKey] = vs.Field(i).Interface()
			}
		}
	} else {
		for i := 0; i < ts.NumField(); i++ {
			var jsonKey string
			field := ts.Field(i)
			jsonTag := field.Tag.Get("json")

			if jsonTag == "" {
				jsonKey = field.Name
			} else {
				jsonKey = strings.Split(jsonTag, ",")[0]
			}

			if fields[0] == "*" || contains(fields, jsonKey) {
				_, ok := nestFields[jsonKey]
				if ok {
					f, n := ParseFields(strings.Join(nestFields[jsonKey], ","))
					if vs.Field(i).Kind() == reflect.Ptr {
						u[jsonKey] = FieldToMap(vs.Field(i).Elem().Interface(), f, n)
					} else {
						u[jsonKey] = FieldToMap(vs.Field(i).Interface(), f, n)
					}
				} else {
					u[jsonKey] = vs.Field(i).Interface()
				}
			}
		}
	}
	return u
}
