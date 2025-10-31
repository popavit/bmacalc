package calc

import (
	"fmt"
	"strconv"
)

type Basis100 struct{}

func (b *Basis100) mapGroup() map[string]map[string][]string {

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

func (b *Basis100) readCoil(group, channel string) (int, error) {
	return 0, fmt.Errorf("пока не реализивано")
}

func (b *Basis100) readDiscreteInput(group, channel string) (int, error) {

	// вычисляем стартовый в зависимости от типа группы
	startAddr := 0
	switch group[:2] {
	case "DI":
		startAddr = 0
	case "AI":
		startAddr = 0x4000
	case "DO":
		startAddr = 0x8000
	case "AO":
		startAddr = 0xC000
	default:
		return 0, fmt.Errorf("неверно указан тип группы: %q", group[:2])
	}

	channels := map[string]int{"DI": 16, "AI": 8, "DO": 10, "AO": 8}
	// переводим номер группы в int и проверяем на ошибку
	if groupNum, err := strconv.Atoi(group[2:]); err == nil {
		//если ошибки нет и указана группа с 1 по 40
		if groupNum > 0 && groupNum <= 40 {
			// переводим строку канала в int
			if iChannel, err := strconv.Atoi(channel); err == nil {
				numOfChannels := channels[group[:2]] // берем количество каналов
				// проверяем канал, должен быть с 1 по 16
				if iChannel > 0 && iChannel <= numOfChannels {
					intervalBetweenGroup := 0x100 // интервал между группами
					numOfBits := 8                // количество битов
					groupNum--                    // смещаем на -1, так как расчет адреса 1, начинается с 0x0000
					// возвращаем адрес
					return finalCalc(startAddr+groupNum*intervalBetweenGroup, iChannel, numOfBits), nil
				} else {
					return 0, fmt.Errorf("неверно введен канал: %q", channel)
				}
			} else {
				return 0, fmt.Errorf("не удается преобразовать строку канала в int: %q", channel)

			}
		} else {
			return 0, fmt.Errorf("группы не существует: %q", group)
		}

	} else {
		return 0, fmt.Errorf("неверно указан номер или некорректная строка группы: %q", group)
	}
}

func (b *Basis100) readHoldingRegister(group, channel string) (res int, err error) {
	// список адресов для контура регулирования
	controlLoop := map[string]int{
		"ctrlMod":           0x0000,
		"ctrlConfig":        0x0001,
		"coefGroup":         0x0002,
		"specAlgNum":        0x0003,
		"setpoint":          0x0100,
		"valveValue":        0x0102,
		"Ko":                0x0104,
		"Kp":                0x0106,
		"Ti":                0x0108,
		"Td":                0x010A,
		"Tf":                0x010C,
		"specKo":            0x010E,
		"specKp":            0x0110,
		"specTi":            0x0112,
		"specD1":            0x0114,
		"specD2":            0x0116,
		"valveClosingValue": 0x0118,
	}

	// список адресов для чтения адресов
	deviceTime := map[string]int{
		"YEAR": 0xFF00, "MONTH": 0xFF01,
		"DAY": 0xFF02, "HOUR": 0xFF03,
		"MIN": 0xFF04, "SEC": 0xFF05,
	}

	// вычисляем стартовый
	startAddr := 0
	switch group[:2] {
	case "CL":
		startAddr = 0
		numOfCLs := 100 // количество контуров регулирования
		if numCL, err := strconv.Atoi(group[2:]); err == nil && numCL > 0 && numCL <= numOfCLs {
			intervalBetweenCL := 0x200 // интервал между контурами регулирования
			numCL--                    // смещаем на -1, так как расчет адреса 1, начинается с 0x0000
			// вычисляем стартовый адрес конкретного контура
			CLAddr := startAddr + intervalBetweenCL*numCL
			// если параметр существует в списке
			if parametrAddr, ok := controlLoop[channel]; ok {
				// возвращаем адрес
				return CLAddr + parametrAddr, nil
			} else {
				return 0, fmt.Errorf("неверно введен параметр контура: %q", channel)
			}
		} else {
			return 0, fmt.Errorf("неверно введен номер контура или неудалось обработать строку: %q", group)
		}
	case "AO":
		startAddr = 0xD000
		numOfGroups := 40 // количество групп
		// переводим в int и проверяем номер группы на ошибку и на существование (c 1 по 40)
		if groupNum, err := strconv.Atoi(group[2:]); err == nil && groupNum > 0 && groupNum <= numOfGroups {
			numOfChannels := 8 // количество каналов
			groupNum--         // смещаем на -1, так как расчет адреса 1, начинается с 0x0000
			// переводим в int и проверяем номер канала на ошибку и на существование
			if iChannel, err := strconv.Atoi(channel); err == nil && iChannel > 0 && iChannel <= numOfChannels {
				intervalBetweenGroup := 0x10 // интервал между группами
				// вычисляем адрес конкретной группы
				groupAddr := startAddr + intervalBetweenGroup*groupNum
				numOfWords := 2 // количество слов для интервала
				return finalCalc(groupAddr, iChannel, numOfWords), nil
			} else {
				return 0, fmt.Errorf("неверно введен канал: %q", channel)
			}
		} else {
			return 0, fmt.Errorf("неверно введена номер группы или неудалось обработать строку: %q", group)
		}
	default:
		switch group {
		case "TBL":
			// Если выбрана текущая, возвращаем 0xE002
			if channel == "CURRENT" {
				return 0xE002, nil
			}
			startAddr = 0xEE00
			numOfTbl := 256 // количество таблиц
			// переводим в int и проверяем номер таблицы на ошибку и на существование
			if tblNum, err := strconv.Atoi(channel); err == nil && tblNum > 0 && tblNum <= numOfTbl {
				numOfWords := 2 // количество слов для интервала
				return finalCalc(startAddr, tblNum, numOfWords), nil
			} else {
				return 0, fmt.Errorf("неверно введен номер таблицы: %q", channel)
			}
		case "TIME":
			startAddr = 0xFF00
			if addr, ok := deviceTime[channel]; ok {
				return addr, nil
			} else {
				return 0, fmt.Errorf("неверно введен параметр:%q", channel)
			}
		case "STATE":
			switch channel {
			case "sysError":
				return 0xFF11, nil
			case "dublication":
				return 0xFF12, nil
			default:
				return 0, fmt.Errorf("неверно введен параметр: %q", channel)
			}
		default:
			return 0, fmt.Errorf("неверно введена группа (CLx/AOx/TIME/TBL/STATE): %q", group)
		}
	}
}

func (b *Basis100) readInputRegister(group, channel string) (int, error) {

	startAddr := 0
	switch group[:2] {
	case "AI":
		startAddr = 0
	case "AO":
		startAddr = 0x1000
	default:
		return 0, fmt.Errorf("неверно введена группа: %q", group)
	}

	numOfGroups := 40 // количество групп
	// переводим и проверяем номер группы на ошибку и на существование
	if groupNum, err := strconv.Atoi(group[2:]); err == nil && groupNum > 0 && groupNum <= numOfGroups {
		numOfChannels := 8 // количество каналов
		// переводим и проверяем номер канала на ошибку и на существование
		if iChannel, err := strconv.Atoi(channel); err == nil && iChannel > 0 && iChannel <= numOfChannels {
			intervalBetweenGroup := 0x10 // интервал между группами
			groupNum--                   // смещаем на -1, так как расчет адреса 1, начинается с 0x0000
			// вычисляем адрес конкретной группы
			groupAddr := startAddr + intervalBetweenGroup*groupNum
			numOfWords := 2 // количество слов для интервала
			return finalCalc(groupAddr, iChannel, numOfWords), nil
		} else {
			return 0, fmt.Errorf("неверно введен канал: %q", channel)
		}
	} else {
		return 0, fmt.Errorf("не удалось считать номер группы или вне диапазона групп(1..40)%q", group)
	}
}

func (b *Basis100) writeSingleCoil(group, channel string) (int, error) {
	return 0, fmt.Errorf("пока не реализивано")
}
func (b *Basis100) writeSingleRegister(group, channel string) (int, error) {
	return 0, fmt.Errorf("пока не реализивано")
}

func (b *Basis100) writeMultipleRegister(group, channel string) (int, error) {
	return 0, fmt.Errorf("функция не реализована")
}
