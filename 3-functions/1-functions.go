package main

import (
	"fmt"
)

func main() {
	port := 3000
	_, err := startWebServer(port)
	fmt.Println(port, err)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	// return errors.New("Something went wrong")
	return port, nil
}
