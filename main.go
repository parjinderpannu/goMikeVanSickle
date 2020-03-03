package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	port := 3000
	startWebServer(port, 2)
}

func startWebServer(port, numberOfRetries int) {
	fmt.Println("Starting server ...")
	// do imp stuff
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
}
