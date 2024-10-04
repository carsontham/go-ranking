package handlers

type UserRequestBody struct {
	Name  string `json:"user" validate:"required"`
	Email string `json:"email" validate:"required"`
	Score int64  `json:"score" validate:"required,valid_balance,valid_amount"`
}
