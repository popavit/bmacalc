package calc

import (
	"fmt"
	"strconv"
)

func (b *Basis100) mapGroup() map[string]map[string][]string {

	res := map[string]map[string][]string{
		"f1": {},
		"f2": {},
		"f3": {
			"TIME": {
				"YEAR", "MONTH",
				"DAY", "HOUR",
				"MIN", "SEC"},
			"STATE": {"sysError", "dublication"},
			"TBL":   {"CURRENT"},
		},
		"f4":  {},
		"f5":  {},
		"f6":  {},
		"f16": {},
	}

	// добавляем в f2 группы и каналы
	for _, group := range []string{"AI", "AO", "DI", "DO"} {
		groupCount := 40 // количество групп
		for i := 1; i <= groupCount; i++ {
			channels := []string{}
			switch group {
			case "AI":
				fallthrough
			case "AO":
				channels = []string{
					"1", "2", "3", "4",
					"5", "6", "7", "8",
				}
			case "DI":
				channels = []string{
					"1", "2", "3", "4",
					"5", "6", "7", "8",
					"9", "10", "11", "12",
					"13", "14", "15", "16",
				}
			case "DO":
				channels = []string{
					"1", "2", "3", "4",
					"5", "6", "7", "8",
					"9", "10",
				}
			}
			res["f2"][fmt.Sprintf("%s%d", group, i)] = channels
		}
	}

	for _, group := range []string{"CL", "AO"} {
		items := []string{}
		numOfGroups := 0
		switch group {
		case "CL":
			items = []string{
				"ctrlMod", "ctrlConfig", "coefGroup",
				"specAlgNum", "setpoint", "valveValue",
				"Ko", "Kp", "Ti", "Td", "Tf",
				"specKo", "specKp", "specTi",
				"specD1", "specD2", "valveClosingValue",
			}
			numOfGroups = 100
		case "AO":
			items = []string{
				"1", "2", "3", "4",
				"5", "6", "7", "8",
			}
			numOfGroups = 40
		}

		for i := 1; i <= numOfGroups; i++ {
			res["f3"][fmt.Sprintf("%s%d", group, i)] = items
		}
	}
	//256 таблиц в f3
	for i := 1; i <= 256; i++ {
		res["f3"]["TBL"] = append(res["f3"]["TBL"], strconv.Itoa(i))
	}

	return res
}
