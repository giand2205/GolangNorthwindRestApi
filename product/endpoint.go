package product

import (
	"GolangNorthwindRestApi/helpers"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type getProductByIdRequest struct {
	ProductID int
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

type updateProductRequest struct {
	ID           int64
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

type deleteProductRequest struct {
	ProductID string
}

type getBestSellersRequest struct {
}

func makeGetProductByIdEndpoint(s Service) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIdRequest)
		product, err := s.GetProductById(&req)
		helpers.Catch(err)
		return product, nil
	}
	return getProductByIdEndpoint
}

func makeGetProductsEndpoint(s Service) endpoint.Endpoint {
	getProductsEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		products, err := s.GetProducts(&req)
		helpers.Catch(err)
		return products, nil
	}
	return getProductsEndpoint
}

func makeAddProductEndpoint(s Service) endpoint.Endpoint {
	addProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		productId, err := s.InsertProduct(&req)
		helpers.Catch(err)
		return productId, nil
	}
	return addProductEndpoint
}

func makeUpdateProductEndpoint(s Service) endpoint.Endpoint {
	updateProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductRequest)
		productId, err := s.UpdateProduct(&req)
		helpers.Catch(err)
		return productId, nil
	}
	return updateProductEndpoint
}

func makeDeleteProductEndpoint(s Service) endpoint.Endpoint {
	deleteProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)
		productId, err := s.DeleteProduct(&req)
		helpers.Catch(err)
		return productId, nil
	}
	return deleteProductEndpoint
}

func makeBestSellersEndpoint(s Service) endpoint.Endpoint {
	getBestSellersEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.GetBestSellers()
		helpers.Catch(err)
		return result, nil
	}
	return getBestSellersEndpoint
}
