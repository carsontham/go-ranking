package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"go-ranking/app/handlers/rest"
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

func (repo RankingRepo) GetAllUser(sortScoreInDesc bool, minScore int) ([]*User, error) {
	db := repo.database
	if db == nil {
		log.Println("database error")
	}
	baseQuery := "SELECT * FROM ranked_users"

	var args []interface{}

	if minScore > 0 {
		baseQuery += " WHERE score > $1"
		args = append(args, minScore)
	}

	// Add sorting by score if requested
	if sortScoreInDesc {
		baseQuery += " ORDER BY score DESC"
	}

	log.Println(baseQuery)
	var users []*User
	rows, err := db.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

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

func (repo RankingRepo) CreateNewUser(user *User) error {
	db := repo.database
	if db == nil {
		log.Println("database error")
	}

	query := `INSERT INTO ranked_users (name, email, score) VALUES ($1, $2, $3)`
	result, err := db.Exec(query, &user.Name, &user.Email, &user.Score)

	if err != nil {
		return err
	}
	res, _ := result.RowsAffected()
	log.Printf("inserted %d row into user table", res)
	return nil
}

func (repo RankingRepo) CheckUniqueEmail(email string) (bool, error) {
	db := repo.database
	if db == nil {
		log.Println("Database connection error")
		return false, errors.New("database connection error")
	}

	// Query the database to check if the email exists
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM ranked_users WHERE email = $1)"
	row := db.QueryRow(query, email)
	err := row.Scan(&exists)
	if err != nil {
		log.Println("Error checking email uniqueness:", err)
		return false, err
	}

	// Return true if the email does not exist
	return !exists, nil
}

func (repo RankingRepo) GetUserByID(id int64) (*User, error) {
	db := repo.database
	if db == nil {
		log.Println("database error")
	}

	query := "SELECT * FROM ranked_users WHERE id = $1"
	row := db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Score)
	if err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, rest.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (repo RankingRepo) UpdateUserByID(user *User) error {
	db := repo.database
	if db == nil {
		log.Println("database error")
	}
	updateQuery := "UPDATE ranked_users SET name = $1, email = $2, score = $3 WHERE id = $4"
	result, err := db.Exec(updateQuery, user.Name, user.Email, user.Score, user.ID)
	if err != nil {
		log.Println("Error updating user:", err)
		return err
	}
	if err != nil {
		return err
	}
	res, _ := result.RowsAffected()
	log.Printf("inserted %d row into user table", res)
	return nil
}

func (repo RankingRepo) DeleteUserByID(id int64) error {
	db := repo.database
	if db == nil {
		log.Println("database error")
		return errors.New("database not initialized")
	}

	deleteQuery := "DELETE FROM ranked_users WHERE id = $1"
	result, err := db.Exec(deleteQuery, id)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return err
	}
	log.Printf("deleted %d row(s) from user table with ID %d", rowsAffected, id)
	return nil
}

func (repo RankingRepo) GetAllUserRanking() ([]RankedUser, error) {
	db := repo.database
	if db == nil {
		log.Println("database error")
		return nil, errors.New("database not initialized")
	}

	query := ` 	SELECT id, name, email, score,
				RANK() OVER (ORDER BY score DESC) AS rank
				FROM ranked_users
				ORDER BY score DESC;
			`

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var rankedUsers []RankedUser

	for rows.Next() {
		var rankedUser RankedUser

		err := rows.Scan(&rankedUser.ID, &rankedUser.Name, &rankedUser.Email, &rankedUser.Score, &rankedUser.Rank)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		rankedUsers = append(rankedUsers, rankedUser)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error after iterating rows:", err)
		return nil, err
	}

	return rankedUsers, nil
}

func (repo RankingRepo) GetUserRankByID(userID int64) (*RankedUser, error) {
	db := repo.database
	if db == nil {
		log.Println("database error")
		return nil, errors.New("database not initialized")
	}

	query := `
        WITH RankedUsers AS (
        	SELECT id, name, email, score,
        	RANK() OVER (ORDER BY score DESC) AS rank
            FROM ranked_users
        )
        SELECT id, name, email, score, rank FROM RankedUsers WHERE id = $1;
    `

	var rankedUser RankedUser
	err := db.QueryRow(query, userID).Scan(&rankedUser.ID, &rankedUser.Name, &rankedUser.Email, &rankedUser.Score, &rankedUser.Rank)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest.ErrNotFound
		}
		log.Println("Error getting user rank:", err)
		return nil, err
	}
	return &rankedUser, nil
}
