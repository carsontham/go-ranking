package handlers

import (
	"github.com/go-chi/chi/v5"
	"go-ranking/app/handlers/rest"
	"go-ranking/app/repository"
	"log"
	"net/http"
	"strconv"
)

func GetUserRankByID(repo repository.RankingRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id, _ := strconv.ParseInt(chi.URLParam(req, "id"), 10, 64)
		users, err := repo.GetUserRankByID(id)
		if err != nil {
			log.Println(err)
			rest.InternalServerError(w)
			return
		}
		rest.StatusOK(w, users)

	}
}
