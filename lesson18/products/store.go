package products

type ProductStore interface {
	CreateProduct(product Product) (*Product, error)
	ListProduct() ([]Product, error)
	GetProduct(id string) (*Product, error)
	DeleteProduct(id string) error
	UpdateProduct(product Product) (*Product, error)
}
