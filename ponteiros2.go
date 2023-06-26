package main

import "fmt"

func VerificarParImpar(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}
}

func main() {
	num := 7

	fmt.Println("O valor original do inteiro é:", num)
	VerificarParImpar(&num)
	fmt.Println("O novo valor do inteiro é:", num)
}
