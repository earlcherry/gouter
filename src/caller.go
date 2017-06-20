package gouter

import (
	"reflect"
)

func Call(controller interface{}, method string, args... interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(controller).MethodByName(method).Call(inputs)
}


