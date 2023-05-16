package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	// l *log.Logger
}

func NewGoodbye() *Goodbye {
	return &Goodbye{}
}

func (h *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Server is running at port :9090")

	d, _ := io.ReadAll(r.Body)

	fmt.Fprintln(w, "data is : ", d)
}
