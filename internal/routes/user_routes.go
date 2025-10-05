package routes

import (
	"net/http"

	"github.com/caiobribeiro/advance-go-api-learn-project/internal/handlers"
)

func SetupUserRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("POST /user/register", handler.CreateUserHandler())
}
