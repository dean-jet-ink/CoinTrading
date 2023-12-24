package sendorder

type SendOrderOutput struct {
	Id          string  `json:"id"`
	Symbol      string  `json:"symbol"`
	Side        string  `json:"side"`
	OrderType   string  `json:"order_type"`
	Price       int     `json:"price"`
	Size        float64 `json:"size"`
	TimeInForce string  `json:"time_in_force"`
	Status      string  `json:"status"`
}
