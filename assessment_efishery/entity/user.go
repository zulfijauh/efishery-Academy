package entity

type Users struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Nama         string `json:"nama" gorm:"not null;default:null"`
	Gender       string `json:"gender" gorm:"type:char(1);not null;default:null"`
	Email        string `json:"email" gorm:"typevarchar(60)not null;default:null"`
	Phone        string `json:"phone" gorm:"typevarchar(12)not null;default:null"`
	Transactions []Transactions
}

type CreateUserRequest struct {
	Nama   string `json:"nama"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}

type UpdateUserRequest struct {
	Nama         string `json:"nama"`
	Gender       string `json:"gender"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Transactions []Transactions
}

type UserResponse struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}

type DetailedUserResponse struct {
	Nama         string `json:"nama"`
	Gender       string `json:"gender"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Transactions []Transactions
}
