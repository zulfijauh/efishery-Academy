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
	CreateTransaction(transaction entity.Transactions) (entity.TransactionsResponse, error)
	GetAllTransaction() ([]entity.Transactions, error)
	GetTransactionById(id int) (entity.TransactionsResponse, error)
	UpdateTransaction(TransactionRequest entity.UpdateTransactionsRequest) (entity.TransactionsResponse, error)
	DeleteTransaction(id int) error
}

type TransactionUsecase struct {
	TransactionRepository repository.ITransactionsRepository
}

func NewTransactionUsecase(transactionRepository repository.ITransactionsRepository) *TransactionUsecase {
	return &TransactionUsecase{transactionRepository}
}

func (useCase TransactionUsecase) CreateTransaction(transaction entity.Transactions) (entity.TransactionsResponse, error) {
	spare := 6 * time.Hour
	u := entity.Transactions{
		UsersID:         transaction.UsersID,
		ShippingAddress: transaction.ShippingAddress,
		Created:         time.Now(),
		Limit:           time.Now().Add(spare),
		Status:          "Waiting payment",
	}
	transactions, err := useCase.TransactionRepository.Store(u)
	if err != nil {
		return entity.TransactionsResponse{}, err
	}
	transactionRes := entity.TransactionsResponse{
		ID:              transactions.ID,
		UsersID:         transactions.UsersID,
		Cart:            transactions.Cart,
		TotalProduct:    transactions.TotalProduct,
		Subtotal:        transactions.Subtotal,
		ShippingAddress: transactions.ShippingAddress,
		Created:         transactions.Created,
		Limit:           transactions.Limit,
		Proof:           transactions.Proof,
		Status:          transactions.Status,
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
			ID:              transaction.ID,
			UsersID:         transaction.UsersID,
			Cart:            transaction.Cart,
			TotalProduct:    transaction.TotalProduct,
			Subtotal:        transaction.Subtotal,
			ShippingAddress: transaction.ShippingAddress,
			Created:         transaction.Created,
			Limit:           transaction.Limit,
			Proof:           transaction.Proof,
			Status:          transaction.Status,
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
	userRes := entity.TransactionsResponse{
		ID:              transaction.ID,
		UsersID:         transaction.UsersID,
		Cart:            transaction.Cart,
		TotalProduct:    transaction.TotalProduct,
		Subtotal:        transaction.Subtotal,
		ShippingAddress: transaction.ShippingAddress,
		Created:         transaction.Created,
		Limit:           transaction.Limit,
		Proof:           transaction.Proof,
		Status:          transaction.Status,
	}
	return userRes, nil
}

func (useCase TransactionUsecase) UpdateTransaction(transactionRequest entity.UpdateTransactionsRequest, id int) (entity.TransactionsResponse, error) {
	transaction, err := useCase.TransactionRepository.FindByID(id)
	if err != nil {
		return entity.TransactionsResponse{}, err
	}
	new_status, _ := useCase.GetTransactionById(transaction.ID)
	if new_status.Created.After(new_status.Limit) && &new_status.Proof == nil {
		new_status.Status = "canceled"
	} else if new_status.Created.Before(new_status.Limit) && &new_status.Proof != nil {
		new_status.Status = "accepted"
	}
	cartRepository := repository.NewCartRepository(config.DB)
	cartUsecase := NewCartUsecase(cartRepository)
	cart, _ := cartUsecase.CarttoTransaction(transaction.ID)

	subtotal := 0
	for _, i := range cart {
		subtotal += i.Total
	}
	productQuantity := 0
	for _, i := range cart {
		productQuantity += i.Quantity
	}

	transaction.Cart = cart
	transaction.TotalProduct = productQuantity
	transaction.Subtotal = subtotal
	transaction.Proof = transactionRequest.Proof
	transaction.Status = new_status.Status

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
