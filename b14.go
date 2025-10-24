package main

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
	}

	// проверка наличия канала по всем группам, а после расчет
	switch group {
	case "AI", "DI", "EI", "B", "P", "F1", "F2", "W":
		if size, ok := channels[group]; ok && size >= channel {
			numOfBits := 16 // количество битов (для интервала) между адресами
			channel--       // так как начинаем расчет с 0 адреса
			// пример: группа I4 - startAddr = (h:0x0300|d:768)
			// канал 8; получаем расчет:
			// (0x0300|d:768) + (h:0x0040|d:8*8=64) = (h:0x0340|d:832)
			res = startAddr + channel*numOfBits
		}
	}
	return
}

func calc_addr_b14_f3()  {}
func calc_addr_b14_f4()  {}
func calc_addr_b14_f5()  {}
func calc_addr_b14_f6()  {}
func calc_addr_b14_f16() {}
