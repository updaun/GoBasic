package main

import "fmt"

const g float32 = 9.8

type kinEnergy func() float32
type potEnergy func() float32

func calMechEnergy(energyFunc func() float32) float32 {
	result := energyFunc()
	return result
}

func main() {
	var m, v, h float32
	fmt.Scanln(&m, &v, &h)

	kinEnergy := func() float32 {
		return 0.5 * m * v * v
	}
	potEnergy := func() float32 {
		return m * g * h
	}

	ke := calMechEnergy(kinEnergy)
	pe := calMechEnergy(potEnergy)
	fmt.Println(ke, pe, ke+pe)
}
