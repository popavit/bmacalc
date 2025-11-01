package calc

import (
	"fmt"
	"slices"
	"strconv"
)

type BasisPV struct{}

func (b *BasisPV) readCoil(group, channel string) (int, error) {
	return 0, fmt.Errorf("пока не реализивано")
}

func (b *BasisPV) readDiscreteInput(group, channel string) (res int, err error) {
	channels := map[string]int{
		"I": 4, "F": 2,
		"B": 8, "P": 16, "W": 64,
	}

	// выбираем стартовый адрес
	startAddr := 0
	switch group {
	// входные аналоговые значения
	case "I":
		startAddr = 0
		// кнопка Ф1 и Ф2, чтобы это не значило
	case "F":
		startAddr = 0x0100
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
		return 0, fmt.Errorf("неверно введена группа: %q", group)
	}

	// преобразуем строку канала в int
	iChannel, e := strconv.Atoi(channel)
	// проверяем на ошибку и на то, что каналы не отрицательные
	if e != nil || iChannel <= 0 {
		return 0, fmt.Errorf("неверно введен канал: %q", channel)
	}

	// проверка наличия канала по всем группам, а после расчет
	switch group {
	case "I", "DI", "EI", "B", "P", "F", "W":
		if size, ok := channels[group]; ok && size >= iChannel {
			numOfBits := 16 // количество битов (для интервала) между адресами
			// так как начинаем расчет с 0 адреса
			// пример: группа I4 - startAddr = (h:0x0300|d:768)
			// канал 8; получаем расчет:
			// (0x0300|d:768) + (h:0x0040|d:8*8=64) = (h:0x0340|d:832)
			res = finalCalc(startAddr, iChannel, numOfBits)
		} else {
			return 0, fmt.Errorf("неверно введен канал: %q", channel)
		}
	default:
		return 0, fmt.Errorf("неверно введена группа: %q", group)
	}
	return
}

func (b *BasisPV) readHoldingRegister(group, channel string) (res int, err error) {
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
		"specD":      0x0110,
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
		"P": 16, "HI": 4, "HP": 16, "HB": 8,
	}

	startAddr := 0
	// задаем стартовый канал; в случае времени, основного
	// и внешнего контуров - возвращаем значение сразу
	switch group {
	case "CLM":
		if addr, ok := controlLoop[channel]; ok {
			return addr, nil
		}
	case "CLEXT":
		// Во внешнем контуре нет режима работы и спец. алгоритмов
		if slices.Contains([]string{"ctrlMod", "specq", "specD"}, channel) {
			return 0, fmt.Errorf("во внешнем контуре нет параметра %q", channel)
		}
		if addr, ok := controlLoop[channel]; ok {
			shift := 0x1000 // смещение по адресу для внешнего контура
			addr += shift
			return addr, nil
		}

	case "P":
		startAddr = 0x8000
	case "HI":
		startAddr = 0x9000
	case "HP":
		startAddr = 0x9020
	case "HB":
		startAddr = 0x9040
	case "TIME":
		if addr, ok := deviceTime[channel]; ok {
			return addr, nil
		} else {
			return 0, fmt.Errorf("неверно введен параметр: %q", channel)
		}
	default:
		return 0, fmt.Errorf("неверно введена группа: %q", group)
	}

	// преобразуем строку channel в int
	// проверяем на ошибку и на то, что каналы не отрицательные
	if iChannel, e := strconv.Atoi(channel); e == nil && iChannel > 0 {
		// проверяем на наличие группы и размер группы
		if size, ok := channels[group]; ok && size >= iChannel {
			numOfWords := 2 // количество слов (для интервала между адресами)
			res = finalCalc(startAddr, iChannel, numOfWords)
		} else {
			return 0, fmt.Errorf("неверно введен канал (%q)", channel)
		}
	} else {
		return 0, fmt.Errorf("не удалось преобразовать строку канала (%q) в int или нулевой/отрицательный канал", channel)
	}

	return res, nil
}

func (b *BasisPV) readInputRegister(group, channel string) (res int, err error) {
	// карта с ключем - группа, значения
	channels := map[string]map[string]int{
		"I": {"1": 0x0, "2": 0x2, "3": 0x4, "4": 0x6},
		"P": {"1": 0x200, "2": 0x202, "3": 0x204, "4": 0x206,
			"5": 0x208, "6": 0x20A, "7": 0x20C, "8": 0x20E,
			"9": 0x210, "10": 0x212, "11": 0x214, "12": 0x216,
			"13": 0x218, "14": 0x21A, "15": 0x21C, "16": 0x21E},
		"B": {"1": 0x300, "2": 0x302, "3": 0x304, "4": 0x306,
			"5": 0x308, "6": 0x30A, "7": 0x30C, "8": 0x30E},
	}

	// проверяем наличие группы
	if groupChannels, ok := channels[group]; ok {
		// проверяем наличие канала в группе
		if addr, ok := groupChannels[channel]; ok {
			res = addr
		} else {
			return 0, fmt.Errorf("неверно введен канал: %q", channel)
		}
	} else {
		return 0, fmt.Errorf("неверно введена группа: %q", group)
	}
	return res, nil
}

func (b *BasisPV) writeSingleCoil(group, channel string) (int, error) {
	return 0, fmt.Errorf("пока не реализивано")
}
func (b *BasisPV) writeSingleRegister(group, channel string) (int, error) {
	return 0, fmt.Errorf("пока не реализивано")
}
func (b *BasisPV) writeMultipleRegister(group, channel string) (int, error) {
	return 0, fmt.Errorf("функция не реализована")
}
