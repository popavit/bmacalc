package calc

import "strconv"

func (b *Basis14) mapGroup() map[string]map[string][]string {

	ch8 := []string{
		"1", "2", "3", "4",
		"5", "6", "7", "8",
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

	res := map[string]map[string][]string{
		"f1": {},
		"f2": {
			"I":   ch8,
			"F":   {"1", "2"},
			"DI":  ch8,
			"EXT": ch12,
			"B":   ch8,
			"P":   ch16,
			"W":   {},
		},
		"f3": {
			"CLM": {
				"ctrlMod", "ctrlConfig", "coefGroup",
				"specAlgNum", "setpoint", "valveValue",
				"Ko", "Kp", "Ti", "Td", "Tf",
				"specq", "specD",
			},
			"CLEXT": {
				"ctrlConfig", "coefGroup",
				"specAlgNum", "setpoint", "valveValue",
				"Ko", "Kp", "Ti", "Td", "Tf",
			},
			"TIME": {
				"YEAR", "MONTH",
				"DAY", "HOUR",
				"MIN", "SEC",
			},
			"P":    ch16,
			"HI":   ch8,
			"HEXT": ch8,
			"HP":   ch16,
			"HB":   ch8,
		},
		"f4": {
			"I":   ch8,
			"EXT": ch8,
			"P":   ch16,
			"B":   ch8,
		},
		"f5":  {},
		"f6":  {},
		"f16": {},
	}

	// 64 канала W
	for i := 1; i <= 64; i++ {
		temp := res["f2"]
		temp["W"] = append(temp["W"], strconv.Itoa(i))
	}

	return res

}
