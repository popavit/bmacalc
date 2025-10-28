package main

import (
	"bmacalc/internal/calc"
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	query := scanner.Text()

	basisType, mFunc, group, channel, err := calc.ParseString(query)
	if err != nil {
		panic(err)
	}

	device, err := calc.NewBasis(basisType)
	if err != nil {
		panic(err)
	}

	if addr, err := device.CalcAddr(mFunc, group, channel); err == nil {
		fmt.Printf("hex: %X, dec: %d\n", addr, addr)
	} else {
		fmt.Println(err)
	}

}
