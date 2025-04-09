package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"github.com/nathanSeixeiro/gateway-payments/internal/repository"
	"github.com/nathanSeixeiro/gateway-payments/internal/service"
	"github.com/nathanSeixeiro/gateway-payments/internal/web/server"
)

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// string connection to postgres
	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		GetEnv("DB_HOST", "db"),
		GetEnv("DB_PORT", "5432"),
		GetEnv("DB_USER", "postgres"),
		GetEnv("DB_PASSWORD", "postgres"),
		GetEnv("DB_NAME", "gateway_payments"),
		GetEnv("DB_SSLMODE", "disable"),
	)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error opening database")
	}
	defer db.Close()

	// dependency injection
	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	// create server
	server := server.NewServer(accountService, os.Getenv("HTTP_PORT"))
	server.ConfigureRoutes()
	fmt.Println("Server is running on port", os.Getenv("HTTP_PORT"))
	server.Start()
}
