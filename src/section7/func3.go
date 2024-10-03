// 함수 기초(3)

package main

import "fmt"

func multiply(x, y int) (int, int) {
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 1, b * 2, c * 3, d * 4, e * 5
}

func main() {
	// 다중 값 반환
	// 예제1
	a, b := multiply(10, 5)
	c, _ := multiply(10, 5)
	_, d := multiply(10, 5)

	println("ex1 : ", a, b)
	println("ex1 : ", c)
	println("ex1 : ", d)

	// 예제2
	x1, x2, x3, x4, x5 := arrayMultiply(1, 2, 3, 4, 5)
	y1, _, y3, _, y5 := arrayMultiply(1, 2, 3, 4, 5)
	fmt.Println("ex2 : ", x1, x2, x3, x4, x5)
	fmt.Println("ex2 : ", y1, y3, y5)
}
