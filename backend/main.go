package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
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

	// Execute query
	var version string
	err = conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	fmt.Println("PostgreSQL version:", version)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
