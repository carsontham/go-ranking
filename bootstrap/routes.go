package bootstrap

import "go-ranking/app/handlers"

func (s *Server) SetUpRoutes() {
	// TODO
	s.router.Post("/users", handlers.CreateNewUser())

}
