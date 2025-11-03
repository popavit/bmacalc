package calc

import (
	"fmt"
	"slices"
	"strconv"
)

type Basis12 struct{}

func (b *Basis12) readCoil(group, channel string) (int, error) {
	return 0, fmt.Errorf("пока не реализовано")
}

func (b *Basis12) readDiscreteInput(group, channel string) (res int, err error) {

	// карта с ключем - группа, значения
	channels := map[string]map[string]int{
		"I": {"2": 0x0008, "3": 0x0010, "4": 0x0018, "13": 0x0060},
		"P": {"1": 0x00C0, "2": 0x00C8, "3": 0x00D0},
		"W": {"1": 0x00F0, "2": 0x00F8, "3": 0x0100,
			"4": 0x0108, "5": 0x0110, "6": 0x0118,
			"7": 0x0120, "8": 0x0128, "3B": 0x0130},
	}

	return addrFromGroupMap(channels, group, channel)
}

func (b *Basis12) readHoldingRegister(group, channel string) (res int, err error) {
	// список адресов для контура регулирования
	controlLoop := map[string]int{
		"ctrlMod":    0x4000,
		"ctrlConfig": 0x4001,
		"coefGroup":  0x4002,
		"specAlgNum": 0x4003,
		"setpoint":   0x4100,
		"valveValue": 0x4102,
		"Ko":         0x4104,
		"Kp":         0x4106,
		"Ti":         0x4108,
		"Td":         0x410A,
		"Tf":         0x410C,
		"specKo":     0x410E,
		"specKp":     0x4110,
		"specTi":     0x4112,
		"specD1":     0x4114,
		"specD2":     0x4116,
	}

	// список адресов для чтения адресов
	deviceTime := map[string]int{
		"YEAR": 0x3000, "MONTH": 0x3001,
		"DAY": 0x3002, "HOUR": 0x3003,
		"MIN": 0x3004, "SEC": 0x3005,
	}

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
		// во внешнем контуре нет режима работы и спец. алгоритмов
		if slices.Contains([]string{"ctrlMod", "specD1",
			"specD2", "specKo", "specKp",
			"specTi", "Td"}, channel) {
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
		// преобразуем channel в int
		if iChannel, e := strconv.Atoi(channel); e == nil {
			// проверяем диапазон
			if iChannel > 0 && iChannel <= 3 {
				startAddr := 0x8000
				numOfWords := 2
				return computeAddress(startAddr, iChannel, numOfWords), nil
			} else {
				return 0, fmt.Errorf("канал вне диапазона: %q", channel)
			}
		} else {
			return 0, fmt.Errorf("не удается преобразовать строку канала (%q) в int", channel)
		}
	case "TIME":
		if addr, ok := deviceTime[channel]; ok {
			return addr, nil
		} else {
			return 0, fmt.Errorf("неверно введен параметр: %q", channel)
		}
	default:
		return 0, fmt.Errorf("неверно введена группа: %q", group)
	}

}

func (b *Basis12) readInputRegister(group, channel string) (res int, err error) {

	// карта с ключем - группа, значения
	channels := map[string]map[string]int{
		"I": {"2": 0x0002, "3": 0x004, "4": 0x006, "13": 0x0018},
		"P": {"1": 0x0030, "2": 0x0032, "3": 0x0034},
	}

	return addrFromGroupMap(channels, group, channel)
}

func (b *Basis12) writeSingleCoil(group, channel string) (int, error) {
	return 0, fmt.Errorf("пока не реализовано")
}
func (b *Basis12) writeSingleRegister(group, channel string) (int, error) {
	return 0, fmt.Errorf("пока не реализовано")
}

func (b *Basis12) writeMultipleRegister(group, channel string) (int, error) {
	return 0, fmt.Errorf("функция не реализована")
}
