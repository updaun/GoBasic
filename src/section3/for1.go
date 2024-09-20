// For문(1)

package main

import "fmt"

func main() {

	// 반복문 - For

	// Go에서 유일하게 제공되는 반복문
	// 다양한 사용법 숙지

	// 예제 1
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	// 에러 발생 1

	// for i := 0; i < 5; i++
	// {
	// 	fmt.Println("ex1 : ", i)
	// }

	// 에러 발생 2
	// for i := 0; i < 5; i++
	// fmt.Println("ex1 : ", i)

	// 예제 2 (무한 루프)
	// for {
	// 	fmt.Println("ex2 : Hello, Go!")
	// 	fmt.Println("ex2 : Infinite Loop!")
	// }

	// 예제 3 (Range 용법)
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}
}
