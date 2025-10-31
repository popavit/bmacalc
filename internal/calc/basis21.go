package calc

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

type Basis21 struct{}

func (b *Basis21) mapGroup() map[string]map[string][]string {

	return map[string]map[string][]string{
		"f1":  {},
		"f2":  {},
		"f3":  {},
		"f4":  {},
		"f5":  {},
		"f6":  {},
		"f16": {},
	}

}

// CalcAddr расчитывает Modbus адрес устройства в
// зависимости от выбранной функции, группы и канала(названия параметра)
//   - readCoil              = "f1"
//   - readDiscreteInput     = "f2"
//   - readHoldingRegister   = "f3"
//   - readInputRegister     = "f4"
//   - writeSingleCoil       = "f5"
//   - writeSingleRegister   = "f6"
//   - writeMultipleRegister = "f16"

func (b *Basis21) readCoil(group, channel string) (int, error) {
	return 0, fmt.Errorf("функция не реализована")
}

func (b *Basis21) readDiscreteInput(group, channel string) (res int, err error) {

	// список каналов с: ключ - группа,
	// значение - колличество каналов в группе.
	channels := map[string]int{
		"I1": 16, "I2": 16, "I3": 16,
		"I4":  8,
		"I11": 12, "I12": 12, "I13": 12, "I14": 12, "I15": 12,
		"I21": 12, "I22": 12, "I23": 12,
		"I31": 12, "I32": 12, "I33": 12,
		"P": 24, "B": 128,
		"W1": 5,
		"W2": 10, "W3": 10, "W4": 10,
		"W5": 20, "W6": 20, "W7": 20, "W8": 20, "W9": 20,
	}

	// упорядочиваем список для расчетов (используются индексы)
	orderChannelsI := []string{
		"I1", "I2", "I3", "I4", "I11",
		"I12", "I13", "I14", "I15", "I21", "I22",
		"I23", "I31", "I32", "I33",
	}

	// упорядочиваем список для расчетов (используются индексы)
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
	case "B":
		startAddr = 0x3000
	case "W1", "W2", "W3", "W4", "W5",
		"W6", "W7", "W8", "W9":
		// по индексу из списка выше расчитываем начальный адрес для группы
		if i := slices.Index(orderChannelsW, group); i != -1 {
			interval := 0x0100 // интервал адресов между группами
			startAddr = 0x1000 + i*interval
		}
	default:
		return 0, fmt.Errorf("неверно введена группа: %q", group)
	}

	// преобразуем строку channels в int
	iChannel, err := strconv.Atoi(channel)
	// проверяем на ошибку и, что канал не отрицательный
	if err != nil || iChannel <= 0 {
		return 0, fmt.Errorf("не удается преобразовать строку канала (%q) в int или нулевое/отрицательное значение", channel)
	}
	// проверка наличия канала по группе в карте channels, а после расчет
	if size, ok := channels[group]; ok && size >= iChannel && iChannel > 0 {
		numOfBits := 8 // количество битов (для интервала) между адресами
		// пример: группа I4 - startAddr = (h:0x0300|d:768)
		// канал 8; получаем расчет:
		// (0x0300|d:768) + (h:0x0040|d:8*8=64) = (h:0x0340|d:832)
		res = finalCalc(startAddr, iChannel, numOfBits)
	} else {
		return 0, fmt.Errorf("неверно введены группа и/или канал")
	}

	return res, nil
}

func (b *Basis21) readHoldingRegister(group, channel string) (res int, err error) {

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
		"MIN": 0xFF04, "SEC": 0xFF05,
	}

	switch group {
	case "CL1", "CL2", "CL3", "CL4",
		"CL5", "CL6", "CL7", "CL8":
		if i := slices.Index(controlLoopList, group); i != 1 {
			// по выбраному контуру регулирования, параметру
			// и промежутку между адресами считаем адрес
			interval := 0x1000
			if addr, ok := controlLoop[channel]; ok {
				res = addr + i*interval
			} else {
				return 0, fmt.Errorf("неверно введен канал: %q", channel)
			}
		}
	case "TIME":
		if addr, ok := deviceTime[channel]; ok {
			res = addr
		} else {
			return 0, fmt.Errorf("неверно введен канал: %q", channel)
		}
	case "P":
		if iChannel, e := strconv.Atoi(channel); e == nil {
			// расчетные каналы с 1 по 24
			if iChannel > 0 && iChannel <= 24 {
				startAddr := 0x8000
				numOfWords := 2
				return finalCalc(startAddr, iChannel, numOfWords), nil
			} else {
				return 0, fmt.Errorf("канал %q вне диапазона", channel)
			}
		} else {
			return 0, fmt.Errorf("не удается преобразовать строку канала (%q) в int", channel)
		}
	default:
		return 0, fmt.Errorf("неверно введена группа: %q", group)
	}

	return res, nil
}

