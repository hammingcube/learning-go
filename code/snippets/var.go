package main

import (
	"fmt"
)

var x, y, z = 33, "hello", false

var u int = 33
var v string = "hello"
var w bool = false

var (
	e = 33
	f = "hello"
	g = false
)

func main() {
	var a, b, c = 33, "hello", false
	p, q, r := 33, "hello", false

	fmt.Println(a, b, c)
	fmt.Println(p, q, r)
	fmt.Println(x, y, z)
	fmt.Println(u, v, w)
	fmt.Println(e, f, g)
}
