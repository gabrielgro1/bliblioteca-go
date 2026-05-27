package service

import (
	"biblioteca/domain"
	"biblioteca/repository"
	"errors"
)

type BookService struct {
	repo *repository.BookRepo
}

func NewBookService(repo *repository.BookRepo) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) CreateBook(title string, autor string, year int) error {
	if title == "" {
		return errors.New("Titulo é obrigatorio")
	}

	if autor == "" {
		return errors.New("autor obrigatorio")
	}

	if year <= 0 {
		return errors.New("Ano invalido")
	}

	_, encontrado := s.repo.FindByTitle(title)
	if encontrado {
		return errors.New("livro ja cadastrado")
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
	return s.repo.FindByTitle(title)

}

func (s *BookService) MarkAsRead(id int) error {
	if id <= 0 {
		return errors.New("Id invalido")
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
