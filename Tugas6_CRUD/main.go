package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	product = map[int]*Product{}
	nomor   = 1
)

func CreateProduct(c echo.Context) error {
	p := &Product{
		ID: nomor,
	}
	if err := c.Bind(p); err != nil {
		return err
	}
	product[p.ID] = p
	nomor++
	return c.JSON(http.StatusCreated, p)
}

func GetAllProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, product)
}

func GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, product[id])
}

func UpdateProduct(c echo.Context) error {
	p := new(Product)
	if err := c.Bind(p); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	product[id].Name = p.Name
	return c.JSON(http.StatusOK, product[id])
}

func DelProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(product, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()
	e.POST("/", CreateProduct)
	e.GET("/products", GetAllProduct)
	e.GET("/products/:id", GetProduct)
	e.PUT("/update/:id", UpdateProduct)
	e.DELETE("/delete/:id", DelProduct)

	e.Logger.Fatal(e.Start(":1323"))

}
