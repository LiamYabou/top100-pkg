package db

import (
	"os"
	"github.com/jackc/pgx/v4/pgxpool"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

var testDBURL  = os.Getenv("TEST_DB_DSN")

func Open(dbURL string) (db *pgxpool.Pool, err error) {
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}
	db, err = pgxpool.ConnectConfig(context.Background(), config)
	return db, err
}

func OpenTest() (db *pgxpool.Pool, err error) {
	db, err = pgxpool.Connect(context.Background(), testDBURL)
	return db, err
}

func OpenPQ(dbURL string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", dbURL)
	return
}

func OpenPQtest() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", testDBURL)
	return db, err
}
