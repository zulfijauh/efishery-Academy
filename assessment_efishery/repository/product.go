package repository

import (
	"assessment_efishery/entity"

	"gorm.io/gorm"
)

type IProductsRepository interface {
	Store(product entity.Products) (entity.Products, error)
	FindAll() ([]entity.Products, error)
	FindByID(id int) (entity.Products, error)
	ToCart(id int) (entity.Products, error)
	FindByCategory(category string) ([]entity.Products, error)
	FindByPrice(priceMin int, priceMax int) ([]entity.Products, error)
	Update(product entity.Products) (entity.Products, error)
	Delete(id int) error
}

type ProductsRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{db: db}
}

func (r ProductsRepository) Store(product entity.Products) (entity.Products, error) {
	if err := r.db.Debug().Create(&product).Error; err != nil {
		return entity.Products{}, err
	}
	return product, nil
}

func (r ProductsRepository) FindAll() ([]entity.Products, error) {
	var product []entity.Products
	if err := r.db.Debug().Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r ProductsRepository) FindByPrice(priceMin int, priceMax int) ([]entity.Products, error) {
	var product []entity.Products
	if err := r.db.Debug().Where("harga >= ? AND harga <= ? ", priceMin, priceMax).Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r ProductsRepository) FindByID(id int) (entity.Products, error) {
	var product entity.Products
	if err := r.db.Debug().Where("id = ?", id).First(&product).Error; err != nil {
		return entity.Products{}, err
	}
	return product, nil
}

func (r ProductsRepository) ToCart(id int) (entity.Products, error) {
	var product entity.Products
	if err := r.db.Debug().Select("nama").Where("id = ?", id).Find(&product).Error; err != nil {
		return entity.Products{}, err
	}
	return product, nil
}

func (r ProductsRepository) FindByCategory(category string) ([]entity.Products, error) {
	var product []entity.Products
	if err := r.db.Debug().Where("kategori = ?", category).Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r ProductsRepository) Update(product entity.Products) (entity.Products, error) {
	if err := r.db.Debug().Save(&product).Error; err != nil {
		return entity.Products{}, err
	}
	return product, nil
}

func (r ProductsRepository) Delete(id int) error {
	if err := r.db.Debug().Delete(&entity.Products{}, id).Error; err != nil {
		return err
	}
	return nil
}
