package calc

import "strconv"

func (b *BasisPV) mapGroup() map[string]map[string][]string {

	ch4 := []string{"1", "2", "3", "4"}
	ch8 := []string{
		"1", "2", "3", "4",
		"5", "6", "7", "8",
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
			"I": ch4,
			"F": {"1", "2"},
			"B": ch8,
			"P": ch16,
		},
		"f3": {
			"CLM": {
				"ctrlMod", "ctrlConfig",
				"coefGroup", "specAlgNum",
				"setpoint", "valveValue",
				"Ko", "Kp", "Ti", "Td", "Tf",
				"specq", "specD",
			},
			"CLEXT": {
				"ctrlConfig", "coefGroup",
				"specAlgNum", "setpoint",
				"valveValue",
				"Ko", "Kp", "Ti", "Td", "Tf",
			},
			"P":  ch16,
			"HI": ch4,
			"HP": ch16,
			"HB": ch8,
		},
		"f4": {
			"I": ch4,
			"P": ch16,
			"B": ch8,
		},
		"f5":  {},
		"f6":  {},
		"f16": {},
	}

	// W для 2 функции
	for i := 1; i <= 64; i++ {
		res["f2"]["W"] = append(res["f2"]["W"], strconv.Itoa(i))
	}

	return res
}
