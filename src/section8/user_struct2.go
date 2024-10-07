// 사용자 정의 타입(2)

package main

import "fmt"

type cnt int

func main() {
	// 기본 자료형 사용자 정의 타입
	// 예제1

	a := cnt(15)
	fmt.Println("ex1 : ", a)

	// 예제2
	var b cnt = 15
	fmt.Println("ex2 : ", b)

	// 예제3
	// testConverT(b) // 예외 발생 : 기본 자료형과 사용자 정의 타입은 서로 다른 타입이므로
	testConverT(int(b))

	c := 10

	testConverD(b)
	testConverD(cnt(c))

}

func testConverT(i int) {
	fmt.Println("ex3 : (Default Type) : ", i)
}

func testConverD(i cnt) {
	fmt.Println("ex4 : (Custom Type) : ", i)
}
