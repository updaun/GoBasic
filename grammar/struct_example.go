package main

import "fmt"

type student struct {
	name   string
	gender string
	scores map[string]int
}

// 생성자 함수
func newStudent(name string, gender string) *student {
	s := student{name, gender, map[string]int{}}
	return &s
}

func main() {
	var numStudent, numSubject int
	var sName, sGender string
	var subjectName string
	var subjectScore int

	fmt.Scanln(&numStudent, &numSubject)

	var studentGroup []*student // 슬라이스 선언

	for i := 0; i < numStudent; i++ {
		fmt.Scanln(&sName, &sGender)
		s := newStudent(sName, sGender)

		for j := 0; j < numSubject; j++ {
			fmt.Scanln(&subjectName, &subjectScore)

			s.scores[subjectName] = subjectScore
		}
		studentGroup = append(studentGroup, s)
	}

	for _, s := range studentGroup {
		fmt.Println("----------")
		fmt.Println(s.name, s.gender)

		for k, v := range s.scores {
			fmt.Println(k, v)
		}

	}
	fmt.Println("----------")
}
