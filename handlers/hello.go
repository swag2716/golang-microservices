package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Server is running at port 9000")

	d, _ := io.ReadAll(r.Body)

	log.Printf("Data is :%s", d)

	fmt.Fprintf(w, "data is %s", d)
}
