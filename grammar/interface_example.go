package main

import (
	"fmt"
	"math"
)

type shape interface {
	space() float32
	volume() float32
}

type cylinder struct {
	radius, height float32
}

type cube struct {
	a, b, c float32
}

func (r cylinder) space() float32 {
	return math.Pi*r.radius*r.radius*2 + 2*math.Pi*r.radius*r.height
}

func (r cylinder) volume() float32 {
	return math.Pi * r.radius * r.radius * r.height
}

func (c cube) space() float32 {
	return 2*c.a*c.b + 2*c.a*c.c + 2*c.b*c.c
}

func (c cube) volume() float32 {
	return c.a * c.b * c.c
}

func main() {
	cy1 := cylinder{10, 10}
	cy2 := cylinder{4.2, 15.6}
	cu1 := cube{10.5, 20.2, 20}
	cu2 := cube{4, 10, 23}

	printMeasure(cy1, cy2, cu1, cu2)
}

func printMeasure(m ...shape) {
	for _, val := range m {
		fmt.Printf("%.2f, %.2f\n", val.space(), val.volume())
	}
}
