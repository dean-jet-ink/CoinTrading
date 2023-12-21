package valueobjects

import (
	"cointrading/app/domain/errors"
	"fmt"
)

var (
	btcJPY = &Symbol{1}
	btcUSD = &Symbol{2}
	ethJPY = &Symbol{3}
	ethUSD = &Symbol{4}
)

var symbolNames = map[Symbol]string{
	*btcJPY: "BTC/JPY",
	*btcUSD: "BTC/USD",
	*ethJPY: "ETH/JPY",
	*ethUSD: "ETH/USD",
}

type Symbol struct {
	value int
}

func NewSymbol(value int) (*Symbol, error) {
	symbol := &Symbol{
		value: value,
	}

	if _, ok := symbolNames[*symbol]; !ok {
		message := fmt.Sprint("非対応のsymbolです")
		original := fmt.Sprintf("Unexpected symbol code %v", value)

		myerr := errors.NewMyError(message, original, 500)

		return nil, myerr
	}

	return symbol, nil
}

func (s *Symbol) Value() int {
	return s.value
}

func (s *Symbol) String() string {
	name, _ := symbolNames[*s]

	return name
}

func (s *Symbol) IsBTCJPY() bool {
	return *s == *btcJPY
}

func (s *Symbol) IsBTCUSD() bool {
	return *s == *btcUSD
}

func (s *Symbol) IsETHJPY() bool {
	return *s == *ethJPY
}

func (s *Symbol) IsETHUSD() bool {
	return *s == *ethUSD
}

func Symbols() []*Symbol {
	return []*Symbol{
		btcJPY,
		btcUSD,
		ethJPY,
		ethUSD,
	}
}
