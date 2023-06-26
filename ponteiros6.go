package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func AlterarTituloLivro(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	livro1 := Livro{Titulo: "Livro 1", Autor: "Autor 1"}
	livro2 := Livro{Titulo: "Livro 2", Autor: "Anônimo"}

	fmt.Println("Livro 1 antes:", livro1)
	fmt.Println("Livro 2 antes:", livro2)

	AlterarTituloLivro(&livro1)
	AlterarTituloLivro(&livro2)

	fmt.Println("Livro 1 depois:", livro1)
	fmt.Println("Livro 2 depois:", livro2)
}
