package main

import (
	"bmacalc/internal/calc"
	"fmt"
)

func main() {
	var (
		b  string
		f  string
		gr string
		ch string
	)

	fmt.Scanf("%s %s %s %s", &b, &f, &gr, &ch)

	device, err := calc.NewBasis(b)
	if err != nil {
		panic(err)
	}

	if addr, err := device.CalcAddr(f, gr, ch); err == nil {
		fmt.Printf("hex: %X, dec: %d\n", addr, addr)
	} else {
		fmt.Println(err)
	}

}
