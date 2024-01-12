package dto

type Candle struct {
	Time   string  `json:"time"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
}

type Symbol struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
