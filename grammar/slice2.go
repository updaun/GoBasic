package main

import "fmt"

func main() {
	names := make([]string, 10)

	var name string

	for {
		fmt.Scanln(&name)
		names = append(names, name)
		if name == "1" {
			break
		}
	}

	var result string = names[0]

	for _, i := range names[1:] {
		if len([]rune(result)) < len([]rune(i)) {
			result = i
		}
	}

	fmt.Println(result, len([]rune(result)))
}
