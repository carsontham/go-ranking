package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"go-ranking/app/handlers"
	"go-ranking/app/repository"
)

func (s *Server) SetUpRoutes(repo repository.RankingRepository, validator *validator.Validate) {
	s.router.Post("/users", handlers.CreateNewUser(repo, validator))
	s.router.Get("/users", handlers.GetAllUser(repo))

}
