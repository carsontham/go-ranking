package handlers

import "go-ranking/app/repository"

type UserRequestBody struct {
	ID    int64  `json:"id"`
	Name  string `json:"name" validate:"required,min=1"`
	Email string `json:"email" validate:"required,email"`
	Score int64  `json:"score" validate:"required,gt=0"`
}

type RankedUserResponse struct {
	UserRequestBody
	Rank int64 `json:"rank"`
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
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Score: user.Score,
	}
}

func UserDBModelArrayToViewModelArray(users []*repository.User) []*UserRequestBody {
	viewArray := make([]*UserRequestBody, len(users))
	for index, user := range users {
		viewArray[index] = UserDBModelToViewModel(user)
	}
	return viewArray
}

func UpdateUser(user *repository.User, updatedUser UserRequestBody) *repository.User {
	return &repository.User{
		ID:    user.ID,
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
		Score: updatedUser.Score,
	}
}

func RankedUserDBModelToViewModel(rankedUser *repository.RankedUser) *RankedUserResponse {
	return &RankedUserResponse{
		UserRequestBody: *UserDBModelToViewModel(&rankedUser.User),
		Rank:            rankedUser.Rank,
	}
}

func RankedUserDBModelArrayToViewModelArray(rankedUsers []*repository.RankedUser) []*RankedUserResponse {
	viewArray := make([]*RankedUserResponse, len(rankedUsers))
	for index, rankedUser := range rankedUsers {
		viewArray[index] = RankedUserDBModelToViewModel(rankedUser)
	}
	return viewArray
}
