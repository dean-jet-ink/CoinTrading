package valueobjects

import (
	"cointrading/app/domain/myerror"
	"fmt"
)

var (
	BTCJPY = &Symbol{1}
	ETHJPY = &Symbol{2}
)

var symbolNames = map[Symbol]string{
	*BTCJPY: "BTC/JPY",
	*ETHJPY: "ETH/JPY",
}

var symbolTableNames = map[Symbol]string{
	*BTCJPY: "btc_jpy",
	*ETHJPY: "eth_jpy",
}

type Symbol struct {
	value int
}

func NewSymbol(value int) (*Symbol, error) {
	symbol := &Symbol{value}

	if _, ok := symbolNames[*symbol]; !ok {
		err := fmt.Errorf("%w: Unexpected symbol code %v", myerror.ErrUnexpectedSymbol, value)

		return nil, err
	}

	return symbol, nil
}

func (s *Symbol) Value() int {
	return s.value
}

func (s *Symbol) DisplayValue() string {
	return symbolNames[*s]
}

func (s *Symbol) DisplayValueForTableName() string {
	return symbolTableNames[*s]
}

func (s *Symbol) IsBTCJPY() bool {
	return *s == *BTCJPY
}

func (s *Symbol) IsETHJPY() bool {
	return *s == *ETHJPY
}

func Symbols() []*Symbol {
	return []*Symbol{
		BTCJPY,
		ETHJPY,
	}
}
