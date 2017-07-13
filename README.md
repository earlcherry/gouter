# Gouter
Simple non http router.  
Installation with dep:
```bash
dep ensure github.com/earlcherry/gouter
```
Example:
```go
package main

import (
	"fmt"
	"github.com/earlcherry/gouter"
)

func main() {
	g := gouter.New()
	g.Add("test1", YourStruct{})
	g.Add("test2", YourStruct{})
	g.Add("test3", YourStruct{})
	g.Run("test1")
}

type YourStruct struct{}

func (y YourStruct) Run() {
	fmt.Printf("Yep, its work!")
}
```