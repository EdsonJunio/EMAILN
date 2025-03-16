package main

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"emailn/internal/internalErrors"
	"encoding/json"
	"errors"
	"fmt"
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
	service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

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
		json.NewDecoder(r.Body).Decode(&product)
		fmt.Println(product)
		render.JSON(w, r, product)
	})
	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var request contract.NewCampaign
		render.DecodeJSON(r.Body, &request)
		id, err := service.Create(request)

		if err != nil {

			if errors.Is(err, internalErrors.ErrInternal) {
				render.Status(r, 500)
			} else {
				render.Status(r, 300)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		render.Status(r, 201)
		render.JSON(w, r, map[string]string{"id": id})
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
