package main

import (
	"fmt"
	"stringUtils"
)

var (
	a = 10
	b, c int = 20, 30
)

func main() {
	d := 60
	fmt.Printf(stringUtils.Reverse("\noG ,olleH"))
	fmt.Printf("%v + %v = %v", a, b, a+b)
	fmt.Printf("%v + %v = %v", c, d, add(c,d))
}

func add(n1, n2) int {
	return n1+n2
}
