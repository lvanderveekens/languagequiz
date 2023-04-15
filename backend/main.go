package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/lvanderveekens/language-resources/api"
	"github.com/lvanderveekens/language-resources/postgres"
)

func main() {
	config, err := pgx.ParseConfig("postgres://postgres:postgres@localhost:5432/app?sslmode=disable")
	if err != nil {
		fmt.Println("Error parsing connection config:", err)
		return
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer conn.Close(context.Background())

	fmt.Println("Successfully connected to database!")

	exerciseStorage := postgres.NewExerciseStorage(conn)
	exerciseHandler := api.NewExerciseHandler(exerciseStorage)

	var handlers = api.NewHandlers(exerciseHandler)

	var server = api.NewServer(handlers)
	log.Fatal(server.Start(8080))
}
