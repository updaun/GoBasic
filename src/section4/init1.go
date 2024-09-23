// go 초기화 함수(1)

package main

import (
	"fmt"
)

func init() {
	// init: 패키지 로드 시에 가장 먼저 호출
	// 가장 먼저 호출되는 함수이기 때문에 한 번만 호출
	fmt.Println("Init Method Start!")
}

func main() {
	fmt.Println("Main Method Start!")
}
