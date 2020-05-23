package main

import (
	"fmt"
)

// 1. Declaring variables with primitives
func main() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	c := complex(3, 4)
	fmt.Println(c) // (3+4i)

	r, im := real(c), imag(c)
	fmt.Println(r, im) // 3 4
}
