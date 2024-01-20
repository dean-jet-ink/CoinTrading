package gettradingconfig

import "cointrading/app/application/usecases/dto"

type GetTradingConfigOutput struct {
	Exchange *dto.Exchange `json:"exchange"`
	Symbol   *dto.Symbol   `json:"symbol"`
	Duration *dto.Duration `json:"duration"`
}
