package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"go-ranking/app/handlers/rest"
	"go-ranking/app/repository"
	"log"
	"net/http"
	"strconv"
)

func UpdateUserByID(repo repository.RankingRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id, _ := strconv.ParseInt(chi.URLParam(req, "id"), 10, 64)

		var userReqBody UserRequestBody

		if err := json.NewDecoder(req.Body).Decode(&userReqBody); err != nil {
			log.Println("error body invalid")
			rest.BadRequest(w, errors.New("invalid request body")) // 400
			return
		}

		user, err := repo.GetUserByID(id)
		if err != nil {
			if errors.Is(err, rest.ErrNotFound) {
				rest.NotFound(w)
				return
			}
			rest.InternalServerError(w)
			return
		}

		updatedUser := UpdateUser(user, userReqBody)

		if err := repo.UpdateUserByID(updatedUser); err != nil {
			rest.InternalServerError(w)
		}
		rest.StatusNoContent(w)
	}
}
