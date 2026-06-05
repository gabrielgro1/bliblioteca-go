package repository

import (
	"biblioteca/domain"
	"database/sql"
	_ "github.com/lib/pq"
)

type BookRepoDB struct {
	db *sql.DB
}

func NewBookRepoDB(db *sql.DB) *BookRepoDB {
	return &BookRepoDB{db: db}
}

func (r *BookRepoDB) Save(book domain.Book){
	r.db.Exec(
		"INSERT INTO books (title, author, year, read) VALUES ($1, $2, $3, $4)",
		book.Title, book.Author, book.Year, book.Read,
	)
}

func (r *BookRepoDB) List() []domain.Book {
	rows, err := r.db.Query("SELECT id, title, author, year, read FROM books")
	if err != nil{
		return nil
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next(){
		var b domain.Book
		rows.Scan(&b.Id, &b.Title, &b.Author, &b.Year, &b.Read)
		books = append(books, b)

	}
	return books
}

	func (r *BookRepoDB) FindByTitle(title string) (domain.Book, bool) {
		row := r.db.QueryRow(
		"SELECT id, title, author, year, read FROM books WHERE LOWER(title)= LOWER($1)",
		title,
		)
	var b domain.Book
	err := row.Scan(&b.Id, &b.Title, &b.Author, &b.Year, &b.Read)
	if err != nil {
		return domain.Book{}, false
	}
	return b, true
	}

	func (r *BookRepoDB) MarkAsRead(id int) bool {
		result, err := r.db.Exec("UPDADE books SET read = true WHERE id = &1", id)
		if err != nil {
			return false
		}
		rows, _ := result.RowsAffected()
		return rows > 0 	
	}

	func (r *BookRepoDB) Delete(id int) bool {
		result, err := r.db.Exec("DELETE FROM books WHERE id = $1", id)
		if err != nil {
			return false
		}
		rows, _ := result.RowsAffected()
		return rows > 0
	}