package main

func main() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			continue
		}
		println("continuing...")
	}
	/*
		0
		continuing...
		1
		continuing...
		2
		3
		continuing...
		4
		continuing...
	*/
}
