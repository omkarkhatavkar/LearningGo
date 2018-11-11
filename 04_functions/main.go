package main

import (
	"fmt"
	"strconv"
)

func greeting(name string) string {
	return "Hellp " + name
}

func tellMeAge(age int) string {
	return "Your age is => " + strconv.FormatInt(int64(age), 10)
}

func main() {
	fmt.Println(greeting("Omkar"))
	fmt.Println(tellMeAge(29))
}
