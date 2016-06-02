package main

import (
	"fmt"
	"stringUtils"
)

var (
	a = 40
	b int = 50
)

func main() {
	c := 60
	fmt.Printf(stringUtils.Reverse("\noG ,olleH"))
	fmt.Printf("%v + %v + %v = %v", a, b, c, a+b+c)
}

