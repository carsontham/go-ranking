package repository

// RankingRepository is the interface for repository
type RankingRepository interface {
	CreateNewUser(User) error
	GetAllUser(bool, int) ([]*User, error)
	CheckUniqueEmail(string) (bool, error)
}
