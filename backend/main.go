package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lvanderveekens/language-resources/api"
	"github.com/lvanderveekens/language-resources/postgres"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	zoneName, _ := time.Now().Zone()
	fmt.Printf("Configured time zone: %s", zoneName)

	connString := "postgres://postgres:postgres@localhost:15432/app?sslmode=disable"
	config, err := pgxpool.ParseConfig(connString)
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

	m, err := migrate.New("file://migrations", connString)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	exerciseStorage := postgres.NewExerciseStorage(dbpool)
	exerciseHandler := api.NewExerciseHandler(exerciseStorage)

	var handlers = api.NewHandlers(exerciseHandler)

	var server = api.NewServer(handlers)
	log.Fatal(server.Start(8888))
}
