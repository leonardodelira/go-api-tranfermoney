package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func config() *pgxpool.Config {
	databaseURL := "postgres://postgres:postgres@localhost:5432/transfermoney" //TODO: create .env
	dbConfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}
	return dbConfig
}

func CreateConnectionPostgres() *pgxpool.Pool {
	connPool, err := pgxpool.NewWithConfig(context.Background(), config())
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}

	err = connPool.Ping(context.Background())
	if err != nil {
		fmt.Print(err)
		log.Fatal("Could not ping database")
	}

	fmt.Println("Connected to the database!!")

	return connPool
}
