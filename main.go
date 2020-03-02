package main

import "fmt"

func main() {
	// fmt.Println("Hello, from a module, Gophers")
	var i int // initialize variable
	i = 42    // give value
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
