package data

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Data struct {
	db *sql.DB
}

func New() *Data {
	url, ok := os.LookupEnv("POSTGRES_URL")
	if !ok {
		log.Fatal("postgres uri is not specified")
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("can't connect to postgres: ", err)
	}

	return &Data{
		db: db,
	}
}
