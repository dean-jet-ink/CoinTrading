package valueobjects

import (
	"cointrading/app/domain/errors"
	"fmt"
)

var (
	btcJPY = &Symbol{1}
	ethJPY = &Symbol{2}
)

var symbolNames = map[int]string{
	1: "BTC/JPY",
	3: "ETH/JPY",
}

type Symbol struct {
	value int
}

func NewSymbol(value int) (*Symbol, error) {
	if _, ok := symbolNames[value]; !ok {
		message := "非対応のシンボルです"
		original := fmt.Sprintf("Unexpected symbol code %v", value)

		myerr := errors.NewMyError(message, original, 500)

		return nil, myerr
	}

	return &Symbol{
		value: value,
	}, nil
}

func (s *Symbol) Value() int {
	return s.value
}

func (s *Symbol) DisplayValue() string {
	name, _ := symbolNames[s.value]

	return name
}

func (s *Symbol) IsBTCJPY() bool {
	return *s == *btcJPY
}

func (s *Symbol) IsETHJPY() bool {
	return *s == *ethJPY
}

func Symbols() []*Symbol {
	return []*Symbol{
		btcJPY,
		ethJPY,
	}
}
