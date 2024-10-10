// 인터페이스 기본(3)

package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

// 구조체 Dog 메소드 구현
func (d Dog) bite() {
	fmt.Println(d.name, " : Dog bites!")
}

func (d Dog) sounds() {
	fmt.Println(d.name, " : Dog barks!")
}

func (d Dog) run() {
	fmt.Println(d.name, " : Dog is running!")
}

// 구조체 Cat 메소드 구현
func (c Cat) bite() {
	fmt.Println(c.name, " : Cat bites!")
}

func (c Cat) sounds() {
	fmt.Println(c.name, " : Cat cries!")
}

func (c Cat) run() {
	fmt.Println(c.name, " : Cat is running!")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
	sounds()
	run()
}

// 인터페이스를 파라미터로 받는다.
func act(animal Behavior) {
	animal.bite()
	animal.sounds()
	animal.run()
}

func main() {

	// 인터페이스 구현 예제
	// 인터페이스 규격화 역할 이해
	// 인터페이스에 정의된 메소드 사용 유도
	// 코드의 가독성 및 유지보수성 증가

	// 덕타이핑 예제
	// 덕타이핑 : 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식
	// Go 덕타이핑의 중요한 특징 : 오리처럼 걷고, 꽥꽥댄다면 오리로 간주한다.

	// 예제1
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	// 개 행동 실행
	act(dog)
	// 고양이 행동 실행
	act(cat)

}
