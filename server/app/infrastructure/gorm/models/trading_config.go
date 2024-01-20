package models

type TradingConfig struct {
	Exchange int    `json:"exchange" gorm:"type:int(11);not null"`
	Symbol   int    `json:"symbol" gorm:"type:int(11);not null"`
	Duration string `json:"duration" gorm:"type:varchar(255);not null"`
}
