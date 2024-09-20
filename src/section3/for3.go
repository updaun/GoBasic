// For문(3)

package main

import "fmt"

func main() {
	// 예제 1
loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break loop1
			}
			fmt.Println("ex1 : ", i, j)
		}
	}

	// 예제 2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex2 : ", i)
	}

loop2:
	// 에러 발생(loop 레이블 밑에 관련이 없는 소스 코드)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue loop2
			}
			fmt.Println("ex3 : ", i, j)
		}
	}
}
