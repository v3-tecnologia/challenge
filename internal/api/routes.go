package api

import (
	"v3/internal/api/handlers"
	"v3/pkg/httpcore"

	"github.com/go-chi/chi/v5"
)

func applyRoutes(router chi.Router, ctl *handlers.ApiController) {
	router.Post("/telemetry/gyroscope", httpcore.Handle(ctl.CreateGyroscope))

}
