package repository

// RankingRepository
type RankingRepository interface {
	GetAllUser() ([]*User, error)
}
