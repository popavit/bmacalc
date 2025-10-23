package main

import (
	"slices"
)

func calc_addr_b21_f2(group string, channel int) (res int) {

	res = -1

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

	orderChannelsI := []string{
		"I1", "I2", "I3", "I4", "I11",
		"I12", "I14", "I15", "I21", "I22",
		"I23", "I31", "I32", "I33",
	}

	orderChannelsW := []string{
		"W1", "W2", "W3",
		"W4", "W5", "W6", "W7", "W8",
		"W9",
	}

	startAddr := 0
	switch group {
	case "I1", "I2", "I3", "I4", "I11",
		"I12", "I14", "I15", "I21", "I22",
		"I23", "I31", "I32", "I33":
		if i := slices.Index(orderChannelsI, group); i != -1 {
			startAddr = i * 0x0100
		}
	case "P":
		startAddr = 0x2000
	case "IB", "OB":
		startAddr = 0x3000
	case "W1", "W2", "W3", "W4", "W5",
		"W6", "W7", "W8", "W9":
		if i := slices.Index(orderChannelsW, group); i != -1 {
			startAddr = 0x1000 + i*0x0100
		}
	}
	if size, ok := channels[group]; ok && size <= channels[group] {
		res = startAddr + (channel-1)*8
	}

	return
}

func calc_addr_b21_f4(group string, channel int) (res int) {

	res = -1

	orderChannelsI := []string{
		"I1", "I2", "I3", "I4",
		"I11", "I12", "I13", "I14", "I15",
		"I21", "I22", "I23",
		"I31", "I32", "I33",
		"P", "IB", "OB", "V1", "V2",
	}

	startAddr := 0
	switch group {
	case "I1", "I2", "I3", "I4",
		"I11", "I12", "I13", "I14", "I15",
		"I21", "I22", "I23",
		"I31", "I32", "I33":
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
	if ok := slices.Index(orderChannelsI, group); ok != -1 {
		res = startAddr + (channel-1)*2
	}
	return
}
