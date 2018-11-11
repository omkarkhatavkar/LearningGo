package main

import (
	"fmt"
)

func main() {

	var number int
	fmt.Println("Please Enter The Number : ")
	fmt.Scanf("%d", &number)
	fmt.Println(getFizzBuzz(number))
}

func getFizzBuzz(number int) string {
	if number%15 == 0 {
		return "FizzBuzz!"
	} else if number%3 == 0 {
		return "Fizz !"
	} else if number%5 == 0 {
		return "Buzz !"
	} else {
		return "nothing"
	}
}
