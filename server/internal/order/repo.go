package order

type Repository interface {
	FindAll() []Order
	Create(o Order) Order
}

type repository struct {
	orders []Order
}

func NewRepository() Repository {
	return &repository{orders: []Order{}}
}

func (r *repository) FindAll() []Order {
	return r.orders
}

func (r *repository) Create(o Order) Order {
	r.orders = append(r.orders, o)
	return o
}
