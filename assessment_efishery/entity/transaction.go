package entity

import (
	"database/sql"
	"time"
)

type Transactions struct {
	ID           int `json:"id" gorm:"primaryKey"`
	UsersID      int `json:"users_id" gorm:"not null;default:null"`
	Cart         []Cart
	TotalProduct int            `json:"total_product"`
	Subtotal     int            `json:"subtotal"`
	Address      string         `json:"address" gorm:"default:null"`
	Created      sql.NullTime   `gorm:"default:null"`
	Limit        sql.NullTime   `gorm:"default:null"`
	Proof        sql.NullString `gorm:"default:null"`
	Status       string         `json:"status"`
}

type CreateTransactionsRequest struct {
	UsersID      int `json:"users_id"`
	Cart         []Cart
	TotalProduct int            `json:"total_product"`
	Subtotal     int            `json:"subtotal"`
	Address      string         `json:"address"`
	Created      time.Time      `json:"created_at"`
	Limit        time.Time      `json:"time_limit"`
	Proof        sql.NullString `gorm:"default:null"`
	Status       string         `json:"status"`
}

type UpdateTransactionsRequest struct {
	ID           int `json:"id"`
	UsersID      int `json:"users_id"`
	Cart         []Cart
	TotalProduct int            `json:"total_product"`
	Subtotal     int            `json:"subtotal"`
	Address      string         `json:"address" gorm:"not_null;default:null"`
	Created      sql.NullTime   `gorm:"default:null"`
	Limit        sql.NullTime   `gorm:"default:null"`
	Proof        sql.NullString `gorm:"default:null"`
	Status       string         `json:"status"`
}

type TransactionsResponse struct {
	ID           int `json:"id"`
	UsersID      int `json:"users_id"`
	Cart         []Cart
	TotalProduct int            `json:"total_product"`
	Subtotal     int            `json:"subtotal"`
	Address      string         `json:"address"`
	Created      sql.NullTime   `gorm:"default:null"`
	Limit        sql.NullTime   `gorm:"default:null"`
	Proof        sql.NullString `gorm:"default:null"`
	Status       string         `json:"status"`
}
