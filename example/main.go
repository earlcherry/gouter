package main

import(
	"fmt"
	"github.com/earlcherry/gouter/src"
)


func main() {
	gouter.Call(YourStruct{}, "Method","Yep, its work!")
}

type YourStruct struct{}

func (y YourStruct) Method(test string) {
	fmt.Printf(test)
}
