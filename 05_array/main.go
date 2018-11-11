package main

import "fmt"

func main() {

	fruitnames := []string{"orange", "apple", "banana", "jackfruit"}
	var taste [2]string
	taste[0] = "sweet"
	taste[1] = "bitter"

	fmt.Println(taste)
	//fmt.Println(fruitnames[1:])
	fmt.Println(fruitnames[0:3])
}
