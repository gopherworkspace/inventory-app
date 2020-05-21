package models

type Product struct {
	Id          int     `json:"id" bson:"_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	QTY_Stock   int     `json:"qty_stock"`
	Price       float64 `json:"price"`
	CategoryId  int     `json:"category_id"`
}

type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type Orders struct {
	OrderId       int    `json:"order_id"`
	ProductId     int    `json:"product_id"`
	CustomerTpeId int    `json:"customer_tpe_id"`
	PaymentId     int    `json:"payment_id"`
	OrderQuantity int    `json:"order_quantity"`
	OrderDate     string `json:"order_date"`
	OrderStatus   string `json:"order_status"`
}

type Payment struct {
	PaymentId          int    `json:"payment_id"`
	PaymentDescription string `json:"payment_description"`
}

type Cart struct {
	Name  string         `json:"name"`
	Email string         `json:"email"`
	Items []ProductItems `json:"items"`
}

type ProductItems struct {
	ProductId    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	Quantity     int     `json:"quantity"`
}
