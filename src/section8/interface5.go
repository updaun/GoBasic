// 인터페이스 기본(5)

package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func printValue(s interface{}) {
	fmt.Println("ex1 : ", s)
}

func main() {
	// 인터페이스 활용(빈 인터페이스)
	// 함수내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다. (만능) -> 모든 타입 지정 가능

	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	printValue(dog)
	printValue(cat)
	printValue(15)
	printValue("Hello, Golang!")
	printValue(3.14)
	printValue([]string{"a", "b", "c"})
	printValue(map[string]int{"key1": 1, "key2": 2})
	printValue(&dog)
	printValue(&cat)
	printValue([]Dog{})
	printValue([...]Cat{{"tiger", 23}})
}
