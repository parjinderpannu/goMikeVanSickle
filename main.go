package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	port := 3000
	err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) error {
	fmt.Println("Starting server ...")
	// do imp stuff
	fmt.Println("Server started on port", port)
	return nil
}
