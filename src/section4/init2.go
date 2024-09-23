// go 초기화 함수(2)

package main

import (
	"fmt"
)

// init 함수가 여러 개 여도 되지만 권장하지 않음

func init() {
	fmt.Println("Init1 Method Start!")
}
func init() {
	fmt.Println("Init2 Method Start!")
}
func init() {
	fmt.Println("Init3 Method Start!")
}

func main() {
	fmt.Println("Main Method Start!")
}
