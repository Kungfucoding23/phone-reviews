package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

//MyServer server struct
type MyServer struct {
	server *http.Server
}

//NewServer is a new custom Server
func NewServer(mux *chi.Mux) *http.Server {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}

//Run the server
func (s *MyServer) Run() {
	log.Fatal(s.server.ListenAndServe())
}
