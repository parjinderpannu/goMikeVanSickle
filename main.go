package main

import "fmt"

func main() {

	firstName := "Arthur"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Tricia"
	fmt.Println(ptr, *ptr)

	// var firstName *string = new(string)
	// *firstName = "Arthur"
	// fmt.Println(*firstName)

	// fmt.Println("Hello, from a module, Gophers")
	var i int // initialize variable
	i = 42    // give value
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	// firstName := "Arthur"
	// fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

}
