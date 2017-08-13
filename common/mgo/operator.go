package mgo

import (
	"strings"
	"errors"
	"reflect"
)

type Operator interface {
	Exec(resultPtr interface{}) (err error)
}


func Mgo(resultPtr interface{}, operate string,option map[string]interface{}) error {
	o := getOperator(operate)
	if o == nil {
		return errors.New("the db-operate " + operate + " does not exist!")
	}

	v := reflect.ValueOf(o).Elem()
	for key, val := range option {
		value := v.FieldByName(key)
		if value == (reflect.Value{}) || !value.CanSet() {
			continue
		}
		value.Set(reflect.ValueOf(val))
	}
	return o.Exec(resultPtr)
}

func getOperator(operate string) Operator  {
	switch strings.ToLower(operate) {
	case "list":
		return new(List)
	default:
		return nil
	}
}