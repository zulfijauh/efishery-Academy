package main

import (
	"fmt"
	"tugas_week1/product"
)

func main() {
	list := product.GetProduct()
	uang := 100000
	remainedMoney := uang

	listBelanja := list.BuyProducts(&remainedMoney)
	fmt.Println(listBelanja, "\n")

	// 1B
	mostExpensiveProduct := list.ReturnMax()
	fmt.Println(mostExpensiveProduct, "\n")

	cheapestProduct := list.ReturnMin()
	fmt.Println(cheapestProduct, "\n")

	ceban := list.SepuluhRibu()
	fmt.Println(ceban, "\n")

}
