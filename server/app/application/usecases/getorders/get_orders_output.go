package getorders

type GetOrdersOutput struct {
	Orders []*Order `json:"orders"`
}

type Order struct {
	ID          string  `json:"id"`
	Symbol      string  `json:"symbol"`
	Side        string  `json:"side"`
	OrderType   string  `json:"order_type"`
	Price       int     `json:"price"`
	Size        float64 `json:"size"`
	TimeInForce string  `json:"time_in_force"`
	Status      string  `json:"status"`
}
