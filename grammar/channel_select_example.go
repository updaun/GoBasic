package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go sendMessage(c)

	for i := 10; i > 0; i-- {
		select {
		case msg := <-c:
			fmt.Printf("%s 메세지가 발송되었습니다.\n", msg)
			return
		default:
			fmt.Printf("%d 초 안에 메세지를 입력하시오.\n", i)
			time.Sleep(time.Second)
		}
	}

	fmt.Println("메세지 발송에 실패했습니다.")
}

func sendMessage(c chan string) {
	var input string
	fmt.Scanln(&input)
	c <- input
}
