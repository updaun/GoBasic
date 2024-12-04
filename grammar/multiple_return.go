package main

import "fmt"

func inputNums() []int {
	var score int
	nums := []int{}

	for i := 0; i < 5; i++ {
		fmt.Scanln(&score)
		nums = append(nums, score)
	}

	return nums
}

func calExam(arr ...int) (int, int, int) {
	var total int
	var upper90, under70 int

	for _, n := range arr {
		total += n
		if n >= 90 {
			upper90 += 1
		} else if n < 70 {
			under70 += 1
		}
	}
	return total, upper90, under70
}

func printResult(total, upper90, under70 int) {
	var result bool = true

	if total < 400 {
		fmt.Println("총점이 400점 미만입니다.")
		result = false
	}
	if upper90 < 2 {
		fmt.Println("90점 이상 과목 수가 2개 미만입니다.")
		result = false
	}
	if under70 > 0 {
		fmt.Println("70점 미만 과목이 있습니다.")
		result = false
	}

	if !result {
		fmt.Println("아이패드를 살 수 없습니다.")
	} else {
		fmt.Println("아이패드를 살 수 있습니다.")
	}
}

func main() {
	var nums []int
	var total, upper90, under70 int
	nums = inputNums()
	total, upper90, under70 = calExam(nums...)
	printResult(total, upper90, under70)
}
