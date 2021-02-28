package main

import (
	"encoding/json"
	"net/http"

	"github.com/Kungfucoding23/API-Go-mysql-chi/gadgets/smartphones/web"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

//Routes ...
func Routes(sph *web.CreateSmartphoneHandler) *chi.Mux {
	mux := chi.NewMux()

	//global middlewares
	mux.Use(
		middleware.Logger,    //log every http request
		middleware.Recoverer, // recover if a panic occurs
	)

	mux.Post("/smartphones", sph.SaveSmartphoneHandler)
	mux.Get("/hello", helloHandler)

	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("done-by", "alejandro")

	res := map[string]interface{}{"message": "hello world"}
	_ = json.NewEncoder(w).Encode(res)
}
