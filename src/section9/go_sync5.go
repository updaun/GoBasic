// 고루틴 동기화 기초(5)

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 고루틴 동기화 객체
	// 동기화 상태(조건) 메소드 사용
	// Wait, Notify, NotifyAll : 기타 언어
	// Wait, Signal, Broadcast

	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) // 비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Waiting : ", n)
			condition.Wait() // 고루틴 대기(멈춤)
			fmt.Println("Waiting End : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c
		// fmt.Println("received data : ", <-c)
	}

	/*
		for i := 0; i < 5; i++ {
			mutex.Lock()
			fmt.Println("Wake Goroutine(Signal) : ", i)
			condition.Signal() // 잠자고 있는 고루틴을 깨움(실행대기)
			mutex.Unlock()
		}
	*/

	mutex.Lock()
	fmt.Println("Wake Goroutine(Broadcast)")
	condition.Broadcast() // 잠자고 있는 모든 고루틴을 깨움(실행대기)
	mutex.Unlock()

	time.Sleep(2 * time.Second)
}
