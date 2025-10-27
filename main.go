package main

import (
	"fmt"
	"strconv"
)

type basisType string

const (
	basis12  basisType = "b12"
	basis14  basisType = "b14"
	basis21  basisType = "b21"
	basis100 basisType = "b100"
	basisPV  basisType = "bpv"
)

type modbusFunc string

const (
	readCoil              modbusFunc = "f1"
	readDiscreteInput     modbusFunc = "f2"
	readHoldingRegister   modbusFunc = "f3"
	readInputRegister     modbusFunc = "f4"
	writeSingleCoil       modbusFunc = "f5"
	writeSingleRegister   modbusFunc = "f6"
	writeMultipleRegister modbusFunc = "f16"
)

func main() {
	var (
		b  basisType
		f  modbusFunc
		gr string
		ch string
	)

	fmt.Scanf("%s %s %s %s", &b, &f, &gr, &ch)
	if addr := calc_addr(b, f, gr, ch); addr != -1 {
		fmt.Printf("dec: %d, hex: %X, %s, %s\n", addr, addr, gr, ch)
	} else {
		fmt.Printf("Неверно введены параметры %s %s \n", gr, ch)
	}
}

func calc_addr(basis basisType, mFunc modbusFunc, group string, channel string) (res int) {
	switch basis {
	case basis12:
		switch mFunc {
		case readCoil:
		case readDiscreteInput:
		case readHoldingRegister:
		case readInputRegister:
		case writeSingleCoil:
		case writeSingleRegister:
		case writeMultipleRegister:
		}
	case basis14:
		switch mFunc {
		case readCoil:
		case readDiscreteInput:
			// преобразуем строку канала в int, если успешно - расчитываем адрес
			if iChannel, err := strconv.Atoi(channel); err == nil {
				res = calc_addr_b14_f2(group, iChannel)
			} else {
				//если нет - ошибка
				return -1
			}
		case readHoldingRegister:
			res = calc_addr_b14_f3(group, channel)
		case readInputRegister:
			if iChannel, err := strconv.Atoi(channel); err == nil {
				res = calc_addr_b14_f4(group, iChannel)
			} else {
				//если нет - ошибка
				return -1
			}
		case writeSingleCoil:
		case writeSingleRegister:
		case writeMultipleRegister:
		}
	case basis21:
		switch mFunc {
		case readCoil:
		case readDiscreteInput:
			// преобразуем строку канала в int, если успешно - расчитываем адрес
			if iChannel, err := strconv.Atoi(channel); err == nil {
				res = calc_addr_b21_f2(group, iChannel)
			} else {
				//если нет - ошибка
				return -1
			}
		case readHoldingRegister:
			res = calc_addr_b21_f3(group, channel)
		case readInputRegister:
			// преобразуем строку канала в int, если успешно - расчитываем адрес
			res = calc_addr_b21_f4(group, channel)
		case writeSingleCoil:
		case writeSingleRegister:
		case writeMultipleRegister:
		}
	case basis100:
		switch mFunc {
		case readCoil:
		case readDiscreteInput:
		case readHoldingRegister:
		case readInputRegister:
		case writeSingleCoil:
		case writeSingleRegister:
		case writeMultipleRegister:
		}
	case basisPV:
		switch mFunc {
		case readCoil:
		case readDiscreteInput:
		case readHoldingRegister:
		case readInputRegister:
		case writeSingleCoil:
		case writeSingleRegister:
		case writeMultipleRegister:
		}
	default:
		return -1
	}
	return res
}
