package calc

import "fmt"

type Basis interface {
	CalcAddr(funcCode, group, channel string) (int, error)
	readCoil(group, channel string) (int, error)
	readDiscreteInput(group, channel string) (int, error)
	readHoldingRegister(group, channel string) (int, error)
	readInputRegister(group, channel string) (int, error)
	writeSingleCoil(group, channel string) (int, error)
	writeSingleRegister(group, channel string) (int, error)
	writeMultipleRegister(group, channel string) (int, error)
}

func NewBasis(code string) (Basis, error) {
	switch code {
	case "b12":
		return nil, fmt.Errorf("устройство пока не реализовано")
	case "b14":
		return &Basis14{}, nil
	case "b21":
		return &Basis21{}, nil
	case "b100":
		return nil, fmt.Errorf("устройство пока не реализовано")
	case "bpv":
		return nil, fmt.Errorf("устройство пока не реализовано")
	default:
		return nil, fmt.Errorf("неизвестное устройство: %q", code)
	}
}
