package main

import (
	"fmt"
)

func preencherPrimos(s *[]int, n int) {
	count := 0
	num := 2

	for count < n {
		if isPrimo(num) {
			*s = append(*s, num)
			count++
		}
		num++
	}
}

func isPrimo(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var primos []int
	tamanho := 10

	preencherPrimos(&primos, tamanho)

	fmt.Println("Números primos:", primos)
}
