// IF문(2)

package main

import "fmt"

func main() {
	var a int = 50
	b := 70

	// 예제1
	if a >= 65 {
		fmt.Println("65 이상")
	} else {
		fmt.Println("65 미만")
	}

	// 예제2
	if b >= 70 {
		fmt.Println("70 이상")
	} else {
		fmt.Println("70 미만")
	}

	// 에러 발생
	/*
			if a >= 65 {
			fmt.Println("65 이상")
		} else
		{
			fmt.Println("65 미만")
		}
	*/
}
