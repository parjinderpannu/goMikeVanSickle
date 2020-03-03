package main

import "fmt"

// const (
// 	first = iota
// 	second
// )

// const (
// 	third = iota
// 	fourth
// )

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)
	fmt.Println(u.LastName)

	u2 := user{ID: 1,
		FirstName: "Clark",
		LastName:  "Kent",
	}

	fmt.Println(u2)

	// m := map[string]int{"foo": 42}
	// fmt.Println(m)
	// fmt.Println(m["foo"])

	// m["foo"] = 27
	// fmt.Println(m)

	// delete(m, "foo")
	// fmt.Println(m)

	// slice := []int{1, 2, 3} //dynamic size array through slice
	// fmt.Println(slice)

	// slice = append(slice, 4, 42, 27)
	// fmt.Println(slice)
	// s2 := slice[1:]  // index 1 till last
	// s3 := slice[:2]  // index start till 1 not including 2
	// s4 := slice[1:2] // starting at index 1 and not including 2
	// // arr := [3]int{1, 2, 3}
	// fmt.Println(s2, s3, s4)

	// slice := arr[:]

	// arr[1] = 42
	// slice[2] = 27

	// fmt.Println(arr, slice)

	// var arr [3]int // easier for compiler
	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3
	// fmt.Println(arr)
	// fmt.Println(arr[1])
	// // fmt.Println(arr[4]) // invalid array index 4 (out of bounds for 3-element array)

	// arr2 := [3]int{1, 2, 3}
	// fmt.Println(arr2)

	// fmt.Println(first, second, third, fourth)
	// const pi = 3.1415 // declare & initialize at same line
	// fmt.Println(pi)

	// const c int = 3
	// fmt.Println(c + 3)

	// // a bunch of code
	// fmt.Println(float32(c) + 1.2)

	// firstName := "Arthur"
	// fmt.Println(firstName)

	// ptr := &firstName
	// fmt.Println(ptr, *ptr)

	// firstName = "Tricia"
	// fmt.Println(ptr, *ptr)

	// var firstName *string = new(string)
	// *firstName = "Arthur"
	// fmt.Println(*firstName)

	// fmt.Println("Hello, from a module, Gophers")
	// var i int // initialize variable
	// i = 42    // give value
	// fmt.Println(i)

	// var f float32 = 3.14
	// fmt.Println(f)

	// firstName := "Arthur"
	// fmt.Println(firstName)

	// b := true
	// fmt.Println(b)

	// c := complex(3, 4)
	// fmt.Println(c)

	// r, im := real(c), imag(c)
	// fmt.Println(r, im)

}
