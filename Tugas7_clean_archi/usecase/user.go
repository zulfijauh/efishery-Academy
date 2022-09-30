package usecase

import (
	"clean-archi/entity"
	"clean-archi/repository"
)

type IUserUsecase interface {
	CreateUser(user entity.UserRequest) (entity.User, error)
	GetAllUser() ([]entity.User, error)
	GetUserById(id int) (entity.User, error)
	UpdateUser(UserRequest entity.UserUpdateRequest) (entity.User, error)
	DeleteUser(id int) error
}

type UserUsecase struct {
	UserRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepository}
}

func (useCase UserUsecase) CreateUser(user entity.UserRequest) (entity.UserResponse, error) {
	u := entity.User{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	users, err := useCase.UserRepository.Store(u)

	if err != nil {
		return entity.UserResponse{}, err
	}

	userRes := entity.UserResponse{
		Username: users.Username,
		Email:    users.Email,
		Phone:    users.Phone,
	}

	return userRes, nil
}

func (useCase UserUsecase) GetAllUser() ([]entity.UserResponse, error) {
	users, err := useCase.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	userRes := []entity.UserResponse{}
	for _, user := range users {
		appendUser := entity.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		}
		userRes = append(userRes, appendUser)
	}
	return userRes, nil
}

func (useCase UserUsecase) GetUserById(id int) (entity.UserResponse, error) {
	users, err := useCase.UserRepository.FindByID(id)
	if err != nil {
		return entity.UserResponse{}, err
	}
	userRes := entity.UserResponse{
		ID:       users.ID,
		Username: users.Username,
		Email:    users.Email,
		Phone:    users.Phone,
	}
	return userRes, nil
}

func (useCase UserUsecase) UpdateUser(UserRequest entity.UserUpdateRequest, id int) (entity.UserResponse, error) {
	users, err := useCase.UserRepository.FindByID(id)
	if err != nil {
		return entity.UserResponse{}, err
	}
	updatedUser := entity.User{
		ID:       users.ID,
		Username: UserRequest.Username,
		Email:    UserRequest.Email,
		Phone:    UserRequest.Phone,
	}
	users, err = useCase.UserRepository.Update(updatedUser)
	if err != nil {
		return entity.UserResponse{}, err
	}
	userRes := entity.UserResponse{
		Username: users.Username,
		Email:    users.Email,
		Phone:    users.Phone,
	}
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
