// For문(2)

package main

import "fmt"

func main() {
	// 예제 1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println("ex1 : ", sum1)

	// 예제 2
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
		//j := i++ // Go에서 후치연산은 반환값을 사용할 수 없음
	}
	fmt.Println("ex2 : ", sum2)

	// 예제 3
	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3 : ", sum3)

	// 예제 4
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4 : ", i, j)
	}

	// 에러 발생
	// for i, j := 0, 0; i <= 10; i, j := i++, j+10 {
	// 	fmt.Println("ex4 : ", i, j)
	// }

}
