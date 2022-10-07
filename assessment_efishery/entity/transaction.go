package entity

import "time"

type Transactions struct {
	ID              int `json:"id" gorm:"primary_key"`
	UsersID         int `json:"user_id" gorm:"not_null;default:null"`
	Cart            []Cart
	TotalProduct    int       `json:"total_product"`
	Subtotal        int       `json:"subtotal"`
	ShippingAddress string    `json:"address"`
	Created         time.Time `json:"created_at"`
	Limit           time.Time `json:"time_limit"`
	Proof           string    `json:"proof"`
	Status          string    `json:"status"`
}

type CreateTransactionsRequest struct {
	UsersID         int `json:"user_id"`
	Cart            []Cart
	TotalProduct    int       `json:"total_product"`
	Subtotal        int       `json:"subtotal"`
	ShippingAddress string    `json:"address"`
	Created         time.Time `json:"created_at"`
	Limit           time.Time `json:"time_limit"`
	Proof           string    `json:"proof"`
	Status          string    `json:"status"`
}

type UpdateTransactionsRequest struct {
	ID              int `json:"id"`
	UsersID         int `json:"user_id"`
	Cart            []Cart
	TotalProduct    int       `json:"total_product"`
	Subtotal        int       `json:"subtotal"`
	ShippingAddress string    `json:"address"`
	Created         time.Time `json:"created_at"`
	Limit           time.Time `json:"time_limit"`
	Proof           string    `json:"proof"`
	Status          string    `json:"status"`
}

type TransactionsResponse struct {
	ID              int `json:"id"`
	UsersID         int `json:"user_id"`
	Cart            []Cart
	TotalProduct    int       `json:"total_product"`
	Subtotal        int       `json:"subtotal"`
	ShippingAddress string    `json:"address"`
	Created         time.Time `json:"created_at"`
	Limit           time.Time `json:"time_limit"`
	Proof           string    `json:"proof"`
	Status          string    `json:"status"`
}
