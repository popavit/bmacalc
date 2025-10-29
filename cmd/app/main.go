package main

import (
	"bmacalc/internal/calc"
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Для запроса адреса введите:",
		"\n<тип базиса> <номер функции> <группа/параметр> <канал/номер/параметр>",
		"\nПодробнее в документации.",
	)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		query := scanner.Text()

		basisType, mFunc, group, channel, err := calc.ParseString(query)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		device, err := calc.NewBasis(basisType)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		if addr, err := device.CalcAddr(mFunc, group, channel); err == nil {
			fmt.Printf("hex: %X, dec: %d\n\n", addr, addr)
		} else {
			fmt.Println("Ошибка:", err)
		}
	}

}
