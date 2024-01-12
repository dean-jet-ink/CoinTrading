package backgrounds

import (
	"cointrading/app/application/usecases/getrealtimeticker"
	"cointrading/app/application/usecases/savecandle"
	"cointrading/app/domain/entities"
	"cointrading/app/domain/valueobjects"
	"context"
	"log"
)

type StreamIngestionDataBackground struct {
	getRealTimeTickerUse getrealtimeticker.GetRealTimeTickerUsecase
	saveCandleUse        savecandle.SaveCandleUsecase
}

func NewStreamIngestionDataBackground(
	getRealtimeTickerUse getrealtimeticker.GetRealTimeTickerUsecase,
	saveCandleUse savecandle.SaveCandleUsecase,
) *StreamIngestionDataBackground {
	return &StreamIngestionDataBackground{
		getRealTimeTickerUse: getRealtimeTickerUse,
		saveCandleUse:        saveCandleUse,
	}
}

func (s *StreamIngestionDataBackground) Exec() {
	exchanges := valueobjects.Exchanges()
	symbols := valueobjects.Symbols()
	durations := valueobjects.Durations()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errChan := make(chan error)

	for _, e := range exchanges {
		for _, sy := range symbols {
			for _, d := range durations {
				go s.streamIngestionData(ctx, cancel, errChan, e, sy, d)
			}
		}
	}

	err := <-errChan

	log.Println(err)
}

func (s *StreamIngestionDataBackground) streamIngestionData(
	ctx context.Context,
	cancel context.CancelFunc,
	errChan chan error,
	exchange *valueobjects.Exchange,
	symbol *valueobjects.Symbol,
	duration *valueobjects.Duration,
) {
	tickerChan := make(chan *entities.Ticker)

	getRealTimeTickerInput := &getrealtimeticker.GetRealTimeTickerInput{
		CTX:        ctx,
		TickerChan: tickerChan,
		ErrChan:    errChan,
		Exchange:   exchange,
		Symbol:     symbol,
	}

	go s.getRealTimeTickerUse.Handle(getRealTimeTickerInput)

	for {
		select {
		case t := <-tickerChan:
			saveCandleInput := &savecandle.SaveCandleInput{
				Exchange: exchange,
				Duration: duration,
				Ticker:   t,
			}

			// outputを後でハンドリング
			_, err := s.saveCandleUse.Handle(saveCandleInput)
			if err != nil {
				errChan <- err
				cancel()
				return
			}

		case <-ctx.Done():
			log.Printf("Canceled stream: %s/%s/%s", exchange.DisplayValue(), symbol.DisplayValue(), duration.DisplayValue())
			return
		}
	}
}
