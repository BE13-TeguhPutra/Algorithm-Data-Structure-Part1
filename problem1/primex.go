package main

import (
	"fmt"
	"math"
)

func checkNumber(number int) bool {
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrtNumber; i++ {
		if number%i == 0 {
			return false
		}

	}
	return true

}

func PrimeX(number int) int {

	var x int = -1
	for i := 0; i < number; i++ {
		x = x + 1
		for !checkNumber(x) {
			// fmt.Println("input", x)
			x++
			// fmt.Println("out", x)
		}
		// fmt.Println("setelah", x)

	}
	return x
}

func main() {
	fmt.Println(PrimeX(1))  //2
	fmt.Println(PrimeX(5))  //11
	fmt.Println(PrimeX(8))  //19
	fmt.Println(PrimeX(9))  //23
	fmt.Println(PrimeX(10)) //29

}
