// IF문(1)

package main

import "fmt"

func main() {
	// 제어문(조건문)
	// IF 문 : 반드시 Boolean 검사 -> 1, 0 (사용 불가 : 자동 형 변환 불가)
	// 소괄호 사용 X

	var a int = 20
	b := 20

	// 예제1
	if a >= 15 {
		fmt.Println("15 이상")
	}

	// 예제2
	if b >= 25 {
		fmt.Println("25 이상")
	}

	// 에러 발생1
	/*
		if b >= 25
		{
			fmt.Println("25 이상")
		}
	*/

	// 에러 발생2
	/*
		if b >= 25
			fmt.Println("25 이상")
	*/

	// 에러 발생3
	/*
		if c := 1; c {
			fmt.Println("True")
		}
	*/
	if c := true; c {
		fmt.Println("True")
	}

	// 예제3
	if d := 40; d >= 35 {
		fmt.Println("35 이상")
	}
	// d += 20 // 에러 발생 : if문에서 사용한 변수는 if문 내에서만 사용 가능

}
