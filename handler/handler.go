package handler

import (
	"biblioteca/domain"
	"biblioteca/service"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Handler struct {
	service *service.BookService
	reader  *bufio.Reader
}

func NewHandler(service *service.BookService) *Handler {
	return &Handler{
		service: service,
		reader:  bufio.NewReader(os.Stdin),
	}
}

func (h *Handler) Run() {
	for {
		h.showMenu()

		opcao := h.readLine("Digite uma opção: ")

		switch opcao {

		case "1":
			h.handleCreateBook()
			h.pause()

		case "2":
			h.handleListBooks()
			h.pause()
		case "3":
			h.handleFindBook()
			h.pause()

		case "4":
			h.handleMarkAsRead()
			h.pause()

		case "5":
			h.handleDeleteBook()
			h.pause()

		case "0":
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println()
			fmt.Println("Opção inválida!")
			h.pause()
		}

	}
}

func (h *Handler) handleCreateBook() {
	titulo := h.readLine("Titulo: ")
	autor := h.readLine("Autor: ")

	ano, err := h.readInt("Ano: ")
	if err != nil {
		fmt.Println("Ano invalido!")
		return
	}

	err = h.service.CreateBook(titulo, autor, ano)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Livro cadastrado!")
}

func (h *Handler) handleListBooks() {
	livros := h.service.ListBooks()
	if len(livros) == 0 {
		fmt.Println("Nenhum livro cadastrado.")
		return
	}

	for _, livro := range livros {
		h.printBook(livro)
	}
}

func (h *Handler) handleFindBook() {
	titulo := h.readLine("Digite o titulo: ")

	livro, encontrado := h.service.FindByTitle(titulo)

	if !encontrado {
		fmt.Println("Livro não encontrado: ")
		return
	}

	h.printBook(livro)
}
func (h *Handler) handleMarkAsRead() {
	id, err := h.readInt("Digite o ID do livro: ")
	if err != nil {
		fmt.Println("ID invalido")
		return
	}

	err = h.service.MarkAsRead(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Livro marcado como lido!")
}
func (h *Handler) handleDeleteBook() {
	id, err := h.readInt("Digite o ID do livro pra poder deletar: ")
	if err != nil {
		fmt.Println("ID invalido")
		return
	}

	err = h.service.Delete(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Livro deletado!")

}

func (h *Handler) readLine(message string) string {
	fmt.Print(message)

	text, _ := h.reader.ReadString('\n')

	return strings.TrimSpace(text)
}

func (h *Handler) printBook(livro domain.Book) {
	fmt.Println("----------------------")
	fmt.Println("ID:", livro.Id)
	fmt.Println("Título:", livro.Title)
	fmt.Println("Autor:", livro.Author)
	fmt.Println("Ano:", livro.Year)
	fmt.Println("Lido:", livro.Read)
}

func (h *Handler) readInt(message string) (int, error) {
	texto := h.readLine(message)

	numero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, err
	}

	return numero, nil
}

func (h *Handler) pause() {
	h.readLine("Pressione Enter para continuar ....")

}

func (h *Handler) showMenu() {
	fmt.Println("=== Biblioteca de livros ===")
	fmt.Println("1. Cadastrar livro")
	fmt.Println("2. Listar livros")
	fmt.Println("3. Buscar livros")
	fmt.Println("4. Marcar como lido")
	fmt.Println("5. Deletar livro")
	fmt.Println("0. Sair")

}
