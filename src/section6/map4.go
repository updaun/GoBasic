// 자료형 : 맵(4)

package main

import "fmt"

func main() {
	// 맵(Map)
	// 맵 조회 할 경우에 주의할 점

	// 예제1
	map1 := map[string]int{ // 기본 초기화 값 int 0 string "" float 0.0 bool false
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1 := map1["apple"]
	value2, ok2 := map1["lemon"]
	value3, ok3 := map1["kiwi"] // 두 번째 리턴 값으로 키 존재 유무 확인

	fmt.Println("ex1 : ", value1)
	fmt.Println("ex1 : ", value2, ok2)
	fmt.Println("ex1 : ", value3, ok3)

	// 예제2
	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : Not Exist Kiwi")
	}

	if value, ok := map1["banana"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : Not Exist Banana")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex2 : Not Exist Kiwi")
	}
}
