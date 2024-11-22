// 파일 I/O (1)

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 파일 읽기, 쓰기 -> ioutil 패키지 활용
	// 더욱 편리하고 직관적으로 파일 읽고 쓰기 가능
	// 아래 메소드 확인 가능
	// WriteFile(), ReadFile(), ReadAll() 등 사용 가능
	// https://golang.org/pkg/io/ioutil/

	s := "Hello Golang!\nFile Write Test!\n"

	// 파일모드(읽기, 쓰기, 실행) -> 권한
	// 읽기: 4, 쓰기: 2, 실행: 1
	// 소유자, 그룹, 기타 사용자 순서 (777 -> 모두 읽기, 쓰기, 실행 가능)
	// https://golang.org/pkg/os/#FileMode
	err := ioutil.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644))
	errCheck(err)

	data, err := ioutil.ReadFile("sample.txt")
	errCheck(err)

	fmt.Println("=====================================")
	fmt.Println(string(data))
	fmt.Println("=====================================")
}
