package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/caiobribeiro/advance-go-api-learn-project/internal/dtos"
	"github.com/caiobribeiro/advance-go-api-learn-project/internal/store"
	"github.com/caiobribeiro/advance-go-api-learn-project/internal/utils"
)

func (h *Handler) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// create context
		ctx := r.Context()

		// user request aka dto
		var req dtos.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		}

		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "error while hashing password")
			return
		}

		_, err = h.Queries.CreateUser(ctx, store.CreateUserParams{
			Username: req.Username,
			Email:    req.Email,
			Password: hashedPassword,
		})

		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "error creating user")
			return
		}

		utils.RespondWithSuccess(w, http.StatusCreated, "user created", req.Username)

	}
}
