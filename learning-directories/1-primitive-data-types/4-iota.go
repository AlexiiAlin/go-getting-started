package main

import (
	"fmt"
)

// 4. Using Iota and constant expression

const (
	first  = iota + 6 // iota = 0, result 6
	second = iota + 2 // iota = 1, result 3
	third             // iota = 4 (previous + 1), result 4
)
const (
	fourth = iota // iota = 0
	fifth         // iota = 1
)

func main() {
	fmt.Println(first, second, third, fourth, fifth) // 6 3 4 0 1
}
