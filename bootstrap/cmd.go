package bootstrap

import "go-ranking/app/repository"

func Run() {
	s := NewServer(":3000")
	repo := repository.NewRankingRepo(GetDB())
	s.SetUpRoutes(repo)
	s.RunServer()
}
