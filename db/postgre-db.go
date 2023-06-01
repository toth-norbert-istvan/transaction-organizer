package transaction_organizer_db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Env struct {
	db *sql.DB
}

func Connect() {
	db, err := sql.Open("postgres", "postgresql://postgres:password@localhost:5432/transaction-organizer?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	_ = &Env{db: db}
}
