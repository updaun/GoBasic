//Panic & Recover(3)

package main

import (
	"fmt"
	"os"
)

func fileOpen(filename string) {
	// defer 함수 (Panic 호출 되면 실행)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error: ", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("File Open OK!", f.Name())
	}

	defer f.Close()
}

func main() {
	// Go panic & Recover 최종 실습

	// 예제
	fileOpen("undefined.txt")
}
