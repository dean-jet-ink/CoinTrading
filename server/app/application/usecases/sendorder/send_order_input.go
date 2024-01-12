package sendorder

type SendOrderInput struct {
	Symbol      int     `json:"symbol" validate:"required"`
	Side        string  `json:"side" validate:"required"`
	OrderType   int     `json:"order_type" validate:"required"`
	Price       int     `json:"price,omitempty"`
	Size        float64 `json:"size" validate:"required"`
	TimeInForce int     `json:"time_in_force" validate:"required"`
}
