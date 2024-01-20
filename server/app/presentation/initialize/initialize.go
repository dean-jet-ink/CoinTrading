package initialize

import "cointrading/app/application/usecases/initializetradingconfig"

type InitExecutor struct {
	initializeTradingConfigUsecase initializetradingconfig.InitializeTradingConfigUsecase
}

func NewInitExecutor(initializeTradingConfigUsecase initializetradingconfig.InitializeTradingConfigUsecase) *InitExecutor {
	return &InitExecutor{
		initializeTradingConfigUsecase: initializeTradingConfigUsecase,
	}
}

func (i *InitExecutor) InitializeTradingConfig() {
	i.initializeTradingConfigUsecase.Handle()
}
