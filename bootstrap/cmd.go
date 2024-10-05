package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"go-ranking/app/repository"
)

func Run() {
	s := NewServer(":3000")
	repo := repository.NewRankingRepo(GetDB())
	val := validator.New()
	s.SetUpRoutes(repo, val)
	s.RunServer()
}
