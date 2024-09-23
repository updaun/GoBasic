// 데이터 타입 : Bool(2)

package main

import "fmt"

func main() {
	// 예제1 - 논리 연산자
	fmt.Println("ex1 : ", true && true)
	fmt.Println("ex2 : ", true && false)
	fmt.Println("ex3 : ", false && false)
	fmt.Println("ex4 : ", true || true)
	fmt.Println("ex5 : ", true || false)
	fmt.Println("ex6 : ", false || false)
	fmt.Println("ex7 : ", !true)
	fmt.Println("ex8 : ", !false)

	// 예제2 - 비교 연산자
	num1 := 15
	num2 := 37

	fmt.Println("ex9 : ", num1 < num2)
	fmt.Println("ex10 : ", num1 > num2)
	fmt.Println("ex11 : ", num1 >= num2)
	fmt.Println("ex12 : ", num1 <= num2)
	fmt.Println("ex13 : ", num1 != num2)
	fmt.Println("ex14 : ", num1 == num2)

}
