// 구조체 기본(3)

package main

import "fmt"

type car struct {
	color string
	name  string
}

func main() {
	// 예제1
	c1 := car{"red", "220d"}
	c2 := new(car)
	c2.color, c2.name = "white", "sonata"
	c3 := &car{}
	c4 := &car{"black", "520d"}

	fmt.Println("ex1 : ", c1)
	fmt.Println("ex1 : ", c2)
	fmt.Println("ex1 : ", c3)
	fmt.Println("ex1 : ", c4)
	fmt.Println()

}
