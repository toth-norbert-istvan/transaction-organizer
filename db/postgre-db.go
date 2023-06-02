package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var postgreSqlDBInstance *PostgreSqlDB
var dbConnection *sql.DB
var once sync.Once

type SqlDB interface {
	Connect()
	GetDb() *sql.DB
}

type PostgreSqlDB struct{}

func (psd PostgreSqlDB) Connect() {
	once.Do(func() {
		postgreSqlDBInstance = &PostgreSqlDB{}
		db, err := sql.Open("postgres", "postgresql://postgres:password@localhost:5432/transaction-organizer?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		dbConnection = db
		log.Println("SQL database connection established")
	})
}

func (psd PostgreSqlDB) GetDb() *sql.DB {
	if dbConnection == nil {
		once.Do(func() {
			psd.Connect()
		})
	}
	return dbConnection
}
