package order

type Item struct {
	Name  string  `json:"name"`
	Qty   int     `json:"qty"`
	Price float64 `json:"price"`
}

type Order struct {
	ID      string  `json:"id"`
	TableNo string  `json:"table_no"`
	Items   []Item  `json:"items"`
	Total   float64 `json:"total"`
	Status  string  `json:"status"`
}
