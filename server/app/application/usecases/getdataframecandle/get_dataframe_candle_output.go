package getdataframecandle

import "cointrading/app/application/usecases/dto"

type GetDataframeCandleOutput struct {
	Candles []*dto.Candle `json:"candles"`
}