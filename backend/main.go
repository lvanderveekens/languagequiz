package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lvanderveekens/language-resources/api"
	"github.com/lvanderveekens/language-resources/postgres"
)

func main() {
	config, err := pgxpool.ParseConfig("postgres://postgres:postgres@localhost:5432/app?sslmode=disable")
	if err != nil {
		fmt.Println("Error parsing connection config:", err)
		return
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Println("Unable to create connection pool:", err)
		return
	}
	defer dbpool.Close()

	fmt.Println("Successfully connected to database!")

	exerciseStorage := postgres.NewExerciseStorage(dbpool)
	exerciseHandler := api.NewExerciseHandler(exerciseStorage)

	var handlers = api.NewHandlers(exerciseHandler)

	var server = api.NewServer(handlers)
	log.Fatal(server.Start(8080))
}
