package main

import "fmt"

func main() {

	ids := []int{12, 2, 2, 2, 23, 4, 4, 11, 1, 1}

	for _, id := range ids {
		fmt.Printf("value is %d \n", id)
	}
}
