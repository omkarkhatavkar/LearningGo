package main

import "fmt"

func main() {

	fmt.Println("Let's do closure!")
	sum := adder()
	for i := 1; i <= 12; i++ {
		fmt.Println(sum(i))
	}

}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
