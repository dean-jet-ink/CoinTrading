package tradingfactory

import (
	"cointrading/app/config"
	"cointrading/app/domain/factories"
	"cointrading/app/domain/repositories"
	"cointrading/app/domain/valueobjects"
	"cointrading/app/infrastructure/apiclients"
	"log"
)

type TradingFactory struct {
}

func NewTradingFactory() factories.ITradingFactory {
	return &TradingFactory{}
}

func (t *TradingFactory) CreateTradingAPIClient(exchange *valueobjects.Exchange) repositories.TradingAPIClient {
	switch {
	case exchange.IsBitflyer():
		return apiclients.NewBitflyerClient(config.BitflyerApiKey(), config.BitflyerApiSecret())
	default:
		log.Printf("Unexpected exchange code %v. Defaulting to bitflyer", exchange.Value())
		return apiclients.NewBitflyerClient(config.BitflyerApiKey(), config.BitflyerApiSecret())
	}
}
