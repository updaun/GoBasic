// 함수 기초(1)

package main

import (
	"fmt"
	"strconv"
)

// 함수 선언 위치는 어디든 가능
func helloGolang() {
	fmt.Println("ex1 : Hello Golang!")
}

func say_one(m string) {
	fmt.Println("ex2 : ", m)
}

func sum(x int, y int) int {
	return x + y
}

func main() {
	// 함수
	// 선언 : func 키워드로 선언
	// func 함수명(매개변수) (반환타입 or 반환값 변수명) : 반환값 존재
	// func 함수명() (반환타입 or 반환값 변수명) : 매개변수 없음, 반환값 존재
	// func 함수명(매개변수) : 매개변수 존재, 반환값 없음
	// func 함수명() : 매개변수 없음, 반환값 없음
	// 타 언어와 달리 반환 값 (return value) 다수 가능

	// 예제1
	helloGolang()

	// 예제2
	say_one("Hello World!")

	// 예제3
	result := sum(5, 5)
	fmt.Println("ex3 : ", result)
	fmt.Println("ex3 : ", sum(15, 15))               // 숫자로 출력
	fmt.Println("ex3 : ", strconv.Itoa(sum(15, 15))) // 문자열로 변환해서 출력
}
