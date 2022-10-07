package usecase

import (
	"assessment_efishery/entity"
	"assessment_efishery/repository"

	"github.com/jinzhu/copier"
)

// USERS
type IUserUsecase interface {
	CreateUser(user entity.Users) (entity.Users, error)
	GetAllUser() ([]entity.Users, error)
	GetUserById(id int) (entity.Users, error)
	UpdateUser(UserRequest entity.UpdateUserRequest) (entity.UserResponse, error)
	DeleteUser(id int) error
}

type UserUsecase struct {
	UserRepository repository.IUsersRepository
}

func NewUserUsecase(userRepository repository.IUsersRepository) *UserUsecase {
	return &UserUsecase{userRepository}
}

func (useCase UserUsecase) CreateUser(user entity.Users) (entity.Users, error) {
	u := entity.Users{
		Nama:         user.Nama,
		Gender:       user.Gender,
		Email:        user.Email,
		Phone:        user.Phone,
		Transactions: []entity.Transactions{},
	}
	users, err := useCase.UserRepository.Store(u)
	if err != nil {
		return entity.Users{}, err
	}
	userRes := entity.Users{
		Nama:   users.Nama,
		Gender: users.Gender,
		Email:  users.Email,
		Phone:  users.Phone,
	}
	return userRes, nil
}

func (useCase UserUsecase) GetAllUser() ([]entity.Users, error) {
	users, err := useCase.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	userRes := []entity.Users{}
	for _, user := range users {
		appendUser := entity.Users{
			ID:     user.ID,
			Nama:   user.Nama,
			Gender: user.Gender,
			Email:  user.Email,
			Phone:  user.Phone,
		}
		userRes = append(userRes, appendUser)
	}
	return userRes, nil
}

func (useCase UserUsecase) GetUserById(id int) (entity.Users, error) {
	users, err := useCase.UserRepository.FindByID(id)
	if err != nil {
		return entity.Users{}, err
	}
	userRes := entity.Users{
		ID:     users.ID,
		Nama:   users.Nama,
		Gender: users.Gender,
		Email:  users.Email,
		Phone:  users.Phone,
	}
	return userRes, nil
}

func (useCase UserUsecase) UpdateUser(userRequest entity.UpdateUserRequest, id int) (entity.UserResponse, error) {
	users, err := useCase.UserRepository.FindByID(id)
	if err != nil {
		return entity.UserResponse{}, err
	}
	users.Nama = userRequest.Nama
	users.Gender = userRequest.Gender
	users.Email = userRequest.Email
	users.Phone = userRequest.Phone

	copier.CopyWithOption(&users, &userRequest, copier.Option{IgnoreEmpty: true})
	users, _ = useCase.UserRepository.Update(users)
	userRes := entity.UserResponse{}
	copier.Copy(&userRes, &users)
	return userRes, nil
}

func (useCase UserUsecase) DeleteUser(id int) error {
	_, err := useCase.UserRepository.FindByID(id)
	if err != nil {
		return err
	}
	err = useCase.UserRepository.Delete(id)
	return err
}
