package psqldb

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDB() {
	connString := "postgres://admin:password123@postgres:5432/userdb"

	var err error

	Pool, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	err = Pool.Ping(context.Background())
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	log.Println("Connected to PostgreSQL")

	CreateTables()
}

func CreateTables() {

	query := `
	CREATE TABLE IF NOT EXISTS documents (
		id BIGSERIAL PRIMARY KEY,
		owner_type VARCHAR(20) NOT NULL,
		owner_id VARCHAR(100) NOT NULL,

		filename TEXT NOT NULL,
		storage_path TEXT NOT NULL,

		file_size BIGINT NOT NULL,
		mime_type VARCHAR(100) NOT NULL,

		status VARCHAR(20) NOT NULL,

		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	);
	`

	_, err := Pool.Exec(context.Background(), query)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	log.Println("documents table ready")
}

func CloseDB() {
	if Pool != nil {
		Pool.Close()
	}
}
