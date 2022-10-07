package usecase

import (
	"assessment_efishery/config"
	"assessment_efishery/entity"
	"assessment_efishery/repository"
)

type ICartUsecase interface {
	CreateCart(cart entity.CreateCartRequest) (entity.CartResponse, error)
	GetAllCart() ([]entity.Cart, error)
	GetCartById(id int) (entity.Cart, error)
	UpdateCart(CartRequest entity.Cart) (entity.Cart, error)
	DeleteCart(id int) error
	CarttoTransaction(id int) ([]entity.Cart, error)
}

type CartUsecase struct {
	CartRepository repository.ICartRepository
}

func NewCartUsecase(cartRepository repository.ICartRepository) *CartUsecase {
	return &CartUsecase{cartRepository}
}

func (useCase CartUsecase) CreateCart(cart entity.CreateCartRequest) (entity.Cart, error) {
	productRepository := repository.NewProductRepository(config.DB)
	productUsecase := NewProductUsecase(productRepository)
	prod, _ := productUsecase.GetProductById(cart.ProductsID)

	p := entity.Cart{
		TransactionsID: cart.TransactionsID,
		ProductsID:     cart.ProductsID,
		Products:       prod.Nama,
		Harga:          prod.Harga,
		Quantity:       cart.Quantity,
		Total:          prod.Harga * cart.Quantity,
	}
	carts, err := useCase.CartRepository.Store(p)
	if err != nil {
		return entity.Cart{}, err
	}
	cartsRes := entity.Cart{
		ID:             carts.ID,
		TransactionsID: carts.TransactionsID,
		ProductsID:     carts.ProductsID,
		Products:       carts.Products,
		Harga:          carts.Harga,
		Quantity:       carts.Quantity,
		Total:          carts.Total,
	}
	return cartsRes, nil
}

func (useCase CartUsecase) GetAllCart() ([]entity.Cart, error) {
	carts, err := useCase.CartRepository.FindAll()
	if err != nil {
		return nil, err
	}
	cartRes := []entity.Cart{}
	for _, cart := range carts {
		appendCart := entity.Cart{
			ID:             cart.ID,
			TransactionsID: cart.TransactionsID,
			ProductsID:     cart.ProductsID,
			Products:       cart.Products,
			Harga:          cart.Harga,
			Quantity:       cart.Quantity,
			Total:          cart.Total,
		}
		cartRes = append(cartRes, appendCart)
	}
	return cartRes, nil
}

func (useCase CartUsecase) GetCartById(id int) (entity.Cart, error) {
	cart, err := useCase.CartRepository.FindByID(id)
	if err != nil {
		return entity.Cart{}, err
	}
	cartRes := entity.Cart{
		ID:             cart.ID,
		TransactionsID: cart.TransactionsID,
		ProductsID:     cart.ProductsID,
		Products:       cart.Products,
		Harga:          cart.Harga,
		Quantity:       cart.Quantity,
		Total:          cart.Total,
	}
	return cartRes, nil
}

func (useCase CartUsecase) UpdateCart(ProductRequest entity.Cart, id int) (entity.Cart, error) {
	cart, err := useCase.CartRepository.FindByID(id)
	if err != nil {
		return entity.Cart{}, err
	}
	productRepository := repository.NewProductRepository(config.DB)
	productUsecase := NewProductUsecase(productRepository)
	prod, _ := productUsecase.GetProductById(cart.ProductsID)
	updatedProduct := entity.Cart{
		TransactionsID: cart.TransactionsID,
		ProductsID:     cart.ProductsID,
		Products:       prod.Nama,
		Harga:          prod.Harga,
		Quantity:       cart.Quantity,
		Total:          prod.Harga * cart.Quantity,
	}
	cart, err = useCase.CartRepository.Update(updatedProduct)
	if err != nil {
		return entity.Cart{}, err
	}
	productRes := entity.Cart{
		TransactionsID: cart.TransactionsID,
		ProductsID:     cart.ProductsID,
		Quantity:       cart.Quantity,
	}
	return productRes, nil
}

func (useCase CartUsecase) DeleteCart(id int) error {
	_, err := useCase.CartRepository.FindByID(id)
	if err != nil {
		return err
	}
	err = useCase.CartRepository.Delete(id)
	return err
}

func (useCase CartUsecase) CarttoTransaction(id int) ([]entity.Cart, error) {
	carts, err := useCase.CartRepository.AllCartID(id)
	if err != nil {
		return nil, err
	}
	cartRes := []entity.Cart{}
	for _, cart := range carts {
		appendProduct := entity.Cart{
			ID:             cart.ID,
			TransactionsID: cart.TransactionsID,
			ProductsID:     cart.ProductsID,
			Products:       cart.Products,
			Harga:          cart.Harga,
			Quantity:       cart.Quantity,
			Total:          cart.Total,
		}
		cartRes = append(cartRes, appendProduct)
	}
	return cartRes, nil
}
