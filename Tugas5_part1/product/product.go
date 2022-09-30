package product

import "sort"

type Product struct {
	ID     int
	Barang string
	Harga  int
}

type Products []Product

var listProduct = Products{
	{
		ID:     1,
		Barang: "Benih Lele",
		Harga:  50000,
	},
	{
		ID:     2,
		Barang: "Pakan lele cap menara",
		Harga:  25000,
	},
	{
		ID:     3,
		Barang: "Probiotik A",
		Harga:  75000,
	},
	{
		ID:     4,
		Barang: "Probiotik Nila B",
		Harga:  10000,
	},
	{
		ID:     5,
		Barang: "Pakan Nila",
		Harga:  20000,
	},
	{
		ID:     6,
		Barang: "Benih Nila",
		Harga:  20000,
	},
	{
		ID:     7,
		Barang: "Cupang",
		Harga:  5000,
	},
	{
		ID:     8,
		Barang: "Benih Nila",
		Harga:  30000,
	},
	{
		ID:     9,
		Barang: "Benih Cupang",
		Harga:  10000,
	},
	{
		ID:     10,
		Barang: "Probiotik B",
		Harga:  10000,
	},
}

// NOMOR 1 A
func (products Products) SortCheapest() Products {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Harga < products[j].Harga
	})
	return products
}

func (product Products) BuyProducts(remainedMoney *int) Products {
	var result Products

	product.SortCheapest()

	for _, product := range listProduct {
		if *remainedMoney >= product.Harga {
			*remainedMoney -= product.Harga
			result = append(result, product)
		} else {
			break
		}
	}
	return result
}

func GetProduct() Products {
	return listProduct
}

// NOMOR 1 B
func (products Products) ReturnMax() Product {
	result := products[0]

	for i := 1; i < len(products); i++ {
		if products[i].Harga > result.Harga {
			result = products[i]
		}
	}
	return result
}

func (products Products) ReturnMin() Product {
	result := products[0]

	for i := 1; i < len(products); i++ {
		if products[i].Harga < result.Harga {
			result = products[i]
		}
	}
	return result
}

func (products Products) SepuluhRibu() Products {
	var result Products

	for _, product := range listProduct {
		if product.Harga == 10000 {
			result = append(result, product)
		}
	}
	return result
}
