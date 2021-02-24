package tool

import (
	"errors"
	"reflect"
)

func TranslateStruct(src interface{}, tgt interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	if reflect.TypeOf(tgt).Kind() != reflect.Ptr {
		return errors.New("target not ptr")
	}

	tgt_val := reflect.ValueOf(tgt).Elem()
	src_val := reflect.ValueOf(src)

	if src_val.Kind() == reflect.Ptr {
		src_val = src_val.Elem()
	}
	if tgt_val.Kind() != reflect.Struct || src_val.Kind() != reflect.Struct {
		return errors.New("not struct")
	}

	// src_map := make(map[string]bool, 10)
	for i := 0; i < src_val.NumField(); i++ {
		t1 := src_val.Type()
		name1 := t1.Field(i).Name
		// src_map[name1] = true
		v2 := tgt_val.FieldByName(name1)
		if !v2.IsValid() {
			continue
		}
		switch v2.Type().Kind() {
		case reflect.Array, reflect.Slice, reflect.Map, reflect.Ptr, reflect.Chan, reflect.Struct:
			continue
		}
		v1 := src_val.Field(i)
		if v1.Kind() == v2.Kind() {
			v2.Set(v1)
		} else {
			v2.Set(v1.Convert(v2.Type()))
		}
	}
	return nil

}
