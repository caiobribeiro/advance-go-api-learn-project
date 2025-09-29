package routes

import (
	"net/http"

	"github.com/caiobribeiro/advance-go-api-learn-project/internal/handlers"
)

func SetupHealthRoute(serverMux *http.ServeMux, handler *handlers.Handler) {
	serverMux.HandleFunc("/health", handler.HealthHandler())
}
