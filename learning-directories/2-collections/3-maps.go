package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"foo": 42}
	fmt.Println(m)        // map[foo:42]
	fmt.Println(m["foo"]) // 42

	m["foo"] = 27
	fmt.Println(m) // map[foo:27]

	delete(m, "foo")
	fmt.Println(m) // map[]
}
