package data

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type Data struct {
	db    *sql.DB
	cache *redis.Client
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

	url, ok = os.LookupEnv("REDIS_URL")
	if !ok {
		log.Fatal("redis url is not specified")
	}

	password, ok := os.LookupEnv("REDIS_PASSWORD")
	if !ok {
		log.Fatal("redis password is not specified")
	}

	cache := redis.NewClient(
		&redis.Options{
			Addr:     url,
			Password: password,
		},
	)

	return &Data{
		db:    db,
		cache: cache,
	}
}
