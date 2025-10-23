package main

import "fmt"

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
		ch int
	)
	fmt.Scanf("%s %s %s %d", &b, &f, &gr, &ch)

	if addr := calc_addr(b, f, gr, ch); addr != -1 {
		fmt.Printf("dec: %d, hex: %X\n", addr, addr)
	} else {
		fmt.Println("Неверно введены параметры")
	}
}

func calc_addr(basis basisType, mFunc modbusFunc, group string, channel int) (res int) {
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
		case readHoldingRegister:
		case readInputRegister:
		case writeSingleCoil:
		case writeSingleRegister:
		case writeMultipleRegister:
		}
	case basis21:
		switch mFunc {
		case readCoil:
		case readDiscreteInput:
			res = calc_addr_b21_f2(group, channel)
		case readHoldingRegister:
		case readInputRegister:
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
