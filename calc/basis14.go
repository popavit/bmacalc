package calc

import (
	"fmt"
	"slices"
	"strconv"
)

type Basis14 struct{}

func (b *Basis14) readCoil(group, channel string) (int, error) {
	return -1, fmt.Errorf("функция не реализована")
}

func (b *Basis14) readDiscreteInput(group, channel string) (res int, err error) {
	channels := map[string]int{
		"I": 8, "F": 2,
		"DI": 8, "EXT": 12, "B": 8,
		"P": 16, "W": 64,
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
	case "DI":
		startAddr = 0x0200
		// входные каналы на шине расширения
	case "EXT":
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
		return 0, fmt.Errorf("неверно введена группа: %q", group)
	}

	// преобразуем строку канала в int и проверяем
	iChannel, e := strconv.Atoi(channel)
	// проверяем на ошибку и на то, что каналы не отрицательные
	if e != nil || iChannel <= 0 {
		return 0, fmt.Errorf("не удается преобразовать строку канала (%q) в int", channel)
	}

	// проверка наличия канала по всем группам, а после расчет
	switch group {
	case "I", "DI", "EXT", "B", "P", "F", "W":
        if size, ok := channels[group]; ok && size >= iChannel {
            numOfBits := 16 // количество битов (для интервала) между адресами
            res = computeAddress(startAddr, iChannel, numOfBits)
		} else {
			return 0, fmt.Errorf("неверно введен канал: %q", channel)
		}
	default:
		return 0, fmt.Errorf("неверно введена группа: %q", group)
	}
	return res, nil
}

func (b *Basis14) readHoldingRegister(group, channel string) (res int, err error) {
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
		"P": 16, "HI": 8, "HEXT": 8, "HP": 16, "HB": 8,
	}

	startAddr := 0
	// задаем стартовый канал; в случае времени, основного
	// и внешнего контуров - возвращаем значение сразу
	switch group {
	case "CLM":
		if addr, ok := controlLoop[channel]; ok {
			return addr, nil
		} else {
			return 0, fmt.Errorf("неверно введен параметр: %q", channel)
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
		} else {
			return 0, fmt.Errorf("неверно введен параметр: %q", channel)
		}
	case "P":
		startAddr = 0x8000
	case "HI":
		startAddr = 0x9000
	case "HEXT":
		startAddr = 0x9010
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
            return computeAddress(startAddr, iChannel, numOfWords), nil
		} else {
			return 0, fmt.Errorf("неверно введен канал: %q", channel)
		}
	} else {
		return 0, fmt.Errorf("не удалось преобразовать строку канала (%q) в int или нулевой/отрицательный канал", channel)
	}

}

func (b *Basis14) readInputRegister(group, channel string) (res int, err error) {
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
	default:
		return 0, fmt.Errorf("неверно введена группа: %q", group)
	}

	// преобразуем строку канала в int
	iChannel, e := strconv.Atoi(channel)
	// проверяем на ошибку и на то, что каналы не отрицательные
	if e != nil || iChannel <= 0 {
		return 0, fmt.Errorf("не удается преобразовать строку канала (%q) в int или канал нулевой/отрицательный", channel)
	}

	// проверка наличия канала по группе в карте channels,
	// не выходит ли за пределы каналов и, после, расчет
	if size, ok := channels[group]; ok && size >= iChannel {
        numOfWords := 2 // количество слов (для интервала между адресами)
        res = computeAddress(startAddr, iChannel, numOfWords)
	} else {
		return 0, fmt.Errorf("неверно введен канал: %q", channel)
	}

	return res, nil
}

func (b *Basis14) writeSingleCoil(group, channel string) (int, error) {
	return -1, fmt.Errorf("функция не реализована")
}

func (b *Basis14) writeSingleRegister(group, channel string) (int, error) {
	return -1, fmt.Errorf("функция не реализована")
}

func (b *Basis14) writeMultipleRegister(group, channel string) (int, error) {
	return -1, fmt.Errorf("функция не реализована")
}
