package main

import (
	"slices"
	"strconv"
)

func calc_addr_b14_f1() {}

func calc_addr_b14_f2(group string, channel int) (res int) {

	// если вычислений не пройзойдет (неправильно указанные данные)
	// функция вернет -1
	res = -1

	channels := map[string]int{
		"AI": 8, "F1": 1, "F2": 1,
		"DI": 8, "EI": 12, "B": 8,
		"P": 16, "W": 64,
	}

	// выбираем стартовый адрес
	startAddr := 0
	switch group {
	// входные аналоговые значения
	case "AI":
		startAddr = 0
		// кнопка Ф1 и Ф2, чтобы это не значило
	case "F1":
		startAddr = 0x0100
	case "F2":
		startAddr = 0x0110
		// собственные дискретные входа
	case "DI":
		startAddr = 0x0200
		// входные каналы на шине расширения
	case "EI":
		startAddr = 0x0300
		// расчетные каналы
	case "P":
		startAddr = 0x0400
		// внешние каналы
	case "B":
		startAddr = 0x0500
		// выходные каналы
	case "W":
		startAddr = 0x0600
	default:
		return
	}

	// проверка наличия канала по всем группам, а после расчет
	switch group {
	case "AI", "DI", "EI", "B", "P", "F1", "F2", "W":
		if size, ok := channels[group]; ok && size >= channel {
			numOfBits := 16 // количество битов (для интервала) между адресами
			// так как начинаем расчет с 0 адреса
			// пример: группа I4 - startAddr = (h:0x0300|d:768)
			// канал 8; получаем расчет:
			// (0x0300|d:768) + (h:0x0040|d:8*8=64) = (h:0x0340|d:832)
			res = finalCalc(startAddr, channel, numOfBits)
		}
	}
	return
}

func calc_addr_b14_f3(group string, channel string) (res int) {

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
		"specq":      0x010E,
		"specD":      0x0120,
	}

	// список адресов для чтения адресов
	deviceTime := map[string]int{
		"YEAR": 0xFF00, "MONTH": 0xFF01,
		"DAY": 0xFF02, "HOUR": 0xFF03,
		"MIN": 0xFF04, "SEC": 0xFF05,
	}

	// список каналов
	// ключ - группа, значение - колличество каналов в группе.
	channels := map[string]int{
		"P": 16, "HI": 8, "HEXT": 8, "HP": 16, "HB": 8,
	}

	startAddr := 0
	// задаем стартовый канал; в случае времени, основного
	// и внешнего контуров - возвращаем значение сразу
	switch group {
	case "CLM":
		if addr, ok := controlLoop[channel]; ok {
			return addr
		}
	case "CLEXT":
		// Во внешнем контуре нет режима работы и спец. алгоритмов
		if slices.Contains([]string{"ctrlMod", "specq", "specD"}, channel) {
			return
		}
		shift := 0x1000 // смещение по адресу для внешнего контура
		if addr, ok := controlLoop[channel]; ok {
			return addr + shift
		}

	case "P":
		startAddr = 0x8000
	case "HI":
		startAddr = 0x9000
	case "HP":
		startAddr = 0x9010
	case "HB":
		startAddr = 0x9020
	case "TIME":
		return deviceTime[channel]
	}

	// преобразуем строку channel в int, так как далее будем смотреть
	//
	if intChannel, err := strconv.Atoi(channel); err == nil {
		if size, ok := channels[group]; ok && size >= intChannel && intChannel > 0 {
			numOfWords := 2 // количество слов (для интервала между адресами)
			res = finalCalc(startAddr, intChannel, numOfWords)

		}
	}

	return
}
func calc_addr_b14_f4(group string, channel int) (res int) {

	// если вычислений не пройзойдет (неправильно указанные данные)
	// функция вернет -1
	res = -1

	channels := map[string]int{
		"I": 8, "EXT": 8, "P": 16, "B": 8,
	}

	startAddr := 0
	switch group {
	case "I":
		startAddr = 0
	case "EXT":
		startAddr = 0x0100
	case "P":
		startAddr = 0x0200
	case "B":
		startAddr = 0x0300
	}

	// проверка наличия канала по группе в карте channels, а после расчет
	if size, ok := channels[group]; ok && size >= channel && channel > 0 {
		numOfWords := 2 // количество слов (для интервала между адресами)
		res = finalCalc(startAddr, channel, numOfWords)
	}

	return
}

func calc_addr_b14_f5()  {}
func calc_addr_b14_f6()  {}
func calc_addr_b14_f16() {}
