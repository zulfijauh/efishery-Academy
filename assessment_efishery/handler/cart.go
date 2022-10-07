package handler

import (
	"assessment_efishery/entity"
	"assessment_efishery/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartUsecase *usecase.CartUsecase
}

func NewCartHandler(cartUsecase *usecase.CartUsecase) *CartHandler {
	return &CartHandler{cartUsecase: cartUsecase}
}

func (handler CartHandler) CreateCart(c echo.Context) error {
	req := entity.CreateCartRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	cart, err := handler.cartUsecase.CreateCart(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Create cart failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, entity.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "Cart created successfully",
		Data:    cart,
	})
}

func (handler CartHandler) GetAllCart(c echo.Context) error {
	cart, err := handler.cartUsecase.GetAllCart()
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get cart failed",
			Error:   err.Error(),
		})
	}
	if len(cart) == 0 {
		fmt.Println("Database Empty")
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "",
		Data:    cart,
	})
}

func (handler CartHandler) GetCartByID(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	cart, err := handler.cartUsecase.GetCartById(intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Create cart failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "Cart created successfully",
		Data:    cart,
	})
}

func (handler CartHandler) UpdateCart(c echo.Context) error {
	req := entity.Cart{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request Body",
			Error:   err.Error(),
		})
	}
	intId, _ := strconv.Atoi(c.Param("id"))
	cart, err := handler.cartUsecase.UpdateCart(req, intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Update cart failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Cart updated successfully",
		Data:    cart,
	})
}

func (handler CartHandler) DeleteCart(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	err := handler.cartUsecase.DeleteCart(intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete cart failed",
			Error:   err.Error(),
		}))
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Cart deleted successfully",
	})
}
