package main

import "fmt"

func InverterString(ptr *string) {
	str := *ptr
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	*ptr = string(runes)
}

func main() {
	texto := "Oi"

	fmt.Println("O texto original é:", texto)
	InverterString(&texto)
	fmt.Println("O texto invertido é:", texto)
}
