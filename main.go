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

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/swapnika/golang-microservices/handlers"
)

func main() {

	pr := handlers.NewProduct()

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", pr.GetProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", pr.AddProduct)
	postRouter.Use(pr.MiddlewareValidateProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", pr.UpdateProduct)
	putRouter.Use(pr.MiddlewareValidateProduct)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

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
