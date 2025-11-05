package calc

import "strconv"

func (b *Basis21) mapGroup() map[string]map[string][]string {

	ch8 := []string{
		"1", "2", "3", "4",
		"5", "6", "7", "8",
	}
	ch10 := []string{
		"1", "2", "3", "4",
		"5", "6", "7", "8",
		"9", "10",
	}
	ch12 := []string{
		"1", "2", "3", "4",
		"5", "6", "7", "8",
		"9", "10", "11", "12",
	}
	ch16 := []string{
		"1", "2", "3", "4",
		"5", "6", "7", "8",
		"9", "10", "11", "12",
		"13", "14", "15", "16",
	}
	ch20 := []string{
		"1", "2", "3", "4", "5",
		"6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15",
		"16", "17", "18", "19", "20",
	}
	ch24 := []string{
		"1", "2", "3", "4",
		"5", "6", "7", "8",
		"9", "10", "11", "12",
		"13", "14", "15", "16",
		"17", "18", "19", "20",
		"21", "22", "23", "24",
	}
	paramCL := []string{
		"ctrlMod", "ctrlConfig", "coefGroup",
		"specAlgNum", "setpoint", "valveValue",
		"Ko", "Kp", "Ti", "Td", "Tf",
		"specKo", "specKp", "specTi",
		"specD1", "specD2",
	}

	res := map[string]map[string][]string{
		"f1": {},
		"f2": {
			"I1": ch16, "I2": ch16, "I3": ch16,
			"I4":  ch8,
			"I11": ch12, "I12": ch12, "I13": ch12,
			"I14": ch12, "I15": ch12, "I21": ch12,
			"I22": ch12, "I23": ch12, "I31": ch12,
			"I32": ch12, "I33": ch12, "P": ch24,
			"W1": {
				"1", "2", "3", "4", "5",
			},
			"W2": ch10, "W3": ch10, "W4": ch10,
			"W5": ch20, "W6": ch20, "W7": ch20,
			"W8": ch20, "W9": ch20,
		},
		"f3": {
			"CL1": paramCL, "CL2": paramCL,
			"CL3": paramCL, "CL4": paramCL,
			"CL5": paramCL, "CL6": paramCL,
			"CL7": paramCL, "CL8": paramCL,
			"TIME": {
				"YEAR", "MONTH",
				"DAY", "HOUR",
				"MIN", "SEC",
			},
			"P": ch24,
		},
		"f4": {
			"I1": ch16, "I2": ch16, "I3": ch16,
			"I4":  ch8,
			"I11": ch12, "I12": ch12, "I13": ch12,
			"I14": ch12, "I15": ch12, "I21": ch12,
			"I22": ch12, "I23": ch12, "I31": ch12,
			"I32": ch12, "I33": ch12,
		},
		"f5":  {},
		"f6":  {},
		"f16": {},
	}
	// 128 канала W (Modbas f2)
	for i := 1; i <= 128; i++ {
		res["f2"]["B"] = append(res["f2"]["B"], strconv.Itoa(i))
	}

	return res
}
