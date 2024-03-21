package retail

type Products struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Shelves struct {
	Id       int    `json:"id" db:"id"`
	Location string `json:"location" db:"location"`
}

type Orders struct {
	Id        int    `json:"id" db:"id"`
	OrderDate string `json:"order_date" db:"order_date"`
}

type OrderDetails struct {
	OrderId           int      `json:"order_id"`
	ProductId         int      `json:"product_id"`
	ProductName       string   `json:"product_name"`
	ShelfId           int      `json:"shelf_id"`
	OrderCount        int      `json:"order_count"`
	AdditionalShelves []string `json:"additional_shelves"`
	ShelfLocation     string   `json:"shelf_location"`
}
