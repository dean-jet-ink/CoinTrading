package controllers

import (
	"cointrading/app/application/usecases/sendorder"
	"cointrading/app/presentation/router"
	"net/http"
)

type OrderController struct {
	sendOrderUsecase sendorder.SendOrderUsecase
}

func NewOrderController(sendOrderUsecase sendorder.SendOrderUsecase) *OrderController {
	return &OrderController{
		sendOrderUsecase: sendOrderUsecase,
	}
}

func (o *OrderController) SendOrder(c router.Context) {
	input := &sendorder.SendOrderInput{}

	if err := c.BindJSON(input); err != nil {
		handleError(err, c)
		return
	}

	output, err := o.sendOrderUsecase.Handle(input)
	if err != nil {
		handleError(err, c)
		return
	}

	c.JSON(http.StatusOK, output)
}

func (o *OrderController) GetOrders(c router.Context) {

}
