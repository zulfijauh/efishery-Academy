package usecase

import (
	"assessment_efishery/entity"
	"assessment_efishery/repository"
)

// PRODUCTS
type IProductUsecase interface {
	CreateProduct(product entity.CreateProductRequest) (entity.Products, error)
	GetAllProduct() ([]entity.Products, error)
	GetProductById(id int) (entity.Products, error)
	UpdateProduct(ProductRequest entity.CreateProductRequest) (entity.Products, error)
	DeleteProduct(id int) error
}

type ProductUsecase struct {
	ProductRepository repository.IProductsRepository
}

func NewProductUsecase(productRepository repository.IProductsRepository) *ProductUsecase {
	return &ProductUsecase{productRepository}
}

func (useCase ProductUsecase) CreateProduct(product entity.CreateProductRequest) (entity.ProductResponse, error) {
	p := entity.Products{
		Nama:      product.Nama,
		Foto:      product.Foto,
		Harga:     product.Harga,
		Stok:      product.Stok,
		Kategori:  product.Kategori,
		Deskripsi: product.Deskripsi,
	}
	products, err := useCase.ProductRepository.Store(p)
	if err != nil {
		return entity.ProductResponse{}, err
	}
	productRes := entity.ProductResponse{
		ID:       products.ID,
		Nama:     products.Nama,
		Foto:     products.Foto,
		Harga:    products.Harga,
		Stok:     products.Stok,
		Kategori: products.Kategori,
	}
	return productRes, nil
}

func (useCase ProductUsecase) GetAllProduct() ([]entity.ProductResponse, error) {
	products, err := useCase.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}
	productRes := []entity.ProductResponse{}
	for _, product := range products {
		appendProduct := entity.ProductResponse{
			ID:       product.ID,
			Nama:     product.Nama,
			Foto:     product.Foto,
			Harga:    product.Harga,
			Stok:     product.Stok,
			Kategori: product.Kategori,
		}
		productRes = append(productRes, appendProduct)
	}
	return productRes, nil
}

func (useCase ProductUsecase) GetProductByPrice(priceMin int, priceMax int) ([]entity.ProductResponse, error) {
	products, err := useCase.ProductRepository.FindByPrice(priceMin, priceMax)
	if err != nil {
		return nil, err
	}
	productRes := []entity.ProductResponse{}
	for _, product := range products {
		appendProduct := entity.ProductResponse{
			ID:       product.ID,
			Nama:     product.Nama,
			Foto:     product.Foto,
			Harga:    product.Harga,
			Stok:     product.Stok,
			Kategori: product.Kategori,
		}
		productRes = append(productRes, appendProduct)
	}
	return productRes, nil
}

func (useCase ProductUsecase) GetProductById(id int) (entity.DetailedProductResponse, error) {
	products, err := useCase.ProductRepository.FindByID(id)
	if err != nil {
		return entity.DetailedProductResponse{}, err
	}
	productRes := entity.DetailedProductResponse{
		ID:        products.ID,
		Nama:      products.Nama,
		Foto:      products.Foto,
		Harga:     products.Harga,
		Stok:      products.Stok,
		Deskripsi: products.Deskripsi}
	return productRes, nil
}

func (useCase ProductUsecase) GetProductByCategory(kategori string) ([]entity.ProductResponse, error) {
	products, err := useCase.ProductRepository.FindByCategory(kategori)
	if err != nil {
		return nil, err
	}
	productRes := []entity.ProductResponse{}
	for _, product := range products {
		appendProduct := entity.ProductResponse{
			ID:       product.ID,
			Nama:     product.Nama,
			Foto:     product.Foto,
			Harga:    product.Harga,
			Stok:     product.Stok,
			Kategori: product.Kategori,
		}
		productRes = append(productRes, appendProduct)
	}
	return productRes, nil
}

func (useCase ProductUsecase) UpdateProduct(ProductRequest entity.CreateProductRequest, id int) (entity.Products, error) {
	products, err := useCase.ProductRepository.FindByID(id)
	if err != nil {
		return entity.Products{}, err
	}
	updatedProduct := entity.Products{
		ID:        products.ID,
		Nama:      products.Nama,
		Foto:      products.Foto,
		Harga:     products.Harga,
		Stok:      products.Stok,
		Kategori:  products.Kategori,
		Deskripsi: products.Deskripsi,
	}
	product, err := useCase.ProductRepository.Update(updatedProduct)
	if err != nil {
		return entity.Products{}, err
	}
	productRes := entity.Products{
		Nama:      product.Nama,
		Foto:      product.Foto,
		Harga:     product.Harga,
		Stok:      product.Stok,
		Kategori:  product.Kategori,
		Deskripsi: product.Deskripsi,
	}
	return productRes, nil
}

func (useCase ProductUsecase) DeleteProduct(id int) error {
	_, err := useCase.ProductRepository.FindByID(id)
	if err != nil {
		return err
	}
	err = useCase.ProductRepository.Delete(id)
	return err
}
