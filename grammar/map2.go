package main

import "fmt"

func main() {
	subjects := make(map[string]int)
	var avg float32

	var subject string
	var score int

	for {
		fmt.Scanln(&subject, &score)
		if subject == "0" {
			break
		}
		subjects[subject] = score
	}

	avg = 0
	for g, s := range subjects {
		fmt.Println(g, s)
		avg += float32(s)
	}

	fmt.Printf("%.2f", avg/float32(len(subjects)))

}
