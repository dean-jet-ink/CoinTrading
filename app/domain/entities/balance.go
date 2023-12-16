package entities

import "cointrading/app/domain/valueobject"

type Balance struct {
	currencyCode *valueobject.CurrencyCode
	amount       float64
	available    float64
}

func NewBalance(currencyCode string, amount, available float64) *Balance {
	return &Balance{
		currencyCode: valueobject.NewCurrencyCode(currencyCode),
		amount:       amount,
		available:    available,
	}
}

func (b *Balance) CurrencyCode() *valueobject.CurrencyCode {
	return b.currencyCode
}

func (b *Balance) Amount() float64 {
	return b.amount
}

func (b *Balance) Available() float64 {
	return b.available
}
