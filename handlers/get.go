package handlers

import (
	"fmt"
	"net/http"

	"github.com/swapnika/golang-microservices/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 	200: productsResponse

// GetProducts return the products from the data store
func (p *product) GetProducts(w http.ResponseWriter, r *http.Request) {

	prod := data.GetProducts()
	err := prod.ToJSON(w)

	if err != nil {
		fmt.Fprintln(w, "Error : ", err)
	}

}
