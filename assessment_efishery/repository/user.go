package repository

import (
	"assessment_efishery/entity"

	"gorm.io/gorm"
)

type IUsersRepository interface {
	Store(user entity.Users) (entity.Users, error)
	FindAll() ([]entity.Users, error)
	FindByID(id int) (entity.Users, error)
	Update(user entity.Users) (entity.Users, error)
	Delete(id int) error
}

type UsersRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r UsersRepository) Store(user entity.Users) (entity.Users, error) {
	if err := r.db.Debug().Create(&user).Error; err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

func (r UsersRepository) FindAll() ([]entity.Users, error) {
	var user []entity.Users
	if err := r.db.Debug().Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r UsersRepository) FindByID(id int) (entity.Users, error) {
	var user entity.Users
	if err := r.db.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

func (r UsersRepository) Update(user entity.Users) (entity.Users, error) {
	if err := r.db.Debug().Save(&user).Error; err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

func (r UsersRepository) Delete(id int) error {
	if err := r.db.Debug().Delete(&entity.Users{}, id).Error; err != nil {
		return err
	}
	return nil
}
