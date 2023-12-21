package interactors

import (
	"cointrading/app/application/usecases/getrealtimeticker"
	"cointrading/app/domain/entities"
	"cointrading/app/domain/valueobjects"
	"cointrading/app/infrastructure/factories/tradingfactory"
	"log"
)

type getRealTimeTickerInteractor struct {
}

func NewGetRealTimeTickerUsecase() getrealtimeticker.GetRealTimeTickerUsecase {
	return &getRealTimeTickerInteractor{}
}

func (g *getRealTimeTickerInteractor) Handle() {
	g.getRealTimeTickerAllExchange()
}

func (g *getRealTimeTickerInteractor) getRealTimeTickerAllExchange() {
	factory := tradingfactory.NewTradingFactory()

	// exchanges := valueobjects.Exchanges()
	// symbols := valueobjects.Symbols()

	// for _, e := range exchanges {
	// 	client := factory.CreateTradingAPIClient(e)

	// 	for _, s := range symbols {
	// 		tickerChan := make(chan *entities.Ticker)
	// 		client.GetRealTimeTicker(s, tickerChan)
	// 	}
	// }

	bitflyer, _ := valueobjects.NewExchange(1)

	client := factory.CreateTradingAPIClient(bitflyer)

	symbol, _ := valueobjects.NewSymbol(1)

	tickerChan := make(chan *entities.Ticker)

	go client.GetRealTimeTicker(symbol, tickerChan)

	for {
		select {
		case ticker := <-tickerChan:
			log.Printf("%+v", ticker)
		}
	}
}
