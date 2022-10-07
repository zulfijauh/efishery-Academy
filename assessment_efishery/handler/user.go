package handler

import (
	"assessment_efishery/entity"
	"assessment_efishery/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// USER
type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecaes *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecaes}
}

// RegisterUser godoc
// @Summary Register User Account.
// @Description registering a user.
// @Tags User Auth
// @Param Body body entity.CreateUserRequest true "the body to register a FULL ACCESS account (USER)"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users [post]
func (handler UserHandler) CreateUser(c echo.Context) error {
	req := entity.Users{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	user, err := handler.userUsecase.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Create user failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, entity.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "User created successfully",
		Data:    user,
	})
}

func (handler UserHandler) GetAllUser(c echo.Context) error {
	user, err := handler.userUsecase.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get user failed",
			Error:   err.Error(),
		})
	}
	if len(user) == 0 {
		fmt.Println("Database Empty")
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "",
		Data:    user,
	})
}

func (handler UserHandler) GetUserByID(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	user, err := handler.userUsecase.GetUserById(intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Create user failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "User created successfully",
		Data:    user,
	})
}

func (handler UserHandler) UpdateUser(c echo.Context) error {
	req := entity.UpdateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request Body",
			Error:   err.Error(),
		})
	}
	intId, _ := strconv.Atoi(c.Param("id"))
	user, err := handler.userUsecase.UpdateUser(req, intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Update user failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "User updated successfully",
		Data:    user,
	})
}

func (handler UserHandler) DeleteUser(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	err := handler.userUsecase.DeleteUser(intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete user failed",
			Error:   err.Error(),
		}))
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "User deleted successfully",
	})
}
