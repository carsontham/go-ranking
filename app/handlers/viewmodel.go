package handlers

import "go-ranking/app/repository"

type UserRequestBody struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required,min=1"`
	Email string `json:"email" validate:"required,email"`
	Score int64  `json:"score" validate:"required,gt=0"`
}

func UserViewModelToDBModel(user UserRequestBody) *repository.User {
	return &repository.User{
		Name:  user.Name,
		Email: user.Email,
		Score: user.Score,
	}
}

func UserDBModelToViewModel(user *repository.User) *UserRequestBody {
	return &UserRequestBody{
		Name:  user.Name,
		Email: user.Email,
		Score: user.Score,
	}
}
