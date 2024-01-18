package getdataframecandle

type CandleParams struct {
	Message string `json:"message"`
}

type GetDataframeCandleInput struct {
	CandleParams CandleParams `json:"candle_params"`
}
