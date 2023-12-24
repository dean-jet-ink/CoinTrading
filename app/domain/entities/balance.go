package entities

import "cointrading/app/domain/valueobjects"

type Balance struct {
	currencyCode *valueobjects.CurrencyCode
	amount       float64
	available    float64
}

func NewBalance(currencyCode int, amount, available float64) (*Balance, error) {
	c, err := valueobjects.NewCurrencyCode(currencyCode)
	if err != nil {
		return nil, err
	}

	return &Balance{
		currencyCode: c,
		amount:       amount,
		available:    available,
	}, nil
}

func (b *Balance) CurrencyCode() *valueobjects.CurrencyCode {
	return b.currencyCode
}

func (b *Balance) Amount() float64 {
	return b.amount
}

func (b *Balance) Available() float64 {
	return b.available
}
