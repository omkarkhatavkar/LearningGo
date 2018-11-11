package main

import (
	"fmt"
	"math"
)

func main() {
	number, err := getSqt(-99.0)
	fmt.Println(number, err)
}

func getSqt(num float64) (float64, error) {
	if num < 0 {
		return 0.0, fmt.Errorf("\nERROR: No Negative NUMBER (%f)", num)
	} else {
		return math.Sqrt(num), nil
	}
}
