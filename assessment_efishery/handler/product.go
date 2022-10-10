package handler

import (
	"assessment_efishery/entity"
	"assessment_efishery/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
}

func NewProductHandler(productUsecase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

func (handler ProductHandler) CreateProduct(c echo.Context) error {
	req := entity.CreateProductRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	product, err := handler.productUsecase.CreateProduct(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Create product failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, entity.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "Product created successfully",
		Data:    product,
	})
}

func (handler ProductHandler) GetAllProduct(c echo.Context) error {
	product, err := handler.productUsecase.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get product failed",
			Error:   err.Error(),
		})
	}
	if len(product) == 0 {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Database Empty",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "",
		Data:    product,
	})
}

func (handler ProductHandler) GetProductByID(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	product, err := handler.productUsecase.GetProductById(intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Product not found",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "Display product successfully",
		Data:    product,
	})
}

func (handler ProductHandler) GetProductByCategory(c echo.Context) error {
	product, err := handler.productUsecase.GetProductByCategory(c.Param("kategori"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Product not found",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "Display product successfully",
		Data:    product,
	})
}

func (handler ProductHandler) GetProductByPrice(c echo.Context) error {
	priceMin, _ := strconv.Atoi(c.Param("harga"))
	priceMax, _ := strconv.Atoi(c.Param("harga2"))
	product, err := handler.productUsecase.GetProductByPrice(priceMin, priceMax)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Product not found",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusAccepted,
		Message: "Display product successfully",
		Data:    product,
	})
}

func (handler ProductHandler) UpdateProduct(c echo.Context) error {
	req := entity.UpdateProductRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request Body",
			Error:   err.Error(),
		})
	}
	intId, _ := strconv.Atoi(c.Param("id"))
	product, err := handler.productUsecase.UpdateProduct(req, intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Update product failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Product updated successfully",
		Data:    product,
	})
}

func (handler ProductHandler) DeleteProduct(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	err := handler.productUsecase.DeleteProduct(intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete product failed",
			Error:   err.Error(),
		}))
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Product deleted successfully",
	})
}
