package models

import "time"

type Candle struct {
	Time   time.Time `json:"time" gorm:"primaryKey"`
	Open   float64   `json:"open"`
	Close  float64   `json:"close"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Volume float64   `json:"volume"`
}
