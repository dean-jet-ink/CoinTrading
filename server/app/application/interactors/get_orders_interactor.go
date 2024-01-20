package interactors

import (
	"cointrading/app/application/usecases/getorders"
	"cointrading/app/domain/repositories"
	"cointrading/app/domain/valueobjects"
)

type getOrdersInteractor struct {
	tradingClient repositories.TradingAPIClient
}

func NewGetOrdersInteractor(tradingClient repositories.TradingAPIClient) getorders.GetOrdersUsecase {
	return &getOrdersInteractor{
		tradingClient: tradingClient,
	}
}

func (g *getOrdersInteractor) Handle(input *getorders.GetOrdersInput) (*getorders.GetOrdersOutput, error) {
	symbol, err := valueobjects.NewSymbol(input.Symbol)
	if err != nil {
		return nil, err
	}

	orders, err := g.tradingClient.GetOrders(symbol)
	if err != nil {
		return nil, err
	}

	output := &getorders.GetOrdersOutput{}

	for _, o := range orders {
		order := &getorders.Order{
			ID:          o.ID(),
			Symbol:      o.Symbol().DisplayValue(),
			Side:        o.Side().DisplayValue(),
			OrderType:   o.OrderType().DisplayValue(),
			Price:       o.Price(),
			Size:        o.Size(),
			TimeInForce: o.TimeInForce().DisplayValue(),
			Status:      o.Status().DisplayValue(),
		}

		output.Orders = append(output.Orders, order)
	}

	return output, nil
}
