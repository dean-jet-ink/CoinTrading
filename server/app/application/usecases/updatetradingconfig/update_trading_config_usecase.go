package updatetradingconfig

type UpdateTradingConfigUsecase interface {
	Handle(input *UpdateTradingConfigInput) error
}
