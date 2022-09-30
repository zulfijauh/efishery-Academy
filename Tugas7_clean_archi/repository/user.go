package repository

import (
	"clean-archi/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Store(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindByID(id int) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(id int) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) Store(user entity.User) (entity.User, error) {
	if err := r.db.Debug().Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r UserRepository) FindAll() ([]entity.User, error) {
	var user []entity.User
	if err := r.db.Debug().Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r UserRepository) FindByID(id int) (entity.User, error) {
	var user entity.User
	if err := r.db.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r UserRepository) Update(user entity.User) (entity.User, error) {
	if err := r.db.Debug().Save(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r UserRepository) Delete(id int) error {
	if err := r.db.Debug().Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
