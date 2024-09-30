// 자료형 : 슬라이스 심화(3)

package main

import (
	"fmt"
)

func main() {
	// 슬라이스 복사
	// copy(복사 대상, 원본)
	// make로 공간을 할당 후 복사해야 한다. 그렇지 않으면 복사가 안됨
	// 복사 된 슬라이스 값 변경해도 원본에 영향을 주지 않는다.

	// 예제1(복사)
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1) // 복사 안됨

	fmt.Println("ex1 : ", slice2)
	fmt.Println("ex1 : ", slice3)

	fmt.Println()

	// 예제2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)

	b[0] = 7
	b[4] = 10

	fmt.Println("ex2 : ", a)
	fmt.Println("ex2 : ", b)

	fmt.Println()

	// 예제3
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] // 주의! 부분적으로 슬라이스 추출은 참조 -> 원본 값 변경 된다.

	d[1] = 7

	fmt.Println("ex3 : ", c)
	fmt.Println("ex3 : ", d)

	// 예제4
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] // 용량을 7로 설정 후 5까지만 추출

	fmt.Println("ex4 : ", len(f), cap(f), f)
}
