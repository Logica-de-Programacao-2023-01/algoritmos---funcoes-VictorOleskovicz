package main

import "fmt"

func contarVogais(str string) int {
	count := 0
	vogais := "aeiouAEIOU"

	for _, char := range str {
		if vogais.Contains(string(char)) {
			count++
		}
	}

	return count
}

func main() {
	str := "Hello, World!"
	qtdVogais := contarVogais(str)
	fmt.Println("Quantidade de vogais:", qtdVogais)
}
