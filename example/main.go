package main

import (
	"fmt"
	"github.com/earlcherry/gouter/src"
)

func main() {
	g := gouter.New()
	g.Add("test", YourStruct{},"Method")
	g.Add("test23", YourStruct{},"")
	g.Add("test3", YourStruct{},"")
	g.Run("test")
}

type YourStruct struct{}

func (y YourStruct) Method() {
	fmt.Printf("Yep, its work!")
}
