package getorders

type GetOrdersUsecase interface {
	Handle(input *GetOrdersInput) (*GetOrdersOutput, error)
}
