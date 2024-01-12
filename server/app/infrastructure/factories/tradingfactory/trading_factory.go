package tradingfactory

import (
	"cointrading/app/domain/factories"
	"cointrading/app/domain/valueobjects"
)

func NewTradingFactory(exchange *valueobjects.Exchange) factories.ITradingFactory {
	switch {
	case exchange.IsBitflyer():
		return &bitflyerFactory{}
	default:
		return &bitflyerFactory{}
	}
}
