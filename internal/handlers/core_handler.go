package handlers

import (
	"database/sql"

	"github.com/caiobribeiro/advance-go-api-learn-project/internal/store"
)

type Handler struct {
	//  DB instance
	DB *sql.DB
	// Query stores
	Queries *store.Queries
}

func NewHandlers(db *sql.DB, queries *store.Queries) *Handler {
	return &Handler{
		DB:      db,
		Queries: queries,
	}
}
