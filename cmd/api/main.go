package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type Product struct {
	ID   int
	Name string
}

func main() {
	r := chi.NewRouter()
	r.Use(myMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		println("endpoint")
	})
	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]string{"message": "sucess"}
		render.JSON(w, r, obj)
	})
	r.Put("/product", func(w http.ResponseWriter, r *http.Request) {
		var product Product
		render.DecodeJSON(r.Body, &product)
		product.ID = 5
		render.JSON(w, r, product)
	})

	http.ListenAndServe(":3000", r)
}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("berofe")
		next.ServeHTTP(w, r)
		println("after")
	})
}
