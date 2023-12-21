package backgrounds

import "cointrading/app/application/usecases/getrealtimeticker"

type GetRealTimeTickerBackground struct {
	getRealTimeTickerUsecase getrealtimeticker.GetRealTimeTickerUsecase
}

func NewGetRealTimeTickerBackground(getRealtimeTickerUsecase getrealtimeticker.GetRealTimeTickerUsecase) *GetRealTimeTickerBackground {
	return &GetRealTimeTickerBackground{
		getRealTimeTickerUsecase: getRealtimeTickerUsecase,
	}
}

func (g *GetRealTimeTickerBackground) Exec() {
	g.getRealTimeTickerUsecase.Handle()
}
