package models

import (
	//"fmt"
	"reflect"
	"strconv"
)

var ()

func CsvUnmarshal(record []string, v interface{}) error {
	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record) {
		return &FieldMismatch{s.NumField(), len(record)}
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().String() {
		case "string":
			f.SetString(record[i])
		case "int64":
			ival, err := strconv.ParseInt(record[i], 10, 64)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		default:
			return &UnsupportedType{f.Type().String()}
		}
	}
	return nil
}
