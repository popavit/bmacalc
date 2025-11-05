package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/popavit/dmacalc/calc"
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

		device, err := calc.NewDevice(basisType)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		groupList, _ := calc.GetGroup(device, mFunc)
		groupChannels, _ := calc.GetChannel(device, mFunc, group)
		fmt.Printf("Группы: %v\nКаналы у группы %q: %v\n", groupList, group, groupChannels)

		if addr, err := calc.CalcAddr(device, mFunc, group, channel); err == nil {
			fmt.Printf("hex: %X, dec: %d\n\n", addr, addr)
		} else {
			fmt.Println("Ошибка:", err)
		}
	}

}
