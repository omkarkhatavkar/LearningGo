package main

import "fmt"

func main() {
	var name, surname = "Omkar", "khatavkar"
	var age = 29
	surname = "khatavkar"
	size := 1.3
	fmt.Println("Hello "+name, surname+"\n How are you ?")
	fmt.Println("Your age is  ? ")
	fmt.Println(age)
	fmt.Printf("%T", size)
}
