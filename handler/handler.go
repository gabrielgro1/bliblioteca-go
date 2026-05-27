package handler

import (
	"biblioteca/service"
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Handler struct {
	service *service.BookService
	reader *bufio.Reader
}
 
func NewHandler (service *service.BookService) *Handler {
	return &Handler{
		service: service,
		reader: bufio.NewReader(os.Stdin),
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
				continue
			}

			err = h.service.CreateBook(titulo, autor, ano)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Livro cadastrado!")

		case "2":
			livros := h.service.ListBooks()

			if len(livros) == 0 {

				fmt.Println("Nenhum livro cadastrado.")
				continue
			}
			for _, livro := range livros {
				fmt.Println("-------------------")
				fmt.Println("ID", livro.Id)
				fmt.Println("Titulo", livro.Title)
				fmt.Println("Autor", livro.Author)
				fmt.Println("Year", livro.Year)
				fmt.Println("Lido", livro.Read)

			}
		case "3":
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

		case "4":
			fmt.Println("Digite o ID do livro: ")
			idTexto, _ := h.reader.ReadString('\n')
			idTexto = strings.TrimSpace(idTexto)

			id, err := strconv.Atoi(idTexto)
			if err != nil {
				fmt.Print("Id invalido")
				continue

			}

			err = h.service.MarkAsRead(id)

			if err != nil {
				fmt.Print(err)
				continue
			}

			fmt.Println("Livro marcado como lido!")
		case "5":
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
				continue

			}

			fmt.Println("Livro deletado!")

		case "0":
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println()
			fmt.Println("Opção invalida!")
			fmt.Print("Presione Enter para continuar...")
			h.reader.ReadString('\n')
		}

	}
}


