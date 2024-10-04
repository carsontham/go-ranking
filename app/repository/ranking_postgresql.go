package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var _ RankingRepository = (*RankingRepo)(nil)

type GetDB func() *sql.DB

type RankingRepo struct {
	database *sql.DB
}

func NewRankingRepo(db *sql.DB) *RankingRepo {
	return &RankingRepo{
		database: db,
	}
}

func (repo RankingRepo) GetAllUser() ([]*User, error) {
	db := repo.database
	if db == nil {
		log.Println("database error")
	}
	var users []*User
	rows, err := db.Query("SELECT * FROM ranked_users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Score)
		if err != nil {
			fmt.Println("error in scan", err)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
