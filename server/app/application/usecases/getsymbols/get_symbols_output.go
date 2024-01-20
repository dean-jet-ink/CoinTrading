package getsymbols

import "cointrading/app/application/usecases/dto"

type GetSymbolsOutput struct {
	Symbols []*dto.Symbol `json:"symbols"`
}
