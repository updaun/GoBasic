package main

import "fmt"

func main() {

	for i := 1; i < 100; i++ {
		if i%7 == 0 {
			fmt.Printf("%d ", i)
		} else if i%9 == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
