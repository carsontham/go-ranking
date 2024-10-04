package bootstrap

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetDB() *sql.DB {
	const driverName = "postgres"
	const dataSourceName = "postgresql://root:secret@localhost:5432/ranking-db?sslmode=disable"

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	return db
}
