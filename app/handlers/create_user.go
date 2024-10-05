package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"go-ranking/app/handlers/rest"
	"go-ranking/app/repository"
	"log"
	"net/http"
)

func CreateNewUser(repo repository.RankingRepository, v *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var userReqBody UserRequestBody

		if err := json.NewDecoder(req.Body).Decode(&userReqBody); err != nil {
			log.Println("error body invalid")
			rest.BadRequest(w, errors.New("invalid request body")) // 400
			return
		}

		if err := v.Struct(userReqBody); err != nil {
			if ve, ok := err.(validator.ValidationErrors); ok {
				rest.UnprocessableEntity(w, ve) // 422
			} else {
				rest.InternalServerError(w) // 500
			}
			return
		}

		dbUser := UserViewModelToDBModel(userReqBody)

		// Check if email is unique
		isUnique, err := repo.CheckUniqueEmail(dbUser.Email)
		if err != nil {
			log.Println("Database error:", err)
			rest.InternalServerError(w) // 500
			return
		}
		if !isUnique {
			// status 409
			rest.StatusConflict(w, "Email is already in use") // 409
			return
		}

		if err := repo.CreateNewUser(dbUser); err != nil {
			rest.InternalServerError(w) // 500
			return
		}
		rest.StatusCreated(w) // 201

	}
}
