package controllers

import (
	"cointrading/app/presentation/router"
)

type TradingConfigController struct {
}

func NewTradingConfigController() *TradingConfigController {
	return &TradingConfigController{}
}

func (t *TradingConfigController) GetTradingConfig(c router.Context) {

}

func (t *TradingConfigController) GetDurations(c router.Context) {

}
