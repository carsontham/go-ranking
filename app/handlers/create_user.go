package handlers

import (
	"encoding/json"
	"go-ranking/app/handlers/rest"
	"go-ranking/app/repository"
	"net/http"
)

func CreateNewUser(repo repository.RankingRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var userReqBody UserRequestBody

		if err := json.NewDecoder(req.Body).Decode(&userReqBody); err != nil {
			rest.BadRequest(w, err)
			return
		}
		//
		//err := service.CreateNewAccount(account)
		//if err != nil {
		//	log.Println(err)
		//	rest.InternalServerError(w)
		//	return
		//}
		rest.StatusOK(w, "account created")

	}
}
