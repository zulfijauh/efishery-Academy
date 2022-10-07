package entity

type Cart struct {
	ID             int    `json:"id" gorm:"primary_key"`
	TransactionsID int    `json:"transactions_id" gorm:"not_null;default:null"`
	ProductsID     int    `json:"products_id" gorm:"not_null;default:null"`
	Products       string `json:"nama_products"`
	Harga          int    `json:"harga_products"`
	Quantity       int    `json:"quantity" gorm:"not_null;default:null"`
	Total          int    `json:"total"`
}

type CreateCartRequest struct {
	TransactionsID int    `json:"transactions_id" gorm:"primary_key"`
	ProductsID     int    `json:"products_id" `
	Products       string `json:"nama_products"`
	Harga          int    `json:"harga_products"`
	Quantity       int    `json:"quantity"`
	Total          int    `json:"total"`
}

type UpdateCartRequest struct {
	TransactionsID int    `json:"transactions_id"`
	ProductsID     int    `json:"products_id"`
	Products       string `json:"nama_products"`
	Harga          int    `json:"harga_products"`
	Quantity       int    `json:"quantity"`
	Total          int    `json:"total"`
}

type CartResponse struct {
	ID             int    `json:"id"`
	TransactionsID int    `json:"transactions_id"`
	ProductsID     int    `json:"products_id"`
	Products       string `json:"nama_products"`
	Harga          int    `json:"harga_products"`
	Quantity       int    `json:"quantity"`
	Total          int    `json:"total"`
}
