package main

import (
	"fmt"
)

func modificarValor(ponteiro *int) {
	*ponteiro = 42
}

func main() {
	var valor int
	ponteiro := &valor
	modificarValor(ponteiro)
	fmt.Println("Novo valor:", valor)
	ponteiro = nil
}
