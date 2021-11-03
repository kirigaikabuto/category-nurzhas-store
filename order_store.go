package category_nurzhas_store

type OrderStore interface {
	CreateOrder(order *Order) (*Order, error)
	ListOrder() ([]Order, error)
}
