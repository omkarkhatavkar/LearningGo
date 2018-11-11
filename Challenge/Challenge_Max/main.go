package main

import (
	"fmt"
)

func main() {
	numbers := []int{55, 2, 2, 3, 4, 55, 4, 56}
	fmt.Println(getMaxNumber(numbers))

}

func getMaxNumber(numbers []int) int {
	maxnumber := 0
	for _, num := range numbers {
		if num > maxnumber {
			maxnumber = num
		}
	}
	return maxnumber
}
