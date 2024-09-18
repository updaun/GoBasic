// 열거형 3
package main

import "fmt"

func main() {
	const (
		_ = iota
		A
		B
		C
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		PLATINUM
	)

	fmt.Println("DEFAULT : ", DEFAULT)
	fmt.Println("SILVER : ", SILVER)
	fmt.Println("GOLD : ", GOLD)
	fmt.Println("PLATINUM : ", PLATINUM)
}
