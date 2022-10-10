package usecase

import (
	"assessment_efishery/entity"
	"assessment_efishery/repository"

	"github.com/jinzhu/copier"
)

// PRODUCTS
type IProductUsecase interface {
	CreateProduct(product entity.CreateProductRequest) (entity.ProductResponse, error)
	GetAllProduct() ([]entity.ProductResponse, error)
	GetProductById(id int) (entity.DetailedProductResponse, error)
	UpdateProduct(productRequest entity.UpdateProductRequest) (entity.DetailedProductResponse, error)
	DeleteProduct(id int) error
	GetProductByCategory(kategori string) ([]entity.ProductResponse, error)
	GetProductByPrice(priceMin int, priceMax int) ([]entity.ProductResponse, error)
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

func (useCase ProductUsecase) UpdateProduct(productRequest entity.UpdateProductRequest, id int) (entity.DetailedProductResponse, error) {
	products, err := useCase.ProductRepository.FindByID(id)
	if err != nil {
		return entity.DetailedProductResponse{}, err
	}
	products.Nama = productRequest.Nama
	products.Foto = productRequest.Foto
	products.Harga = productRequest.Harga
	products.Stok = productRequest.Stok
	products.Kategori = productRequest.Kategori
	products.Deskripsi = productRequest.Deskripsi

	copier.CopyWithOption(&products, &productRequest, copier.Option{IgnoreEmpty: true})
	products, _ = useCase.ProductRepository.Update(products)
	productRes := entity.DetailedProductResponse{}
	copier.Copy(&productRes, &products)
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
