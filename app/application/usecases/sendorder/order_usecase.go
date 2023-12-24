package sendorder

type SendOrderUsecase interface {
	Handle(input *SendOrderInput) (*SendOrderOutput, error)
}
