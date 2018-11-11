package main

import "fmt"

func main() {

	a := 100
	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T type\n", b)
	fmt.Printf("%T type\n", a)

	// to access the value of pointer
	fmt.Println(*b)

}
