package calc

import "fmt"

type Basis interface {
	readCoil(group, channel string) (int, error)
	readDiscreteInput(group, channel string) (int, error)
	readHoldingRegister(group, channel string) (int, error)
	readInputRegister(group, channel string) (int, error)
	writeSingleCoil(group, channel string) (int, error)
	writeSingleRegister(group, channel string) (int, error)
	writeMultipleRegister(group, channel string) (int, error)
	// returnGroup(modbusFunc string) ([]string, error)
	// returnChannel(modbusFunc string) ([]string, error)

	mapGroup() map[string]map[string][]string
}

func NewBasis(code string) (Basis, error) {
	switch code {
	case "b12":
		return &Basis12{}, nil
	case "b14":
		return &Basis14{}, nil
	case "b21":
		return &Basis21{}, nil
	case "b100":
		return &Basis100{}, nil
	case "bpv":
		return &BasisPV{}, nil
	default:
		return nil, fmt.Errorf("неизвестное устройство: %q", code)
	}
}
