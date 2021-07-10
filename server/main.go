package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type config struct {
	db struct {
		dsn string
	}
}

func main() {
	var cfg config

	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://musichan:pa55word@localhost/musichan", "PostgreSQL data source name")
	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Printf("database connection pool established")
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
