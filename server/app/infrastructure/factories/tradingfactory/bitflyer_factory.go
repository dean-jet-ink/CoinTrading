package tradingfactory

import (
	"cointrading/app/config"
	"cointrading/app/domain/repositories"
	"cointrading/app/infrastructure/apiclients/bitflyer"
)

type bitflyerFactory struct {
}

func (b *bitflyerFactory) NewTradingAPIClient() repositories.TradingAPIClient {
	return bitflyer.NewBitflyerClient(config.BitflyerApiKey(), config.BitflyerApiSecret())
}
