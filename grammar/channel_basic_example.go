package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	var num1, num2 int
	c := make(chan int)

	fmt.Scanln(&num1, &num2)

	go func() {
		c <- add(num1, num2)
	}()

	fmt.Println(<-c)
}
