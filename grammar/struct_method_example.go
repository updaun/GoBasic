package main

import "fmt"

const g = 9.8

type dynamics struct {
	m, v, h, ke, pe, me float32
}

func cal_ke(d *dynamics) float32 {
	return 0.5 * d.m * d.v * d.v
}

func cal_pe(d *dynamics) float32 {
	return d.m * g * d.h
}

func main() {
	var n int

	fmt.Scanln(&n)

	var ds []dynamics

	var m, v, h float32
	var ke, pe, me float32

	for i := 0; i < n; i++ {
		fmt.Scanln(&m, &v, &h)
		ds = append(ds, dynamics{m, v, h, 0, 0, 0})
	}

	for i := 0; i < n; i++ {
		ke = cal_ke(&ds[i])
		pe = cal_pe(&ds[i])
		me = ke + pe
		fmt.Println(ke, pe, me)
	}

}
