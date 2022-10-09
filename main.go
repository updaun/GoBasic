package main

import (
	"fmt"
	"strings"
)

// func multiply(a int, b int) int {
func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// fmt.Println(multiply(2, 2))

	// totalLength, upperName := lenAndUpper("updaun")
	// fmt.Println(totalLength, upperName)

	// totalLength, _ := lenAndUpper("updaun")
	// fmt.Println(totalLength)
	repeatMe("updaun", "python", "ai", "sql", "django")
}
