package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/swapnika/webServer/handlers"
)

func main() {

	// l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// hh := handlers.NewHello(l)
	// gg := handlers.NewGoodbye()
	pr := handlers.NewProduct()

	sm := mux.NewRouter()

	// sm.Handle("/", hh)
	// sm.Handle("/goodbye", gg)
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", pr.GetProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", pr.AddProduct)
	postRouter.Use(pr.MiddlewareValidateProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", pr.UpdateProduct)
	putRouter.Use(pr.MiddlewareValidateProduct)
	// sm.Handle("/products/", pr)

	s := &http.Server{
		Addr:         ":9000",
		Handler:      sm,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			fmt.Println("This is the bloddy thing")
			log.Fatal(err)
		}
	}()
	// s.ListenAndServe()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	sig := <-sigChan
	log.Println("Shutting down server", sig)
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	if err := s.Shutdown(tc); err != nil {
		log.Fatal("Server forced to shut down", err)
	}

	fmt.Println("Gracefuuly shut down server")

}
