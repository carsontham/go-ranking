package handlers

import (
	"go-ranking/app/handlers/rest"
	"go-ranking/app/repository"
	"log"
	"net/http"
)

func GetAllUser(repo repository.RankingRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//var userReqBody UserRequestBody
		//
		//if err := json.NewDecoder(req.Body).Decode(&userReqBody); err != nil {
		//	rest.BadRequest(w, err)
		//	return
		//}
		users, err := repo.GetAllUser()
		if err != nil {
			log.Println(err)
			rest.InternalServerError(w)
			return
		}
		rest.StatusOK(w, users)

	}
}
