package service

import (
	"biblioteca/domain"
	"errors"
	"strings"
)

type BookRepository interface {
	Save(book domain.Book)
	List() []domain.Book
	FindByTitle(title string) (domain.Book, bool)
	MarkAsRead(id int) bool
	Delete(id int) bool
}

type BookService struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) CreateBook(title string, autor string, year int) error {
	title = strings.TrimSpace(title)
	autor = strings.TrimSpace(autor)
	
	if title == "" {
		return errors.New("Titulo é obrigatorio")
	}

	if autor == "" {
		return errors.New("Autor obrigatorio")
	}

	if year <= 0 {
		return errors.New("Ano invalido")
	}

	_, encontrado := s.repo.FindByTitle(title)
	if encontrado {
		return errors.New("Livro ja cadastrado")
	}
	book := domain.Book{
		Title:  title,
		Author: autor,
		Year:   year,
		Read:   false,
	}
	s.repo.Save(book)

	return nil
}

func (s *BookService) ListBooks() []domain.Book {
	return s.repo.List()
}

func (s *BookService) FindByTitle(title string) (domain.Book, bool) {
	title = strings.TrimSpace(title)
	if title == ""{
		return domain.Book{}, false
	}
	
	return s.repo.FindByTitle(title)

}

func (s *BookService) MarkAsRead(id int) error {
	if id <= 0 {
		return errors.New("ID invalido")
	}

	marcado := s.repo.MarkAsRead(id)
	if !marcado {
		return errors.New("Livro não encontrado")
	}
	return nil

}

func (s *BookService) Delete(id int) error {
	if id <= 0 {
		return errors.New("ID invalido")

	}
	deletado := s.repo.Delete(id)
	if !deletado {
		return errors.New("Livro não encontrado")
	}
	return nil
}
