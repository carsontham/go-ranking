package repository

type User struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
	Score int64  `db:"score"`
}

type RankedUser struct {
	User
	Rank int64 `db:"rank"`
}
