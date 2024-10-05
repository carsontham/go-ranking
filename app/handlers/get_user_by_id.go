package handlers

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"go-ranking/app/handlers/rest"
	"go-ranking/app/repository"
	"net/http"
	"strconv"
)

func GetUserByID(repo repository.RankingRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id, _ := strconv.ParseInt(chi.URLParam(req, "id"), 10, 64)

		user, err := repo.GetUserByID(id)
		if err != nil {
			if errors.Is(err, rest.ErrNotFound) {
				rest.NotFound(w)
				return
			}
			rest.InternalServerError(w)
			return
		}
		rest.StatusOK(w, user)
	}
}
