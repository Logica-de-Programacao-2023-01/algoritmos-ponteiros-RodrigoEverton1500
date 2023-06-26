package main

import "fmt"

func SomaUltimosDigitos(ptr *int) {
	num := *ptr
	digito1 := num % 10
	digito2 := (num / 10) % 10

	*ptr = digito1 + digito2
}

func main() {
	numero := 1234

	fmt.Println("O número original é:", numero)
	SomaUltimosDigitos(&numero)
	fmt.Println("O novo valor é:", numero)
}
