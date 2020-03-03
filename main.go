package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	port := 3000
	isStarted := startWebServer(port)
	fmt.Println(isStarted)
}

func startWebServer(port int) bool {
	fmt.Println("Starting server ...")
	// do imp stuff
	fmt.Println("Server started on port", port)
	return true
}
