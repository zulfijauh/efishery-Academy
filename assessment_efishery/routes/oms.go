package routes

import (
	"assessment_efishery/handler"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, userHandler *handler.UserHandler) {
	e.POST("/post/users", userHandler.CreateUser)
	e.GET("/get/users", userHandler.GetAllUser)
	e.GET("/get/users/:id", userHandler.GetUserByID)
	e.PUT("/put/users/:id", userHandler.UpdateUser)
	e.DELETE("/del/users/:id", userHandler.DeleteUser)
}

func ProductRoutes(e *echo.Echo, productHandler *handler.ProductHandler) {
	e.POST("/post/products", productHandler.CreateProduct)
	e.GET("/get/products", productHandler.GetAllProduct)
	e.GET("/get/products/:id", productHandler.GetProductByID)
	e.GET("/get/products/category/:kategori", productHandler.GetProductByCategory)
	e.GET("/get/products/price/:harga/:harga2", productHandler.GetProductByPrice)
	e.PUT("/put/products/:id", productHandler.UpdateProduct)
	e.DELETE("/del/products/:id", productHandler.DeleteProduct)
}

func CartRoutes(e *echo.Echo, cartHandler *handler.CartHandler) {
	e.POST("/post/cart", cartHandler.CreateCart)
	e.GET("/get/cart", cartHandler.GetAllCart)
	e.GET("/get/cart/:id", cartHandler.GetCartByID)
	e.PUT("/put/cart/:id", cartHandler.UpdateCart)
	e.DELETE("/del/cart/:id", cartHandler.DeleteCart)
}

func TransactionRoutes(e *echo.Echo, transactionHandler *handler.TransactionHandler) {
	e.POST("/post/transaction", transactionHandler.CreateTransaction)
	e.GET("/get/transaction", transactionHandler.GetAllTransaction)
	e.GET("/get/transaction/:id", transactionHandler.GetTransactionByID)
	e.PUT("/put/transaction/:id", transactionHandler.UpdateTransaction)
	e.DELETE("/del/transaction/:id", transactionHandler.DeleteTransaction)
}
