// 열거형 2

package main

import "fmt"

func main() {
	const (
		A = iota * 10
		B
		C
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
	)

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(Mar)
	fmt.Println(Apr)
	fmt.Println(May)
	fmt.Println(Jun)
}
