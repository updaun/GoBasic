// Switch문(1)
package main

import "fmt"

func main() {
	// 제어문(조건문) - Switch
	// Switch 뒤 표현식(expression) 생략 가능
	// case 뒤 표현식(expression) 사용 가능
	// 자동 break 때문에 fallthrough 존재
	// Type 분기 가능

	// 예제1
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	// 예제2
	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	// 예제3
	switch c := "go"; c {
	case "go":
		fmt.Println("Go!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("일치하는 값 없음")
	}

	// 예제4
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("Golang!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("None!")
	}

	// 예제5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i == j:
		fmt.Println("i는 j와 같다")
	case i > j:
		fmt.Println("i는 j보다 크다")
	}
}
