package entity

type Users struct {
	ID           int    `json:"id" gorm:"primary_key"`
	Nama         string `json:"nama" gorm:"not_null;default:null"`
	Gender       string `json:"gender" gorm:"type:char(30);not_null;default:null"`
	Email        string `json:"email" gorm:"typevarchar(60)not_null;default:null"`
	Phone        string `json:"phone" gorm:"typevarchar(12)not_null;default:null"`
	Transactions []Transactions
}

type CreateUserRequest struct {
	Nama         string `json:"nama" gorm:"not_null;default:null"`
	Gender       string `json:"gender" gorm:"type:char(30);not_null;default:null"`
	Email        string `json:"email" gorm:"typevarchar(60)not_null;default:null"`
	Phone        string `json:"phone" gorm:"typevarchar(12)not_null;default:null"`
	Transactions []Transactions
}

type UpdateUserRequest struct {
	Nama         string `json:"nama" gorm:"not_null;default:null"`
	Gender       string `json:"gender" gorm:"type:char(30);not_null;default:null"`
	Email        string `json:"email" gorm:"typevarchar(60)not_null;default:null"`
	Phone        string `json:"phone" gorm:"typevarchar(12)not_null;default:null"`
	Transactions []Transactions
}

type UserResponse struct {
	Nama         string `json:"nama" gorm:"not_null;default:null"`
	Gender       string `json:"gender" gorm:"type:char(30);not_null;default:null"`
	Email        string `json:"email" gorm:"typevarchar(60)not_null;default:null"`
	Phone        string `json:"phone" gorm:"typevarchar(12)not_null;default:null"`
	Transactions []Transactions
}

type DetailedUserResponse struct {
	Nama         string `json:"nama" gorm:"not_null;default:null"`
	Gender       string `json:"gender" gorm:"type:char(30);not_null;default:null"`
	Email        string `json:"email" gorm:"typevarchar(60)not_null;default:null"`
	Phone        string `json:"phone" gorm:"typevarchar(12)not_null;default:null"`
	Transactions []Transactions
}
