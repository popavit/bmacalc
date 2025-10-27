package main

func finalCalc(startAddr int, channel int, interval int) (res int) {

	channel-- // так как начинаем расчет с 0 адреса
	return startAddr + channel*interval

}
