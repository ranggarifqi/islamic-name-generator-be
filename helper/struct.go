package helper

import (
	"encoding/json"
	"errors"
	"reflect"
)

/** Doesn't support nested struct*/
func GetStructValue[R comparable](obj any, key string, fallback R) (*R, error) {
	objVal := reflect.ValueOf(obj)

	curStruct := objVal.Type()

	if curStruct.Kind() != reflect.Struct {
		return nil, errors.New("provided obj is not a struct")
	}

	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var m map[string]any

	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	val, ok := m[key].(R)

	if !ok {
		return &fallback, nil
	}

	return &val, nil
}
