package main

import (
	"fmt"

	"github.com/ridhwanaw/HelloWorldGo/stringutils"
)

func main() {
	s := "Hello World!"
	A := stringutils.Upper(s)
	fmt.Println(A)
}
