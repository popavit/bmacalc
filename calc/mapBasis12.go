package calc

func (b *Basis12) mapGroup() map[string]map[string][]string {

	return map[string]map[string][]string{
		"f1": {},
		"f2": {"I": {"2", "3", "4", "13"},
			"P": {"1", "2", "3"},
			"W": {"1", "2", "3", "4", "5", "6", "7", "8", "3B"}},
		"f3": {
			"CLM": {"ctrlMod",
				"ctrlConfig", "coefGroup", "specAlgNum", "setpoint",
				"valveValue", "Ko", "Kp", "Ti", "Td", "Tf",
				"specKo", "specKp", "specTi", "specD1",
				"specD2"},
			"CLEXT": {"coefGroup", "specAlgNum", "setpoint",
				"valveValue", "Ko", "Kp", "Ti", "Tf"},
			"TIME": {"YEAR", "MONTH",
				"DAY", "HOUR",
				"MIN", "SEC"},
			"P": {"1", "2", "3"},
		},
		"f4": {
			"I":   {"2", "3", "4", "13"},
			"P":   {"1", "2", "3"},
			"f5":  {},
			"f6":  {},
			"f16": {},
		}}
}
