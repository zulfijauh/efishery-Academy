package usecase

import (
	"assessment_efishery/config"
	"assessment_efishery/entity"
	"assessment_efishery/repository"

	"github.com/jinzhu/copier"
)

type ICartUsecase interface {
	CreateCart(cart entity.CreateCartRequest) (entity.CartResponse, error)
	GetAllCart() ([]entity.CartResponse, error)
	GetCartById(id int) (entity.CartResponse, error)
	UpdateCart(cartRequest entity.UpdateCartRequest) (entity.CartResponse, error)
	DeleteCart(id int) error
	CarttoTransaction(id int) ([]entity.Cart, error)
}

type CartUsecase struct {
	CartRepository repository.ICartRepository
}

func NewCartUsecase(cartRepository repository.ICartRepository) *CartUsecase {
	return &CartUsecase{cartRepository}
}

func (useCase CartUsecase) CreateCart(cart entity.CreateCartRequest) (entity.CartResponse, error) {
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
		return entity.CartResponse{}, err
	}
	cartsRes := entity.CartResponse{
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

func (useCase CartUsecase) GetAllCart() ([]entity.CartResponse, error) {
	carts, err := useCase.CartRepository.FindAll()
	if err != nil {
		return nil, err
	}
	cartRes := []entity.CartResponse{}
	for _, cart := range carts {
		appendCart := entity.CartResponse{
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

func (useCase CartUsecase) GetCartById(id int) (entity.CartResponse, error) {
	cart, err := useCase.CartRepository.FindByID(id)
	if err != nil {
		return entity.CartResponse{}, err
	}
	cartRes := entity.CartResponse{
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

func (useCase CartUsecase) UpdateCart(cartRequest entity.UpdateCartRequest, id int) (entity.CartResponse, error) {
	cart, err := useCase.CartRepository.FindByID(id)
	if err != nil {
		return entity.CartResponse{}, err
	}
	productRepository := repository.NewProductRepository(config.DB)
	productUsecase := NewProductUsecase(productRepository)
	prod, _ := productUsecase.GetProductById(cart.ProductsID)

	cart.ProductsID = cartRequest.ProductsID
	cart.Products = prod.Nama
	cart.Harga = prod.Harga
	cart.Quantity = cartRequest.Quantity
	cart.Total = prod.Harga * cart.Quantity

	copier.CopyWithOption(&cart, &cartRequest, copier.Option{IgnoreEmpty: true})
	cart, _ = useCase.CartRepository.Update(cart)
	cartRes := entity.CartResponse{}
	copier.Copy(&cartRes, &cart)
	return cartRes, nil
}

func (useCase CartUsecase) DeleteCart(id int) error {
	_, err := useCase.CartRepository.FindByID(id)
	if err != nil {
		return err
	}
	err = useCase.CartRepository.Delete(id)
	return err
}

func (useCase CartUsecase) CarttoTransaction(transactions_id int) ([]entity.Cart, error) {
	carts, err := useCase.CartRepository.AllCartID(transactions_id)
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
