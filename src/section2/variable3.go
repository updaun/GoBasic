// 변수3
package main

import "fmt"

func main() {
	// 짧은 선언
	// 반드시 함수 안에서만 사용(전역 x), 선언 후 할당 예외 발생
	// 일회성으로 사용
	// 주로 제한된 범위의 함수내에서 사용할 경우 코드 가독성을 높일 수 있음 -> 추천

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	// shortVar1 := 10 // 예외 발생
	// shortVar3 = 10 // 예외 발생

	fmt.Println("shortVar1 : ", shortVar1)
	fmt.Println("shortVar2 : ", shortVar2)
	fmt.Println("shortVar3 : ", shortVar3)

	// Example
	if i := 10; i < 11 {
		fmt.Println("Short variable test success!")
	}
}
