package helper

import (
	"errors"
	"reflect"

	"github.com/fatih/structs"
)

func GetStructValue[R any](obj any, key string, fallback R) (*R, error) {
	objVal := reflect.ValueOf(obj)

	curStruct := objVal.Elem()
	if curStruct.Kind() != reflect.Struct {
		return nil, errors.New("provided obj is not a struct")
	}

	m := structs.Map(obj)

	val := m[key]

	if reflect.ValueOf(&val).Elem().IsZero() {
		return &fallback, nil
	}

	return val.(*R), nil

}
