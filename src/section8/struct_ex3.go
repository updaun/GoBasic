// 구조체 심화(3)

package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) CalculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CalculateP(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main() {
	// 정리 : 구조체 인스턴스 값 변경 시 -> 포인터 전달, 보통의 경우 -> 값 전달
	kim := Account{"245-901", 10000000, 0.015}
	lee := Account{"245-902", 20000000, 0.025}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	fmt.Println()

	lee.CalculateD(1000000)
	kim.CalculateP(1000000)

	fmt.Println("ex2 : ", int(lee.balance))
	fmt.Println("ex2 : ", int(kim.balance))
}
