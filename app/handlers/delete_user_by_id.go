package handlers

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"go-ranking/app/handlers/rest"
	"go-ranking/app/repository"
	"log"
	"net/http"
	"strconv"
)

func DeleteUserByID(repo repository.RankingRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(req, "id"), 10, 64)
		if err != nil {
			log.Println("Invalid user ID:", err)
			rest.BadRequest(w, errors.New("invalid user ID")) // 400
			return
		}

		// Check if the user exists
		user, err := repo.GetUserByID(id)
		if err != nil {
			if errors.Is(err, rest.ErrNotFound) {
				rest.NotFound(w)
				return
			}
			rest.InternalServerError(w)
			return
		}

		// User exists, proceed to delete
		if err := repo.DeleteUserByID(user.ID); err != nil {
			log.Println("Error deleting user:", err)
			rest.InternalServerError(w)
			return
		}

		// Successfully deleted user
		rest.StatusNoContent(w) // 204 No Content
	}
}
