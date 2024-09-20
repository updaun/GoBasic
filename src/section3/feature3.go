// Go 특징(3)

package main

import "fmt"

func main() {
	// 코드 서식 지정
	// 한 사람이 코딩한 것 같은 일관적인 스타일 유지
	// 코드 스타일 유지
	// gofmt -h : 사용법
	// gofmt -w : 원본파일에 반영

	// 예제1
	for i := 10; i <= 100; i += 10 {
		fmt.Println("ex1 : ", i)
	}
}
