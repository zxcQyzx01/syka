package main

import (
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
	// Загрузка конфигурации
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализация провайдера
	provider, err := initProvider(config)
	if err != nil {
		log.Fatalf("Ошибка инициализации провайдера: %v", err)
	}

	// Инициализация сервисов и обработчиков
	geoService := service.NewGeoService(provider)
	responder := customhttp.NewResponder()
	h := handler.NewHandler(geoService, responder)

	// Регистрация маршрутов
	r := chi.NewRouter()
	registerRoutes(r, h)

	// Запуск сервера
	log.Fatal(http.ListenAndServe(":8080", r))
}

func loadConfig() (*config.Config, error) {
	return config.LoadConfig(), nil
}

func initProvider(config *config.Config) (*dadata.Provider, error) {
	return dadata.NewProvider(config.ApiKey, config.SecretKey), nil
}

func registerRoutes(r *chi.Mux, h *handler.Handler) {
	h.RegisterRoutes(r)
}
