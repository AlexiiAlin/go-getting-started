package main

import (
	"fmt"
)

// 2. Working with pointers
func main() {
	// Arithmetic pointer operations are removed from go
	// Ex: *(firstName + 1)

	var firstName *string = new(string)
	*firstName = "Arthur"
	fmt.Println(*firstName) // Arthur

	secondName := "Arthur2"
	fmt.Println(secondName) // Arthur2

	ptr := &secondName
	fmt.Println(ptr, *ptr) // 0x40c128 Arthur2

	secondName = "Tricia"
	fmt.Println(ptr, *ptr) // 0x40c128 Tricia

}
