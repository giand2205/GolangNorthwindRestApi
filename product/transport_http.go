package product

import (
	"GolangNorthwindRestApi/helpers"
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
	"strconv"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductByIdHandler := kithttp.NewServer(makeGetProductByIdEndpoint(s), getProductByIdRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductByIdHandler)
	getProductsHandler := kithttp.NewServer(makeGetProductsEndpoint(s), getProductsRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)
	addProductHandler := kithttp.NewServer(makeAddProductEndpoint(s), addProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addProductHandler)
	updateProductHandler := kithttp.NewServer(makeUpdateProductEndpoint(s), updateProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/{id}", updateProductHandler)
	deleteProductHandler := kithttp.NewServer(makeDeleteProductEndpoint(s), deleteProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteProductHandler)
	getBestSellerHandler := kithttp.NewServer(makeBestSellersEndpoint(s), getBestSellersRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/bestSellers", getBestSellerHandler)
	return r
}

func getProductByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIdRequest{
		ProductID: productId,
	}, nil
}

func getProductsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helpers.Catch(err)
	return request, nil
}

func addProductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helpers.Catch(err)
	return request, nil
}

func updateProductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := updateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helpers.Catch(err)
	return request, nil
}

func deleteProductRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	productId := chi.URLParam(r, "id")
	return deleteProductRequest{
		ProductID: productId,
	}, nil
}

func getBestSellersRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return getBestSellersRequest{}, nil
}
