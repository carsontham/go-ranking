package handlers

import (
	"go-ranking/app/handlers/rest"
	"go-ranking/app/repository"
	"log"
	"net/http"
)

func GetAllUserRank(repo repository.RankingRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		users, err := repo.GetAllUserRanking()
		if err != nil {
			log.Println(err)
			rest.InternalServerError(w)
			return
		}
		rest.StatusOK(w, RankedUserDBModelArrayToViewModelArray(users))

	}
}
