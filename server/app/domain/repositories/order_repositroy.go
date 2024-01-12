package repositories

type OrderRepository interface {
	Create()
	Update()
	FindByID()
	FindOrdersBySymbol()
}
