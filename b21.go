package main

import (
	"slices"
)

func calc_addr_b21_f1() {}

func calc_addr_b21_f2(group string, channel int) (res int) {

	// если вычислений не пройзойдет (неправильно указанные данные)
	// функция вернет -1
	res = -1

	// ключ - группа, значение - колличество каналов в группе.
	channels := map[string]int{
		"I1": 16, "I2": 16, "I3": 16,
		"I4":  8,
		"I11": 12, "I12": 12, "I13": 12, "I14": 12, "I15": 12,
		"I21": 12, "I22": 12, "I23": 12,
		"I31": 12, "I32": 12, "I33": 12,
		"P": 24, "IB": 128, "OB": 128,
		"W1": 5,
		"W2": 10, "W3": 10, "W4": 10,
		"W5": 10, "W6": 10, "W7": 10, "W8": 10, "W9": 10,
	}

	// список для расчетов (используются индексы)
	orderChannelsI := []string{
		"I1", "I2", "I3", "I4", "I11",
		"I12", "I13", "I14", "I15", "I21", "I22",
		"I23", "I31", "I32", "I33",
	}

	// список для расчетов (используются индексы)
	orderChannelsW := []string{
		"W1", "W2", "W3",
		"W4", "W5", "W6", "W7", "W8",
		"W9",
	}

	startAddr := 0
	switch group {
	case "I1", "I2", "I3", "I4", "I11",
		"I12", "I13", "I14", "I15", "I21", "I22",
		"I23", "I31", "I32", "I33":
		// по индексу из списка выше расчитываем начальный адрес для группы
		if i := slices.Index(orderChannelsI, group); i != -1 {
			interval := 0x0100 // интервал адресов между группами
			startAddr = i * interval
		}
	case "P":
		startAddr = 0x2000
	case "IB", "OB":
		startAddr = 0x3000
	case "W1", "W2", "W3", "W4", "W5",
		"W6", "W7", "W8", "W9":
		// по индексу из списка выше расчитываем начальный адрес для группы
		if i := slices.Index(orderChannelsW, group); i != -1 {
			interval := 0x0100 // интервал адресов между группами
			startAddr = 0x1000 + i*interval
		}
	}

	// проверка наличия канала по группе в карте channels, а после расчет
	if size, ok := channels[group]; ok && size <= channels[group] {
		numOfBits := 8 // количество битов (для интервала) между адресами
		channel--      // так как начинаем расчет с 0 адреса
		// пример: группа I4 - startAddr = (h:0x0300|d:768)
		// канал 8; получаем расчет:
		// (0x0300|d:768) + (h:0x0040|d:8*8=64) = (h:0x0340|d:832)
		res = startAddr + channel*numOfBits
	}

	return
}

func calc_addr_b21_f3(group string, channel string) (res int) {

	// если вычислений не пройзойдет (неправильно указанные данные)
	// функция вернет -1
	res = -1

	// список адресов для контура регулирования
	controlLoop := map[string]int{
		"ctrlMod":    0x0000,
		"ctrlConfig": 0x0001,
		"coefGroup":  0x0002,
		"specAlgNum": 0x0003,
		"setpoint":   0x0100,
		"valveValue": 0x0102,
		"Ko":         0x0104,
		"Kp":         0x0106,
		"Ti":         0x0108,
		"Td":         0x010A,
		"Tf":         0x010C,
		"specKo":     0x010E,
		"specKp":     0x0110,
		"specTi":     0x0112,
		"specD1":     0x0114,
		"specD2":     0x0116,
	}

	// список контуров регулирования для вычислений (по индексу)
	controlLoopList := []string{
		"CL1", "CL2", "CL3", "CL4",
		"CL5", "CL6", "CL7", "CL8",
	}

	// список адресов для чтения адресов
	deviceTime := map[string]int{
		"YEAR": 0xFF00, "MONTH": 0xFF01,
		"DAY": 0xFF02, "HOUR": 0xFF03,
		"MIN": 0xFF02, "SEC": 0xFF03,
	}

	switch group {
	case "CL1", "CL2", "CL3", "CL4",
		"CL5", "CL6", "CL7", "CL8":
		if i := slices.Index(controlLoopList, group); i != 1 {
			// по выбраному контуру регулирования, параметру
			// и промежутку между адресами считаем адрес
			interval := 0x1000
			res = controlLoop[channel] + i*interval
		}
	case "TIME":
		res = deviceTime[channel]
	}

	return
}

func calc_addr_b21_f4(group string, channel int) (res int) {

	// если вычислений не пройзойдет (неправильно указанные данные)
	// функция вернет -1
	res = -1

	// список для расчетов (используются индексы)
	orderChannelsI := []string{
		"I1", "I2", "I3", "I4",
		"I11", "I12", "I13", "I14",
		"I15", "I21", "I22", "I23",
		"I31", "I32", "I33",
		"P", "IB", "OB", "V1", "V2",
	}

	startAddr := 0
	switch group {
	case "I1", "I2", "I3", "I4",
		"I11", "I12", "I13", "I14", "I15",
		"I21", "I22", "I23",
		"I31", "I32", "I33":
		// по индексу из списка выше расчитываем начальный адрес для группы
		if i := slices.Index(orderChannelsI, group); i != -1 {
			startAddr = i * 0x0020
		}
	case "P":
		startAddr = 0x2000
	case "IB", "OB":
		startAddr = 0x3000
	case "V1":
		startAddr = 0x1000
	case "V2":
		startAddr = 0x1010

	}
	// проверка наличия канала по группе в карте channels, а после расчет
	if ok := slices.Index(orderChannelsI, group); ok != -1 {
		numOfWords := 2 // количество слов (для интервала) между адресами
		channel--       // так как начинаем расчет с 0 адреса
		res = startAddr + channel*numOfWords
	}

	return
}

func calc_addr_b21_f5()  {}
func calc_addr_b21_f6()  {}
func calc_addr_b21_f16() {}
