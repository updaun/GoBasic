// 자료형 : 포인터(1)

package main

import "fmt"

func main() {
	// 포인터
	// Go : 포인터 지원(C)
	// 변수의 지역성, 연속된 메모리 참조 ..., 힙, 스택
	// 파이썬, 자바 (JVM) -> 컴파일러, 인터프리터
	// 포인터지원 X (파이썬, C#, JAVA 등)
	// 주소의 값은 직접 변경 불가능 (잘못된 코딩으로 인한 버그 방지)
	// *(애스터리스크)로 사용
	// nil로 초기화 (nil != 0)

	// 예제1
	var a *int            // 방법1
	var b *int = new(int) // 방법2

	fmt.Println(a) // nil
	fmt.Println(b) // 0xc0000140d8

	i := 7
	fmt.Println("ex1 : ", i, &i)

	a = &i // 주소값 대입
	b = &i

	*a = 77

	fmt.Println("ex1 : ", a, &i)
	fmt.Println("ex1 : ", &a)
	fmt.Println("ex1 : ", *a) // 역참조값 출력

	fmt.Println()

	fmt.Println("ex1 : ", b, &i)
	fmt.Println("ex1 : ", &b)
	fmt.Println("ex1 : ", *b)

	var c = &i
	d := &i

	// *d = 10

	fmt.Println("ex1 : ", c, &i)
	fmt.Println("ex1 : ", &c)
	fmt.Println("ex1 : ", *c)

	fmt.Println()

	fmt.Println("ex1 : ", d, &i)
	fmt.Println("ex1 : ", &d)
	fmt.Println("ex1 : ", *d)

}
