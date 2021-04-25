package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	PostgresDB *sql.DB
)

func ConnectPostgresDB() bool {
	var err error
	PostgresDB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println(PostgresDB.Stats())
	return true
}

func InitDB() bool {
	if _, err := PostgresDB.Exec("DROP TABLE IF EXISTS person;"); err != nil {
		log.Fatal(err)
		return false
	}
	if _, err := PostgresDB.Exec("CREATE TABLE IF NOT EXISTS person (id VARCHAR(20) PRIMARY KEY, name VARCHAR(20), password VARCHAR(50), email VARCHAR(60));"); err != nil {
		log.Fatal(err)
		return false
	}

	if _, err := PostgresDB.Exec("DROP TABLE IF EXISTS paymentrecord;"); err != nil {
		log.Fatal(err)
		return false
	}
	if _, err := PostgresDB.Exec("CREATE TABLE IF NOT EXISTS paymentrecord (id serial PRIMARY KEY, class VARCHAR(20), payment INTEGER);"); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
