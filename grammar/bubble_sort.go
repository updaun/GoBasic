package main

import "fmt"

func bubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

}

func inputNums() []int {
	var n, num int
	fmt.Scanln(&n)
	var nums []int
	for i := 0; i < n; i++ {
		fmt.Scanln(&num)
		nums = append(nums, num)
	}
	return nums
}

func outputNums(nums []int) {
	for _, n := range nums {
		fmt.Printf("%d ", n)
	}
}

func main() {
	nums := inputNums()
	bubbleSort(nums)
	outputNums(nums)
}
