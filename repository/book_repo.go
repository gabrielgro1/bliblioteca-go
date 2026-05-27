package repository

import (
	"biblioteca/domain"
	"strings"
)

type BookRepo struct {
	books  []domain.Book
	nextId int
}

func NewBookRepo() *BookRepo {
	return &BookRepo{
		books:  []domain.Book{},
		nextId: 1,
	}
}
func (r *BookRepo) Save(book domain.Book) {
	book.Id = r.nextId
	r.nextId++
	r.books = append(r.books, book)
}

func (r *BookRepo) List() []domain.Book {
	return r.books
}

func (r *BookRepo) FindByTitle(title string) (domain.Book, bool) {
	for _, book := range r.books {
		if strings.EqualFold(book.Title, title) {
			return book, true
		}
	}
	return domain.Book{}, false
}
func (r *BookRepo) MarkAsRead(id int) bool {
	for i := range r.books {
		if r.books[i].Id == id {
			r.books[i].Read = true
			return true
		}
	}
	return false
}
func (r *BookRepo) Delete(id int) bool {
	for i := range r.books {
		if r.books[i].Id == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return true
		}
	}

	return false

}
