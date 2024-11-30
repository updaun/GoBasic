package main

import "fmt"

func main() {
	var num int
	//i,j는 두 개의 반복문에 쓰일 변수

	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		temp := ""
		for j := 0; j < i; j++ {
			temp += "o "
		}
		temp += "* "
		fmt.Println(temp)
	}
}
