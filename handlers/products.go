package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/swapnika/webServer/data"
)

type product struct {
}

func NewProduct() *product {
	return &product{}
}

// func (p *product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		p.getProducts(w, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addProduct(w, r)
// 		return
// 	}

// 	if r.Method == http.MethodPut {
// 		reg := regexp.MustCompile(`/products/([0-9]+)`)
// 		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

// 		if len(g) != 1 {
// 			fmt.Println(len(g))
// 			fmt.Println("one")
// 			http.Error(w, "Invalid URL", http.StatusBadRequest)
// 			return
// 		}

// 		fmt.Println("g[0]", g[0])

// 		if len(g[0]) != 2 {
// 			fmt.Println("two")
// 			http.Error(w, "Invalid URL", http.StatusBadRequest)
// 			return
// 		}

// 		idString := g[0][1]
// 		id, err := strconv.Atoi(idString)

// 		if err != nil {
// 			fmt.Println("three")
// 			http.Error(w, "Invalid URL", http.StatusBadRequest)
// 			return
// 		}

// 		fmt.Println("id : ", id)

// 		p.updateProduct(id, w, r)

// 	}

// 	w.WriteHeader(http.StatusMethodNotAllowed)
// }

func (p *product) GetProducts(w http.ResponseWriter, r *http.Request) {

	prod := data.GetProducts()
	err := prod.ToJSON(w)

	if err != nil {
		fmt.Fprintln(w, "Error : ", err)
	}

}

func (p *product) AddProduct(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Post request")

	prod := r.Context().Value(keyProduct{}).(data.Product)

	data.AddProduct(&prod)
}

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
