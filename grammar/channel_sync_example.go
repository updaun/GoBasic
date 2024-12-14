package main

import (
	"fmt"
	// "time"
)

func main() {

	c := make(chan bool)

	go func() {
		for i := 0; i < 20; i++ {
			c <- true
		}
		fmt.Println("송신 루틴 완료")
	}()

	for j := 0; j < 20; j++ {
		<-c
	}

	// time.Sleep(time.Second*3)

}
