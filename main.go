package main

import (
	"api/configs"
	"api/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
