package usecase

import (
	"assessment_efishery/config"
	"assessment_efishery/entity"
	"assessment_efishery/repository"

	"github.com/jinzhu/copier"
)

// USERS
type IUserUsecase interface {
	CreateUser(user entity.CreateUserRequest) (entity.Users, error)
	GetAllUser() ([]entity.UserResponse, error)
	GetUserById(id int) (entity.DetailedUserResponse, error)
	UpdateUser(UserRequest entity.UpdateUserRequest) (entity.UserResponse, error)
	DeleteUser(id int) error
}

type UserUsecase struct {
	UserRepository repository.IUsersRepository
}

func NewUserUsecase(userRepository repository.IUsersRepository) *UserUsecase {
	return &UserUsecase{userRepository}
}

func (useCase UserUsecase) CreateUser(user entity.CreateUserRequest) (entity.UserResponse, error) {
	u := entity.Users{
		Nama:         user.Nama,
		Gender:       user.Gender,
		Email:        user.Email,
		Phone:        user.Phone,
		Transactions: []entity.Transactions{},
	}
	users, err := useCase.UserRepository.Store(u)
	if err != nil {
		return entity.UserResponse{}, err
	}
	userRes := entity.UserResponse{
		ID:     users.ID,
		Nama:   users.Nama,
		Gender: users.Gender,
		Email:  users.Email,
		Phone:  users.Phone,
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

func (useCase UserUsecase) GetUserById(id int) (entity.DetailedUserResponse, error) {
	users, err := useCase.UserRepository.FindByID(id)
	if err != nil {
		return entity.DetailedUserResponse{}, err
	}
	transactionRepository := repository.NewTransactionRepository(config.DB)
	transactionUsecase := NewTransactionUsecase(transactionRepository)
	transaction, _ := transactionUsecase.TransactiontoUser(users.ID)
	userRes := entity.DetailedUserResponse{
		Nama:         users.Nama,
		Gender:       users.Gender,
		Email:        users.Email,
		Phone:        users.Phone,
		Transactions: transaction,
	}
	return userRes, nil
}

func (useCase UserUsecase) UpdateUser(userRequest entity.UpdateUserRequest, id int) (entity.DetailedUserResponse, error) {
	users, err := useCase.UserRepository.FindByID(id)
	if err != nil {
		return entity.DetailedUserResponse{}, err
	}
	// Add transaction(s) related to customers
	transactionRepository := repository.NewTransactionRepository(config.DB)
	transactionUsecase := NewTransactionUsecase(transactionRepository)
	transaction, _ := transactionUsecase.TransactiontoUser(users.ID)

	users.Nama = userRequest.Nama
	users.Gender = userRequest.Gender
	users.Email = userRequest.Email
	users.Phone = userRequest.Phone
	users.Transactions = transaction

	copier.CopyWithOption(&users, &userRequest, copier.Option{IgnoreEmpty: true})
	users, _ = useCase.UserRepository.Update(users)
	userRes := entity.DetailedUserResponse{}
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
