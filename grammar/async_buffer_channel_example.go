package main

import "fmt"

func main() {
	c := make(chan bool, 50)

	go func() {
		for i := 0; i < 20; i++ {
			c <- true
		}
		fmt.Println("송신 루틴 완료")
	}()

	for j := 0; j < 20; j++ {
		fmt.Printf("수신한 데이터 : %t\n", <-c)
	}
}
