package controllers

import (
	"cointrading/app/application/usecases/getdurations"
	"cointrading/app/application/usecases/getexchanges"
	"cointrading/app/application/usecases/getsymbols"
	"cointrading/app/application/usecases/gettradingconfig"
	"cointrading/app/application/usecases/updatetradingconfig"
	"cointrading/app/domain/myerror"
	"cointrading/app/presentation/router"
	"fmt"
	"net/http"
)

type TradingConfigController struct {
	getTradingConfigUsecase    gettradingconfig.GetTradingConfigUsecase
	getExchangesUsecase        getexchanges.GetExchangesUsecase
	getSymbolsUsecase          getsymbols.GetSymbolsUsecase
	getDurationsUsecase        getdurations.GetDurationsUsecase
	updateTradingConfigUsecase updatetradingconfig.UpdateTradingConfigUsecase
}

func NewTradingConfigController(
	getTradingConfigUsecase gettradingconfig.GetTradingConfigUsecase,
	getExchangesUsecase getexchanges.GetExchangesUsecase,
	getSymbolsUsecase getsymbols.GetSymbolsUsecase,
	getDurationsUsecase getdurations.GetDurationsUsecase,
	updateTradingConfigUsecase updatetradingconfig.UpdateTradingConfigUsecase,
) *TradingConfigController {
	return &TradingConfigController{
		getTradingConfigUsecase:    getTradingConfigUsecase,
		getExchangesUsecase:        getExchangesUsecase,
		getSymbolsUsecase:          getSymbolsUsecase,
		getDurationsUsecase:        getDurationsUsecase,
		updateTradingConfigUsecase: updateTradingConfigUsecase,
	}
}

func (t *TradingConfigController) GetTradingConfig(c router.Context) {
	output := t.getTradingConfigUsecase.Handle()

	c.JSON(http.StatusOK, output)
}

func (t *TradingConfigController) GetExchanges(c router.Context) {
	output := t.getExchangesUsecase.Handle()

	c.JSON(http.StatusOK, output)
}

func (t *TradingConfigController) GetSymbols(c router.Context) {
	output := t.getSymbolsUsecase.Handle()

	c.JSON(http.StatusOK, output)
}

func (t *TradingConfigController) GetDurations(c router.Context) {
	output := t.getDurationsUsecase.Handle()

	c.JSON(http.StatusOK, output)
}

func (t *TradingConfigController) UpdateTrdingConfig(c router.Context) {
	input := &updatetradingconfig.UpdateTradingConfigInput{}

	if err := c.BindJSON(input); err != nil {
		err = fmt.Errorf("%w: %v", myerror.ErrBadRequest, err)
		handleError(err, c)
		return
	}

	if err := t.updateTradingConfigUsecase.Handle(input); err != nil {
		handleError(err, c)
		return
	}

	c.NoContent(http.StatusOK)
}
