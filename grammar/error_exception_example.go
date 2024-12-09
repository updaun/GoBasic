package main

import (
	"fmt"
	// "log"
)

func inputSubNum() (result int, err error) {
	var num int
	fmt.Scanln(&num)

	if num > 0 {
		return num, nil
	}

	return num, fmt.Errorf("잘못된 과목 수입니다.")

}

func average(num int) (result float64, err error) {
	var score, total int
	var avg float64

	for i := 0; i < num; i++ {
		fmt.Scanln(&score)
		if score < 0 {
			return 0, fmt.Errorf("잘못된 점수입니다.")
		}
		total += score
	}

	avg = float64(total) / float64(num)

	return avg, nil
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			return
		}
	}()

	num, err1 := inputSubNum()

	if err1 != nil {
		panic(err1)
	}

	result, err2 := average(num)

	if err2 != nil {
		panic(err2)
	}

	fmt.Println(result)
}
