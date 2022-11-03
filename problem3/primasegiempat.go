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

func PrimaSegiEmpat(wide, high, start int) string {
	var result string
	var number int = start
	var sumNumber int

	for i := 1; i <= high; i++ {
		for j := 1; j <= wide; j++ {
			number++

			for !checkNumber(number) { //ketika fo false maka akan dilanjutkan ke proses selanjutnya jka masih true akan terus berulang.
				number++
			}
			result = result + fmt.Sprintf("%d", number) + "	"
			sumNumber = sumNumber + number
		}
		result = result + "\n"

	}
	result += fmt.Sprintf("%d", sumNumber)
	return result

}

func main() {

	// fmt.Println("=========================")
	fmt.Println(PrimaSegiEmpat(2, 3, 13))
	// fmt.Println("=========================")
	// fmt.Println(checkNumber(15))
	// /*
	// 	17	19
	// 	23	29
	// 	31	37
	// 	156
	// */
	fmt.Println(PrimaSegiEmpat(5, 2, 1))
	// // /*
	// // 	2	3	5	7	11
	// // 	13	17	19	23	29
	// // 	129
	// // */
}
