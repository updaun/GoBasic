package main

import (
	"fmt"
	"strings"
)

// func multiply(a int, b int) int {
func multiply(a, b int) int {
	return a * b
}

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func repeatMe(words ...string) {
	fmt.Println(words)
}

// naked return
func lenAndUpper(name string) (length int, uppercase string) {
	// defer : 이 함수가 실행되고 나서 실행되는 부분
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
	// return 1

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	// return 1

	total := 0
	for _, number := range numbers {
		// total = total + number
		total += number
	}
	return total
}

func main() {
	// fmt.Println(multiply(2, 2))

	// totalLength, upperName := lenAndUpper("updaun")
	// fmt.Println(totalLength, upperName)

	// totalLength, _ := lenAndUpper("updaun")
	// fmt.Println(totalLength)

	// repeatMe("updaun", "python", "ai", "sql", "django")

	// total := superAdd(1, 2, 3, 4, 5, 6)
	// fmt.Println(total)

	// superAdd(1, 2, 3, 4, 5, 6)
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

}
