//Panic & Recover(2)

package main

import "fmt"

func runFunc() {
	defer func() {
		s := recover()
		fmt.Println("Error Message: ", s)
	}()

	panic("Error occurred!")
	fmt.Println("test") // 실행되지 않음
}

func main() {
	// Go Recover 함수
	// 에러 복구 가능
	// Panic으로 발생한 에러를 복구 후 코드 실행을 재개할 수 있게함
	// 프로그램 종료 되지 않는다.
	// 즉, 코드 흐름을 정상상태로 복구하는 기능
	// Panic 에서 설정한 메세지를 받아올 수 있다.

	// 예제
	runFunc()

	fmt.Println("Hello Golang!")
}
