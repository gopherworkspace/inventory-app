package models

type Invoice struct {
	UserName      string  `json:"user_name"`
	Email         string  `json:"email"`
	Item          []Items `json:"items"`
	TotalPrice    float64 `json:"total_price"`
	DeliveredDate string  `json:"delivered_date"`
}

type Items struct {
	ItemName  string
	ItemPrice float64
	Quntity   int
}
