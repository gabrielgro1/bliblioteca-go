package main

import (
	"biblioteca/handler"
	"biblioteca/repository"
	"biblioteca/service"
)

func main() {
	repo := repository.NewBookRepo()
	bookService := service.NewBookService(repo)
	menuHandler := handler.NewHandler(bookService)

	menuHandler.Run()
}
