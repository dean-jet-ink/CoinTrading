package getdurations

import "cointrading/app/application/usecases/dto"

type GetDurationsOutput struct {
	Durations []*dto.Duration `json:"durations"`
}
