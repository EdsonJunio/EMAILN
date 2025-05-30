package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaignService: campaignService,
	}
	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns", endpoints.HandlerError(handler.CampaignGet))

	http.ListenAndServe(":3000", r)
}
