package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/swapnika/golang-microservices/data"
)

// swagger:route PUT /products/{id} products updateProducts
// responses:
// 	201: noContent

// UpdateProducts update a produst in the data store with the specific id
func (p *product) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("put invalid")
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}
	fmt.Println("Put request")

	prod := r.Context().Value(keyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not Found", http.StatusNotFound)
		return
	}
}
