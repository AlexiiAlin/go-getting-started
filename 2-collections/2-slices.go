package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3}
	slice := arr[:]
	fmt.Println(arr, slice) // [1 2 3] [1 2 3]

	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice) // [1 42 27] [1 42 27]

	// Another example (more commonly used)

	slice2 := []int{1, 2, 3}
	fmt.Println(slice2) // [1 2 3]

	slice2 = append(slice2, 4, 42, 27)
	fmt.Println(slice2) // [1 2 3 4 42 27]

	s3 := slice2[1:]  // [2 3 4 42 27]
	s4 := slice2[:2]  // [1 2]
	s5 := slice2[1:2] // [2]

	fmt.Println(s3, s4, s5) // [2 3 4 42 27] [1 2] [2]

}
