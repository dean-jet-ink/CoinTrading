package interactors

import (
	"cointrading/app/application/usecases/sendorder"
	"cointrading/app/domain/entities"
	"cointrading/app/domain/repositories"
)

type orderInteractor struct {
	tradingClient repositories.TradingAPIClient
}

func NewOrderUsecase(tradingClient repositories.TradingAPIClient) sendorder.SendOrderUsecase {
	return &orderInteractor{
		tradingClient: tradingClient,
	}
}

func (o *orderInteractor) Handle(input *sendorder.SendOrderInput) (*sendorder.SendOrderOutput, error) {
	order, err := entities.NewOrder("", input.Symbol, input.Side, input.OrderType, input.Price, input.Size, input.TimeInForce, 0)
	if err != nil {
		return nil, err
	}

	orderId, err := o.tradingClient.SendOrder(order)
	if err != nil {
		return nil, err
	}

	order, err = o.tradingClient.GetOrder(orderId)
	if err != nil {
		return nil, err
	}

	output := &sendorder.SendOrderOutput{
		Id:          order.Id(),
		Symbol:      order.Symbol().DisplayValue(),
		Side:        order.Side().DisplayValue(),
		OrderType:   order.OrderType().DisplayValue(),
		Price:       order.Price(),
		Size:        order.Size(),
		TimeInForce: order.TimeInForce().DisplayValue(),
		Status:      order.Status().DisplayValue(),
	}

	return output, nil
}