func (b *Basis21) readInputRegister(group, channel string) (res int, err error) {

	// список каналов
	// ключ - группа, значение - колличество каналов в группе.
	channels := map[string]int{
		"I1": 16, "I2": 16, "I3": 16, "I4": 8,
		"I11": 12, "I12": 12, "I13": 12, "I14": 12,
		"I15": 12, "I21": 12, "I22": 12, "I23": 12,
		"I31": 12, "I32": 12, "I33": 12,
		"P": 24, "B": 128, "V1": 8, "V2": 8,
	}

	// список для расчетов (используются индексы)
	orderChannelsI := []string{
		"I1", "I2", "I3", "I4",
		"I11", "I12", "I13", "I14",
		"I15", "I21", "I22", "I23",
		"I31", "I32", "I33",
		"P", "B", "V1", "V2",
	}

	// карта с каналами исторических (суточных)
	channelsHI := map[string]int{
		"HI1": 8, "HI2": 8, "HI3": 8,
		"HI11": 8, "HI12": 8, "HI13": 8,
		"HI14": 8, "HI15": 8, "HP": 24,
	}
	// список для расчетов (используются индексы)
	orderChannelsHI := []string{
		"HI1", "HI2", "HI3",
		"HI11", "HI12", "HI13",
		"HI14", "HI15", "HP"}

	startAddr := 0
	switch group {
	case "I1", "I2", "I3", "I4",
		"I11", "I12", "I13", "I14",
		"I15", "I21", "I22", "I23",
		"I31", "I32", "I33":
		// по индексу из списка выше расчитываем начальный адрес для группы
		if i := slices.Index(orderChannelsI, group); i != -1 {
			interval := 0x0020 // интервал между группами
			startAddr = i * interval
		} else {
			return 0, fmt.Errorf("неверно введена группа: %q", group)
		}
	case "V1":
		startAddr = 0x1000
	case "V2":
		startAddr = 0x1010
	case "P":
		startAddr = 0x2000
	case "B":
		startAddr = 0x3000
	case "HI1", "HI2", "HI3",
		"HI11", "HI12", "HI13",
		"HI14", "HI15", "HP":
		if i := slices.Index(orderChannelsHI, group); i != -1 {
			startAddr = 0xA000 // HI1
			interval := 0x0480 // интервал между группами
			startAddr += i * interval
		} else {
			return 0, fmt.Errorf("неверно введена группа: %q", group)
		}
	default:
		return 0, fmt.Errorf("неверно введена группа: %q", group)
	}

	// перевод в int, проверка наличия канала по группе в списке каналов, а после расчет
	if iChannel, e := strconv.Atoi(channel); e == nil {
		if size, ok := channels[group]; ok && size >= iChannel && iChannel > 0 {
			numOfWords := 2 // количество слов (для интервала между адресами)
			res = finalCalc(startAddr, iChannel, numOfWords)
			// если нет, но есть в списке исторических каналов
		} else {
			return 0, fmt.Errorf("неверно введен канал: %q", channel)
		}

	} else if size, ok := channelsHI[group]; ok { // если данные часовые

		// вычисляем канал, день и час
		re := regexp.MustCompile(`^(\d+)([tyb])(\d{1,2})$`)
		matches := re.FindStringSubmatch(channel)

		if len(matches) == 4 {
			iChannel, err = strconv.Atoi(matches[1])
			if err != nil {
				return 0, fmt.Errorf("в строке параметра (%q) неверно указан канал", channel)
			}

			hour, e := strconv.Atoi(matches[3])
			if e != nil {
				return 0, fmt.Errorf("в строке параметра (%q) неверно указаны часы", channel)
			}

			day := matches[2]
			// в зависимости от дня добавляем к startAddr
			switch day {
			case "t":
				startAddr += 0
			case "y":
				startAddr += 0x0030
			case "b":
				startAddr += 0x0060
			default:
				return 0, fmt.Errorf("в строке параметра (%q) неверно указан день", channel)
			}

			// если часы
			if size >= iChannel && iChannel > 0 && hour >= 0 && hour < 24 {
				res = finalCalc(startAddr, iChannel, 0x0090) + hour*2
			} else {
				return 0, fmt.Errorf("неверно указан параметр: %q", channel)
			}
		} else {
			return 0, fmt.Errorf("неверно указан параметр: %q", channel)
		}

	} else {
		return 0, fmt.Errorf("не удалось вычислить канал (%q) или неверно указана группа (%q)", channel, group)
	}
	return
}

func (b *Basis21) writeSingleCoil(group, channel string) (int, error) {
	return -1, fmt.Errorf("функция не реализована")
}

func (b *Basis21) writeSingleRegister(group, channel string) (int, error) {
	return -1, fmt.Errorf("функция не реализована")
}

func (b *Basis21) writeMultipleRegister(group, channel string) (int, error) {
	return -1, fmt.Errorf("функция не реализована")
}
