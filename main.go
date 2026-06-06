package main

import (
	"database/sql"
	"fmt"
	"os"

	"biblioteca/handler"
	"biblioteca/repository"
	"biblioteca/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar .env")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Erro ao conectar ao banco", err)
		os.Exit(1)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Banco não respondeu", err)
		os.Exit(1)
	}

	repo := repository.NewBookRepoDB(db)
	BookService := service.NewBookService(repo)
	menuHandler := handler.NewHandler(BookService)

	menuHandler.Run()

}
