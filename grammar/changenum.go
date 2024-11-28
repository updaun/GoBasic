package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Scanln(&num1, &num2, &num3)

	var changenum1 float32 = float32(num1)
	var changenum2 uint = uint(num2)
	var changenum3 int64 = int64(num3)

	fmt.Printf("%T, %f\n", changenum1, changenum1)
	fmt.Printf("%T, %d\n", changenum2, changenum2)
	fmt.Printf("%T, %d\n", changenum3, changenum3)
}
