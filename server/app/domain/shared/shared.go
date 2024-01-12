package shared

import (
	"cointrading/app/domain/repositories"
	"cointrading/app/domain/valueobjects"
	"cointrading/app/infrastructure/factories/tradingfactory"
)

var (
	exchange = valueobjects.Bitflyer
	symbol   = valueobjects.BTCJPY
	duration = valueobjects.Day
)

func Exchange() *valueobjects.Exchange {
	return exchange
}

func SetExchange(e *valueobjects.Exchange) {
	exchange = e
}

func Symbol() *valueobjects.Symbol {
	return symbol
}

func SetSymbol(s *valueobjects.Symbol) {
	symbol = s
}

func Duration() *valueobjects.Duration {
  return duration
}

func SetDuration(d *valueobjects.Duration) {
  duration = d
}

func APIClient() repositories.TradingAPIClient {
	factory := tradingfactory.NewTradingFactory(exchange)

	return factory.NewTradingAPIClient()
}
