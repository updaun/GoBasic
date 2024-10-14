// 채널(Channel) 기초(5)

package main

import (
	"fmt"
	// "runtime"
)

func main() {
	// 채널(Channel)
	// Close : 채널 닫기, 주의 -> 닫힌 채널에 값 전송 시 패닉(예외) 발생
	// Range : 채널안에서 값을 꺼낸다. (순회), 채널 닫아야(Close) 반복문 종료 -> 채널이 열려 있고 값이 계속 들어오면 계속 돌아감

	// 예제1
	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		close(ch) // 5회 채널에 데이터를 보내고 채널을 닫는다.
	}()

	for i := range ch { // 채널에서 값을 꺼내온다. (채널이 Close 될 때까지)
		fmt.Println("ex1 : ", i)
	}
}
