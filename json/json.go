// Reflection adapter for jsonparser
package json

import (
	"reflect"
	"github.com/buger/jsonparser"
	"github.com/d7561985/gotools/log"
)

func Read(p []byte, v reflect.Value, tag string, root []string) error {

	switch v.Kind() {
	case reflect.String:
		vv, _, _, err := jsonparser.Get(p, tag)
		if err != nil {
			return err
		}
		v.SetString(string(vv))
	case reflect.Slice:
		root = append(root, tag)

		item := reflect.New(v.Type().Elem()).Elem()

		_, err := jsonparser.ArrayEach(p, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

			//! support only string data on slices
			if item.Kind() == reflect.String {
				item.SetString(string(value))
			} else {
				log.ErrorF("reflect.Slice not suported field type: %s", item.Kind().String())
			}

			v.Set(reflect.Append(v, item))
		}, root...)
		if err != nil {
			return err
		}

		root = root[:len(root)-1]

	case reflect.Struct:
		//! no empty tag
		if tag != "" {
			root = append(root, tag)
		}

		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			tg := field.Tag.Get("json")

			v2 := v.FieldByName(field.Name)
			err := Read(p, v2, tg, root)
			if err != nil {
				return err
			}
		}
		if tag != "" {
			root = root[:len(root)-1]
		}
	}
	return nil
}

func Unmarshal(p []byte, i interface{}) error {
	v := reflect.ValueOf(i).Elem()
	return Read(p, v, "", make([]string, 0, 10))
}
