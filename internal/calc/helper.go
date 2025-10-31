package calc

import (
	"fmt"
	"strings"
)

// finalCalc вычисляет конечный адрес канала.
//
// Параметры:
//   - startAddr: базовый адрес начала группы.
//   - channel: номер канала (начиная с 1).
//   - interval: шаг адреса между каналами.
//
// Возвращает:
//   - Итоговый вычисленный адрес (int).
func finalCalc(startAddr int, channel int, interval int) (res int) {
	// так как начинаем расчет с 0 адреса, к примеру:
	// для канала AI1.1 адрес будет 0 * interval,
	// а для AI1.2 будет 1 * interval
	channel--
	return startAddr + channel*interval

}

// ParseString преобразует строку в 4 значения:
//   - b  - тип БАЗИС
//   - f  - функция Modbus
//   - gr - группа
//   - ch - канал
func ParseString(query string) (b, f, gr, ch string, err error) {
	parse := strings.Fields(query)
	if len(parse) == 4 {
		return parse[0], parse[1], parse[2], parse[3], nil
	}
	return "", "", "", "", fmt.Errorf("некорректный ввод запроса: %q", query)
}

// CalcAddr (конструкция) расчитывает Modbus адрес устройства в
// зависимости от выбранной функции, группы и канала(названия параметра)
//   - readCoil              = "f1"
//   - readDiscreteInput     = "f2"
//   - readHoldingRegister   = "f3"
//   - readInputRegister     = "f4"
//   - writeSingleCoil       = "f5"
//   - writeSingleRegister   = "f6"
//   - writeMultipleRegister = "f16"
func CalcAddr(b Basis, modbusFunc, group, channel string) (int, error) {

	switch modbusFunc {
	case "f1":
		return b.readCoil(group, channel)
	case "f2":
		return b.readDiscreteInput(group, channel)
	case "f3":
		return b.readHoldingRegister(group, channel)
	case "f4":
		return b.readInputRegister(group, channel)
	case "f5":
		return b.writeSingleCoil(group, channel)
	case "f6":
		return b.writeSingleRegister(group, channel)
	// case "f16":
	// 	return b.writeMultipleRegister(group, channel)
	default:
		return -1, fmt.Errorf("неизвестная функция: %q", modbusFunc)
	}
}

// getGroup выводит группы по запрошенной модбас функции
func GetGroup(b Basis, modbusFunc string) ([]string, error) {
	groups := b.mapGroup()

	list, ok := groups[modbusFunc]
	if !ok {
		return nil, fmt.Errorf("модбас функция %s не найдена", modbusFunc)
	}
	keys := make([]string, 0, len(list))
	for k := range list {
		keys = append(keys, k)
	}

	return keys, nil
}

// getChannel выводит группы по запрошенной модбас функции и группе
func GetChannel(b Basis, modbusFunc, group string) ([]string, error) {
	groups := b.mapGroup()

	groupList, ok := groups[modbusFunc]
	if !ok {
		return nil, fmt.Errorf("модбас функция %s не найдена", modbusFunc)
	}

	if channelList, ok := groupList[group]; ok {
		return channelList, nil
	} else {
		return nil, fmt.Errorf("группа %s не найдена", group)
	}
}
