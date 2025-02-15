package main

import (
	"emailn/internal/domain/campaign"
	"github.com/go-playground/validator/v10"
)

func main() {
	contacts := []campaign.Contact{{Email: "admin@example.com"}, {Email: ""}}
	campaign := campaign.Campaign{Contacts: contacts}
	validate := validator.New()
	err := validate.Struct(campaign)
	if err == nil {
		println("Nenhum erro")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			switch v.Tag() {
			case "required":
				println(v.StructField() + " is required")
			case "min":
				println(v.StructField() + " is required with min " + v.Param())
			case "max":
				println(v.StructField() + " is required with max " + v.Param())
			case "email":
				println(v.StructField() + " is email" + v.Param())
			}
		}
	}
}
