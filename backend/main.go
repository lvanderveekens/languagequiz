package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/lvanderveekens/language-resources/api"
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

	var version string
	err = conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	fmt.Println("PostgreSQL version:", version)

	var server = api.NewServer()
	log.Fatal(server.Start(8080))
}
