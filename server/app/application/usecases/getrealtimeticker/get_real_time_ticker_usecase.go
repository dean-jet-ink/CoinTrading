package getrealtimeticker

type GetRealTimeTickerUsecase interface {
	Handle(input *GetRealTimeTickerInput)
}
