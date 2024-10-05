package handlers

import (
	"errors"
	"go-ranking/app/handlers/rest"
	"go-ranking/app/repository"
	"log"
	"net/http"
	"strconv"
)

func GetAllUser(repo repository.RankingRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		// query params allowed - sortDesc and minScore
		sortByScoreDesc := req.URL.Query().Get("sortDesc")
		minScoreStr := req.URL.Query().Get("minScore")

		minScore := 0
		if minScoreStr != "" {
			score, err := strconv.Atoi(minScoreStr)
			if err != nil {
				log.Println("Invalid min_score value:", err)
				rest.BadRequest(w, errors.New("invalid min score"))
				return
			}
			minScore = score
		}

		sortScoreDesc := sortByScoreDesc == "true"

		users, err := repo.GetAllUser(sortScoreDesc, minScore)
		if err != nil {
			log.Println(err)
			rest.InternalServerError(w)
			return
		}
		rest.StatusOK(w, users)

	}
}
