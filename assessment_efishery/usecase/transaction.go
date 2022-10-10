package usecase

import (
	"assessment_efishery/config"
	"assessment_efishery/entity"
	"assessment_efishery/repository"
	"time"

	"github.com/jinzhu/copier"
)

// TRANSACTION
type ITransactionUsecase interface {
	CreateTransaction(transaction entity.CreateTransactionsRequest) (entity.TransactionsResponse, error)
	GetAllTransaction() ([]entity.Transactions, error)
	GetTransactionById(id int) (entity.TransactionsResponse, error)
	UpdateTransaction(TransactionRequest entity.UpdateTransactionsRequest) (entity.TransactionsResponse, error)
	DeleteTransaction(id int) error
	TransactiontoUser(user_id int) ([]entity.Transactions, error)
}

type TransactionUsecase struct {
	TransactionRepository repository.ITransactionsRepository
}

func NewTransactionUsecase(transactionRepository repository.ITransactionsRepository) *TransactionUsecase {
	return &TransactionUsecase{transactionRepository}
}

func (useCase TransactionUsecase) CreateTransaction(transaction entity.CreateTransactionsRequest) (entity.TransactionsResponse, error) {
	u := entity.Transactions{
		UsersID: transaction.UsersID,
	}
	transactions, err := useCase.TransactionRepository.Store(u)
	if err != nil {
		return entity.TransactionsResponse{}, err
	}
	transactionRes := entity.TransactionsResponse{
		ID:           transactions.ID,
		UsersID:      transactions.UsersID,
		Cart:         transactions.Cart,
		TotalProduct: transactions.TotalProduct,
		Subtotal:     transactions.Subtotal,
		Address:      transactions.Address,
		Created:      transactions.Created,
		Limit:        transactions.Limit,
		Proof:        transactions.Proof,
		Status:       transactions.Status,
	}
	return transactionRes, nil
}

func (useCase TransactionUsecase) GetAllTransaction() ([]entity.Transactions, error) {
	transactions, err := useCase.TransactionRepository.FindAll()
	if err != nil {
		return nil, err
	}
	transactionRes := []entity.Transactions{}
	for _, transaction := range transactions {
		appendTransaction := entity.Transactions{
			ID:           transaction.ID,
			UsersID:      transaction.UsersID,
			Cart:         transaction.Cart,
			TotalProduct: transaction.TotalProduct,
			Subtotal:     transaction.Subtotal,
			Address:      transaction.Address,
			Created:      transaction.Created,
			Limit:        transaction.Limit,
			Proof:        transaction.Proof,
			Status:       transaction.Status,
		}
		transactionRes = append(transactionRes, appendTransaction)
	}
	return transactionRes, nil
}

func (useCase TransactionUsecase) GetTransactionById(id int) (entity.TransactionsResponse, error) {
	transaction, err := useCase.TransactionRepository.FindByID(id)
	if err != nil {
		return entity.TransactionsResponse{}, err
	}
	cartRepository := repository.NewCartRepository(config.DB)
	cartUsecase := NewCartUsecase(cartRepository)
	cart, _ := cartUsecase.CarttoTransaction(transaction.ID)
	userRes := entity.TransactionsResponse{
		ID:           transaction.ID,
		UsersID:      transaction.UsersID,
		Cart:         cart,
		TotalProduct: transaction.TotalProduct,
		Subtotal:     transaction.Subtotal,
		Address:      transaction.Address,
		Created:      transaction.Created,
		Limit:        transaction.Limit,
		Proof:        transaction.Proof,
		Status:       transaction.Status,
	}
	return userRes, nil
}

func (useCase TransactionUsecase) UpdateTransaction(transactionRequest entity.UpdateTransactionsRequest, id int) (entity.TransactionsResponse, error) {
	transaction, err := useCase.TransactionRepository.FindByID(id)
	if err != nil {
		return entity.TransactionsResponse{}, err
	}
	spare := 6 * time.Hour
	// Automaticly add cart(s) to transactions entity
	cartRepository := repository.NewCartRepository(config.DB)
	cartUsecase := NewCartUsecase(cartRepository)
	cart, _ := cartUsecase.CarttoTransaction(transaction.ID)
	// Automaticly calculate subtotal and total of product based on cart
	subtotal := 0
	for _, i := range cart {
		subtotal += i.Total
	}
	productQuantity := 0
	for _, i := range cart {
		productQuantity += i.Quantity
	}
	// Automaticly update transaction status based on payment condition
	new, _ := useCase.GetTransactionById(transaction.ID)
	if new.Created.Valid == true {
		new.Created.Time = new.Created.Time
		new.Limit.Time = new.Limit.Time
	} else if new.Created.Valid == false {
		new.Created.Time = time.Now()
		new.Created.Valid = true
		new.Limit.Time = time.Now().Add(spare)
		new.Limit.Valid = true
	}

	new.Status = "waiting for payment"
	if transactionRequest.Proof.String == "" && new.Proof.Valid == true {
		transactionRequest.Proof.String = new.Proof.String
		transactionRequest.Proof.Valid = true
		new.Status = "accepted"
	} else if transactionRequest.Proof.String == "" {
		transactionRequest.Proof.Valid = false
		if new.Created.Time.After(new.Limit.Time) && new.Proof.Valid == false {
			new.Status = "canceled"
		}
	} else {
		transactionRequest.Proof.Valid = true
		new.Status = "accepted"
	}

	transaction.Cart = cart
	transaction.TotalProduct = productQuantity
	transaction.Subtotal = subtotal
	transaction.Address = transactionRequest.Address
	transaction.Created = new.Created
	transaction.Limit = new.Limit
	transaction.Proof = transactionRequest.Proof
	transaction.Status = new.Status

	copier.CopyWithOption(&transaction, &transactionRequest, copier.Option{IgnoreEmpty: true})
	transaction, _ = useCase.TransactionRepository.Update(transaction)
	transactionRes := entity.TransactionsResponse{}
	copier.Copy(&transactionRes, &transaction)
	return transactionRes, nil
}

func (useCase TransactionUsecase) DeleteTransaction(id int) error {
	_, err := useCase.TransactionRepository.FindByID(id)
	if err != nil {
		return err
	}
	err = useCase.TransactionRepository.Delete(id)
	return err
}

func (useCase TransactionUsecase) TransactiontoUser(users_id int) ([]entity.Transactions, error) {
	transactions, err := useCase.TransactionRepository.TransactionsOfUsers(users_id)
	if err != nil {
		return nil, err
	}
	transactionRes := []entity.Transactions{}
	for _, transaction := range transactions {
		appendTransaction := entity.Transactions{
			ID:           transaction.ID,
			UsersID:      transaction.UsersID,
			Cart:         transaction.Cart,
			TotalProduct: transaction.TotalProduct,
			Subtotal:     transaction.Subtotal,
			Address:      transaction.Address,
			Created:      transaction.Created,
			Limit:        transaction.Limit,
			Proof:        transaction.Proof,
			Status:       transaction.Status,
		}
		transactionRes = append(transactionRes, appendTransaction)
	}
	return transactionRes, nil
}
