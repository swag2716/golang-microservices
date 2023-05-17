// Package classification product API.
//
// Documentation for product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/swapnika/golang-microservices/data"
)

// A list of products return im the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type productsNoContent struct {
}

// swagger:parameters updateProducts
type productIDParameterWrapper struct {
	// The id of the product to update in the database
	// in: path
	// required: true
	ID int `json:"id"`
}

type product struct {
}

func NewProduct() *product {
	return &product{}
}

type keyProduct struct{}

func (p *product) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)

		if err != nil {
			fmt.Println("Error deserializing product", err)
			return
		}

		err = prod.Validate()
		if err != nil {
			fmt.Println("Error validating product", err)
			return
		}

		ctx := context.WithValue(r.Context(), keyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
