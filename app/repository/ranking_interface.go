package repository

//go:generate mockgen -source=ranking_interface.go -package repositorytest -destination ../../tests/repositorytest/ranking_repo_mock.go

// RankingRepository is the interface for repository
type RankingRepository interface {
	CreateNewUser(*User) error
	GetAllUser(bool, int) ([]*User, error)
	CheckUniqueEmail(string) (bool, error)
	GetUserByID(int64) (*User, error)
	UpdateUserByID(*User) error
	DeleteUserByID(int64) error
	GetAllUserRanking() ([]RankedUser, error)
	GetUserRankByID(int64) (*RankedUser, error)
}
