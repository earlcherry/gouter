package main

import (
	"fmt"
	"github.com/earlcherry/gouter/src"
)

func main() {
	g := gouter.New()
	g.Add("ololo", YourStruct{})
	g.Add("test23", YourStruct{})
	g.Add("test3", YourStruct{})
	err:= g.Run("test234")
	if err!= nil{
		fmt.Printf(err.Error())
	}
}

type YourStruct struct {
	Text string
}

func (y YourStruct) Run()  {
	fmt.Printf("Yep, its work!",y.Text)
}
