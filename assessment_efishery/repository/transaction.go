package repository

import (
	"assessment_efishery/entity"

	"gorm.io/gorm"
)

type ITransactionsRepository interface {
	Store(transaction entity.Transactions) (entity.Transactions, error)
	FindAll() ([]entity.Transactions, error)
	FindByID(id int) (entity.Transactions, error)
	Update(transaction entity.Transactions) (entity.Transactions, error)
	Delete(id int) error
	TransactionsOfUsers(user_id int) ([]entity.Transactions, error)
}

type TransactionsRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionsRepository {
	return &TransactionsRepository{db: db}
}

func (r TransactionsRepository) Store(transaction entity.Transactions) (entity.Transactions, error) {
	if err := r.db.Debug().Create(&transaction).Error; err != nil {
		return entity.Transactions{}, err
	}
	return transaction, nil
}

func (r TransactionsRepository) FindAll() ([]entity.Transactions, error) {
	var transaction []entity.Transactions
	if err := r.db.Debug().Find(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r TransactionsRepository) FindByID(id int) (entity.Transactions, error) {
	var transaction entity.Transactions
	if err := r.db.Debug().Where("id = ?", id).First(&transaction).Error; err != nil {
		return entity.Transactions{}, err
	}
	return transaction, nil
}

func (r TransactionsRepository) Update(transaction entity.Transactions) (entity.Transactions, error) {
	if err := r.db.Debug().Save(&transaction).Error; err != nil {
		return entity.Transactions{}, err
	}
	return transaction, nil
}

func (r TransactionsRepository) Delete(id int) error {
	if err := r.db.Debug().Delete(&entity.Transactions{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r TransactionsRepository) TransactionsOfUsers(users_id int) ([]entity.Transactions, error) {
	var transaction []entity.Transactions
	if err := r.db.Debug().Where("users_id = ?", users_id).Find(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}
