package product

import "GolangNorthwindRestApi/helpers"

type Service interface {
	GetProductById(param *getProductByIdRequest) (*Product, error)
	GetProducts(params *getProductsRequest) (*ProductList, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *updateProductRequest) (int64, error)
	DeleteProduct(params *deleteProductRequest) (int64, error)
	GetBestSellers() (*ProductTopResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetProductById(param *getProductByIdRequest) (*Product, error) {
	//Business Logic
	return s.repo.GetProductById(param.ProductID)
}

func (s *service) GetProducts(params *getProductsRequest) (*ProductList, error) {
	products, err := s.repo.GetProducts(params)
	helpers.Catch(err)
	totalProducts, err := s.repo.GetTotalProducts()
	helpers.Catch(err)
	return &ProductList{Data: products, TotalRecords: totalProducts}, nil
}

func (s *service) InsertProduct(params *getAddProductRequest) (int64, error) {
	return s.repo.InsertProduct(params)
}

func (s *service) UpdateProduct(params *updateProductRequest) (int64, error) {
	return s.repo.UpdateProduct(params)
}

func (s *service) DeleteProduct(params *deleteProductRequest) (int64, error) {
	return s.repo.DeleteProduct(params)
}

func (s *service) GetBestSellers() (*ProductTopResponse, error) {
	products, err := s.repo.GetBestSellers()
	helpers.Catch(err)
	total, err := s.repo.GetTotalVentas()
	helpers.Catch(err)

	return &ProductTopResponse{Data: products, TotalVentas: total}, err
}
