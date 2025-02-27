package handler

import (
	"syka/internal/controller"
	"syka/internal/service"

	"syka/internal/customhttp"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	geoController *controller.AddressController
	service       service.GeoServicer
}

func NewHandler(geoService service.GeoServicer, responder customhttp.Responder) *Handler {
	return &Handler{
		service:       geoService,
		geoController: controller.NewAddressController(geoService, responder),
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/address", func(r chi.Router) {
			r.Post("/search", h.geoController.Search)
			r.Post("/geocode", h.geoController.Geocode)
		})
	})
}
