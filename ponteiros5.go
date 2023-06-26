package main

import (
	"fmt"
	"math"
)

func AtualizarMedia(ptr *float64) {
	valor := *ptr
	media := (valor + math.Pi) / 2
	*ptr = media
}

func main() {
	numero := 3.14

	fmt.Println("O número original é:", numero)
	AtualizarMedia(&numero)
	fmt.Println("O novo valor é:", numero)
}
