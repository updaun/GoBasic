package main

import "fmt"

/*타입문 작성은 선택입니다*/

func calCoin(coinCount int, coinValue int) func() int {
	return func() int { //클로저
		return coinCount * coinValue
	}
}

func main() {
	var coin10, coin50, coin100, coin500 int
	fmt.Scan(&coin10, &coin50, &coin100, &coin500)

	if coin10 < 0 || coin50 < 0 || coin100 < 0 || coin500 < 0 {
		fmt.Println("잘못된입력입니다.")
		return
	}

	add10 := calCoin(coin10, 10)
	add50 := calCoin(coin50, 50)
	add100 := calCoin(coin100, 100)
	add500 := calCoin(coin500, 500)

	totalmoney := add10() + add50() + add100() + add500()

	fmt.Println(totalmoney)
}
