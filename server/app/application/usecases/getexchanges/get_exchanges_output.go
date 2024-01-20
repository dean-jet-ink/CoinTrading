package getexchanges

import "cointrading/app/application/usecases/dto"

type GetExchangesOutput struct {
	Exchanges []*dto.Exchange `json:"exchanges"`
}
