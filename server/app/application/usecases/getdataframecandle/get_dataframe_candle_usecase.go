package getdataframecandle

type GetDataframeCandleUsecase interface {
	Handle(input *GetDataframeCandleInput) (*GetDataframeCandleOutput, error)
}
