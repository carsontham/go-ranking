package handlers

import "go-ranking/app/repository"

type UserRequestBody struct {
	ID    int64  `json:"id"`
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

func UpdateUser(user *repository.User, updatedUser UserRequestBody) *repository.User {
	return &repository.User{
		ID:    user.ID,
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
		Score: updatedUser.Score,
	}
}
