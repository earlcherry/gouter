package gouter

import (
	"reflect"
)

type Gouter struct {
	routes map[string]Goute
}

type Goute struct {
	Controller interface{}
	Method     string
}

func New() *Gouter {
	return &Gouter{
		routes: make(map[string]Goute),
	}
}

func (g *Gouter) Add(route string, controller interface{}, method string) {
	g.routes[route] = Goute{
		Controller: controller,
		Method:     method,
	}
}
func (g *Gouter) Run(route string, args... string) {
	goute := g.routes[route]

	method := reflect.ValueOf(goute.Controller).MethodByName(goute.Method)
	argsCount := method.Type().NumIn()

	inputs := make([]reflect.Value, argsCount)
	for i := 0; i < argsCount; i++ {
		if i >= len(args) {
			inputs[i] = reflect.ValueOf("")
		} else {
			inputs[i] = reflect.ValueOf(args[i])
		}
	}

	method.Call(inputs)
}
