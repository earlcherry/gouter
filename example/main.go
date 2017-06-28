package main

import (
	"fmt"
	"github.com/earlcherry/gouter/src"
	"os"
	"reflect"
)

func main() {
	args := os.Args[1:]
	g := gouter.New()
	g.Add("test", YourStruct{}, "Method")
	g.Add("test23", YourStruct{}, "")
	g.Add("test3", YourStruct{}, "")

	inputs := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		inputs[i] = reflect.ValueOf(args[i])
	}
	if len(inputs) < 2 {
		panic("Not enough arguments")
	}

	reflect.ValueOf(g).MethodByName("Run").Call(inputs)
}

type YourStruct struct{}

func (y YourStruct) Method(a string, b string, c string) {
	fmt.Printf(a)
	fmt.Printf(b)
	fmt.Printf(c)
	fmt.Printf("Yep, its work!")
}
