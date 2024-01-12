package controllers

import (
	"cointrading/app/application/usecases/getdataframecandle"
	"cointrading/app/domain/myerror"
	"cointrading/app/presentation/router"
	"fmt"
)

type CandleController struct {
	getDataframeCandleUsecase getdataframecandle.GetDataframeCandleUsecase
}

func NewCandleController(getDataframeCandleUsecase getdataframecandle.GetDataframeCandleUsecase) *CandleController {
	return &CandleController{
		getDataframeCandleUsecase: getDataframeCandleUsecase,
	}
}

func (ca *CandleController) GetDataframeCandleStream(c router.Context) {
	upgrader := NewWebSocketUpgrader()

	conn, err := upgrader.Upgrade(c)
	if err != nil {
		handleError(err, c)
		return
	}
	defer conn.Close()

	for {
		input := &getdataframecandle.GetDataframeCandleInput{}

		if err := conn.ReadJSON(input); err != nil {
			err = fmt.Errorf("%w: Recieved bad request: %v", myerror.ErrBadRequest, err)
			handleError(err, c)
			return
		}

		output, err := ca.getDataframeCandleUsecase.Handle(input)
		if err != nil {
			handleError(err, c)
			return
		}

		conn.WriteJSON(output)
	}
}
