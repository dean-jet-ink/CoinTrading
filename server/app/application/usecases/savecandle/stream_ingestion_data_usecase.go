package savecandle

type SaveCandleUsecase interface {
	Handle(input *SaveCandleInput) (*SaveCandleOutput, error)
}
