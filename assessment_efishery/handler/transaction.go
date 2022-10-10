package handler

import (
	"assessment_efishery/entity"
	"assessment_efishery/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionUsecase *usecase.TransactionUsecase
}

func NewTransactionHandler(transactionUsecase *usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{transactionUsecase: transactionUsecase}
}

func (handler TransactionHandler) CreateTransaction(c echo.Context) error {
	req := entity.CreateTransactionsRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	transaction, err := handler.transactionUsecase.CreateTransaction(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Create transaction failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, entity.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "Transaction created successfully",
		Data:    transaction,
	})
}

func (handler TransactionHandler) GetAllTransaction(c echo.Context) error {
	transaction, err := handler.transactionUsecase.GetAllTransaction()
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get transaction failed",
			Error:   err.Error(),
		})
	}
	if len(transaction) == 0 {
		fmt.Println("Database Empty")
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "",
		Data:    transaction,
	})
}

func (handler TransactionHandler) GetTransactionByID(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	transaction, err := handler.transactionUsecase.GetTransactionById(intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Create transaction failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "Transaction created successfully",
		Data:    transaction,
	})
}

func (handler TransactionHandler) UpdateTransaction(c echo.Context) error {
	req := entity.UpdateTransactionsRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request Body",
			Error:   err.Error(),
		})
	}
	intId, _ := strconv.Atoi(c.Param("id"))
	transaction, err := handler.transactionUsecase.UpdateTransaction(req, intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Update transaction failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Transaction updated successfully",
		Data:    transaction,
	})
}

func (handler TransactionHandler) DeleteTransaction(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	err := handler.transactionUsecase.DeleteTransaction(intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete transaction failed",
			Error:   err.Error(),
		}))
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Transaction deleted successfully",
	})
}
