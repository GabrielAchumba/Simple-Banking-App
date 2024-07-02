package conversion

import (
	"reflect"
)

// MergeStructs merges fields from src into dst
func SpreadOperation(dst, src interface{}) {
	dstVal := reflect.ValueOf(dst).Elem()
	srcVal := reflect.ValueOf(src).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		dstField := dstVal.FieldByName(srcVal.Type().Field(i).Name)

		if dstField.IsValid() && dstField.CanSet() {
			dstField.Set(srcField)
		}
	}
}
