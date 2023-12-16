package valueobject

var (
	btcJPY = NewSymbol(1)
	btcUSD = NewSymbol(2)
	ethJPY = NewSymbol(3)
	ethUSD = NewSymbol(4)
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

func NewSymbol(value int) *Symbol {
	return &Symbol{
		value: value,
	}
}

func (s *Symbol) Value() int {
	return s.value
}

func (s *Symbol) String() string {
	name, ok := symbolNames[*s]

	if !ok {
		return "Unknown"
	}

	return name
}

func (s *Symbol) ToList() []*Symbol {
	return []*Symbol{
		btcJPY,
		btcUSD,
		ethJPY,
		ethUSD,
	}
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
