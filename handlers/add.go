package handlers

import (
	"fmt"
	"net/http"

	"github.com/swapnika/golang-microservices/data"
)

func (p *product) AddProduct(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Post request")

	prod := r.Context().Value(keyProduct{}).(data.Product)

	data.AddProduct(&prod)
}
