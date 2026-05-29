package handler

import (
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
	var opcao string

	for {
		fmt.Println("=== Bliblioteca de livros ===")
		fmt.Println("1. Cadastrar livro")
		fmt.Println("2. Listar livros")
		fmt.Println("3. Buscar livros")
		fmt.Println("4. Marcar como lido")
		fmt.Println("5. Deletar livro")
		fmt.Println("0. Sair")
		fmt.Print("Digite uma opção: ")
		opcao, _ = h.reader.ReadString('\n')
		opcao = strings.TrimSpace(opcao)

		switch opcao {

		case "1":
			h.handleCreateBook()

		case "2":
			h.ListBooks()

			
		case "3":
			h.handleFindBook()

		case "4":
			h.handleMarkAsRead()

		case "5":
			h.handlerDeleteBook()


		case "0":
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println()
			fmt.Println("Opção invalida!")
			fmt.Println("Presione Enter para continuar...")
			h.reader.ReadString('\n')
		}

	}
}

func (h *Handler) handleCreateBook() {
	fmt.Print("Titulo: ")
	titulo, _ := h.reader.ReadString('\n')
	titulo = strings.TrimSpace(titulo)

	fmt.Print("Autor: ")
	autor, _ := h.reader.ReadString('\n')
	autor = strings.TrimSpace(autor)

	fmt.Print("Ano: ")
	anoTexto, _ := h.reader.ReadString('\n')
	anoTexto = strings.TrimSpace(anoTexto)

	ano, err := strconv.Atoi(anoTexto)
	if err != nil {
		fmt.Println("Ano invalido!")
		return
	}
}
	err = h.service.CreateBook(titulo, autor, ano)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Livro cadastrado!")

func (h *Handler) handleListBooks() 
	livros := h.service.ListBooks()
	if len(livros) == 0 {
		fmt.Println("Nenhum livro cadastrado.")
		return
			}

	for _, livro := range livros {
				fmt.Println("-------------------")
				fmt.Println("ID", livro.Id)
				fmt.Println("Titulo", livro.Title)
				fmt.Println("Autor", livro.Author)
				fmt.Println("Year", livro.Year)
				fmt.Println("Lido", livro.Read)

	}

func (h.*Handler) MarkFindBook() {
	fmt.Println("Digite o titulo: ")
			titulo, _ := h.reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)

			livro, encontrado := h.service.FindByTitle(titulo)

			if !encontrado {
				fmt.Println("Livro não encontrado: ")
				continue
			}
			fmt.Println("Livro encontrado")
			fmt.Println("Id", livro.Id)
			fmt.Println("Titulo", livro.Title)
			fmt.Println("Autor", livro.Author)
			fmt.Println("Year", livro.Year)
			fmt.Println("Lido", livro.Read)
}	

func (h *Handler) MarkAsRead() {
	fmt.Println("Digite o ID do livro: ")
			idTexto, _ := h.reader.ReadString('\n')
			idTexto = strings.TrimSpace(idTexto)

			id, err := strconv.Atoi(idTexto)
			if err != nil {
				fmt.Println("Id invalido")
				return

			}

			err = h.service.MarkAsRead(id)

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Livro marcado como lido!")
		}
func (h *Handler) handlerDeleteBook() {
	fmt.Println("Digite o ID pra poder deletar")
			idtexto, _ := h.reader.ReadString('\n')
			idtexto = strings.TrimSpace(idtexto)

			id, err := strconv.Atoi(idtexto)
			if err != nil {
				fmt.Println("ID invalido")
				continue
			}

			err = h.service.Delete(id)

			if err != nil {
				fmt.Println(err)
				return
									
		
			fmt.Println("Livro deletado!")

}

}
