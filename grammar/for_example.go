package main

import "fmt"

func main() {
	var num1 int
	fmt.Scanln(&num1)

	for i := 1; i < 10; i++ {
		fmt.Printf("%d X %d = %d\n", num1, i, num1*i)
	}
}
