package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"languagequiz/api"
	"languagequiz/postgres"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed loading .env file: ", err)
	}

	zoneName, _ := time.Now().Zone()
	fmt.Println("Configured time zone: ", zoneName)

	connString := os.Getenv("POSTGRES_CONN_STRING")
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatal("Error parsing connection config: ", err)
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Unable to create connection pool: ", err)
	}
	defer dbpool.Close()

	err = dbpool.Ping(context.Background())
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	fmt.Println("Database connection successful!")

	m, err := migrate.New("file://migrations", connString)
	if err != nil {
		log.Fatal("Failed to create Migrate instance: ", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Failed to migrate the database: ", err)
	}

	quizStorage := postgres.NewQuizStorage(dbpool)
	quizHandler := api.NewQuizHandler(quizStorage)

	feedbackHandler := api.NewFeedbackHandler(os.Getenv("DISCORD_BOT_TOKEN"), os.Getenv("DISCORD_FEEDBACK_CHANNEL_ID"))

	var handlers = api.NewHandlers(quizHandler, feedbackHandler)

	var server = api.NewServer(handlers)
	log.Fatal(server.Start(8888))
}
