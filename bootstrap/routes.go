package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"go-ranking/app/handlers"
	"go-ranking/app/repository"
)

func (s *Server) SetUpRoutes(repo repository.RankingRepository, validator *validator.Validate) {
	s.router.Post("/users", handlers.CreateNewUser(repo, validator))
	s.router.Get("/users", handlers.GetAllUser(repo))
	s.router.Get("/users/{id:\\d+}", handlers.GetUserByID(repo))
	s.router.Post("/users/{id:\\d+}", handlers.UpdateUserByID(repo, validator))
	s.router.Delete("/users/{id:\\d+}", handlers.DeleteUserByID(repo))

	s.router.Get("/users/rank", handlers.GetAllUserRank(repo))
	s.router.Get("/users/rank/{id:\\d+}", handlers.GetUserRankByID(repo))
}
