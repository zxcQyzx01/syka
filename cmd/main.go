package main

import (
	"fmt"
	"log"
	"net/http"
	"syka/internal/config"
	"syka/internal/customhttp"
	"syka/internal/handler"
	"syka/internal/infrastructure/dadata"
	"syka/internal/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	config := config.LoadConfig()
	fmt.Println(config)

	provider := dadata.NewProvider(config.ApiKey, config.SecretKey)
	fmt.Println(provider)

	geoService := service.NewGeoService(provider)
	responder := customhttp.NewResponder()
	h := handler.NewHandler(geoService, responder)

	h.RegisterRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
