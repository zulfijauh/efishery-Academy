package repository

import (
	"assessment_efishery/entity"

	"gorm.io/gorm"
)

type ICartRepository interface {
	Store(cart entity.Cart) (entity.Cart, error)
	FindAll() ([]entity.Cart, error)
	FindByID(id int) (entity.Cart, error)
	Update(cart entity.Cart) (entity.Cart, error)
	Delete(id int) error
	AllCartID(id int) ([]entity.Cart, error)
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r CartRepository) Store(cart entity.Cart) (entity.Cart, error) {
	if err := r.db.Debug().Create(&cart).Error; err != nil {
		return entity.Cart{}, err
	}
	return cart, nil
}

func (r CartRepository) FindAll() ([]entity.Cart, error) {
	var cart []entity.Cart
	if err := r.db.Debug().Find(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (r CartRepository) FindByID(id int) (entity.Cart, error) {
	var cart entity.Cart
	if err := r.db.Debug().Where("id = ?", id).First(&cart).Error; err != nil {
		return entity.Cart{}, err
	}
	return cart, nil
}

func (r CartRepository) Update(cart entity.Cart) (entity.Cart, error) {
	if err := r.db.Debug().Save(&cart).Error; err != nil {
		return entity.Cart{}, err
	}
	return cart, nil
}

func (r CartRepository) Delete(id int) error {
	if err := r.db.Debug().Delete(&entity.Cart{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r CartRepository) AllCartID(id int) ([]entity.Cart, error) {
	var cart []entity.Cart
	if err := r.db.Debug().Find(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
