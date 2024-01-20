package updatetradingconfig

type UpdateTradingConfigInput struct {
	Exchange int    `json:"exchange"`
	Symbol   int    `json:"symbol"`
	Duration string `json:"duration"`
}
