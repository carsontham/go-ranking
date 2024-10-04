package bootstrap

import (
	"go-ranking/app/handlers"
	"go-ranking/app/repository"
)

func (s *Server) SetUpRoutes(repo repository.RankingRepository) {
	// TODO
	s.router.Post("/users", handlers.CreateNewUser(repo))
	s.router.Get("/users", handlers.GetAllUser(repo))

}
