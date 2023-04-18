package validatorn

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

var ErrValidation = errors.New("validation error")

func Validate(v any) error {

	if reflect.ValueOf(v).Kind() != reflect.Struct {
		return ErrValidation
	}
	value := reflect.ValueOf(v)

	for i := 0; i < reflect.TypeOf(v).NumField(); i++ {

		tag := reflect.TypeOf(v).Field(i).Tag.Get("validate")
		if len(tag) == 0 {
			continue
		}

		if !value.Type().Field(i).IsExported() {
			return ErrValidation
		}

		if reflect.TypeOf(v).Field(i).Type.Kind() == reflect.String {
			err := validate(tag, reflect.ValueOf(v).Field(i).String())
			if err != nil {
				return ErrValidation
			}
		}
	}

	return nil
}

func validate(tag string, value string) error {
	rulesArr := strings.SplitN(tag, ":", 2)
	rVal := rulesArr[1]

	if rVal == "" {
		return ErrValidation
	}
	r := strings.Split(rVal, ",")
	rl, err := strconv.Atoi(r[0])
	if err != nil {
		return ErrValidation
	}
	rr, err := strconv.Atoi(r[1])
	if err != nil {
		return ErrValidation
	}
	if len([]rune(value)) >= rl && len([]rune(value)) <= rr {
		return nil
	}
	return ErrValidation
}
