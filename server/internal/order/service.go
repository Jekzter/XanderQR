package order

type Service interface {
	GetAll() []Order
	Create(o Order) Order
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAll() []Order {
	return s.repo.FindAll()
}

func (s *service) Create(o Order) Order {
	var total float64

	for _, item := range o.Items {
		total += item.Price * float64(item.Qty)
	}

	o.Total = total
	o.Status = "pending"
	return s.repo.Create(o)
}
