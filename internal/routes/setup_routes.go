package routes

import (
	"net/http"

	"github.com/caiobribeiro/advance-go-api-learn-project/internal/handlers"
)

func SetupRoutes(serverMux *http.ServeMux, handler *handlers.Handler) {
	SetupHealthRoute(serverMux, handler)
}
