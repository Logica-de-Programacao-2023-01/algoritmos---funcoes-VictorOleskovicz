package main

import "fmt"

func encontrarPosicao(sliceInteiros []int, valor int) int {
	for i, num := range sliceInteiros {
		if num == valor {
			return i
		}
	}

	return -1
}

func main() {
	slice := []int{2, 4, 6, 8, 10}
	valor := 6
	posicao := encontrarPosicao(slice, valor)
	fmt.Println("Posição do primeiro elemento igual a", valor, ":", posicao)
}
