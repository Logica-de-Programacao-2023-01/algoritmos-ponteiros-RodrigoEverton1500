package main

import (
	"fmt"
)

func SomaNaturais(ptr *int, n int) {
	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	*ptr = s
}

func main() {
	var num int
	n := 5

	SomaNaturais(&num, n)

	fmt.Println("O valor do inteiro é:", num)
}
