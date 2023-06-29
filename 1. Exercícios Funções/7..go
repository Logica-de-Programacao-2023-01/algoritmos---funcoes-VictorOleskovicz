package main

import (
	"errors"
	"fmt"
)

func aplicarFuncaoEmSlice(sliceInteiros []int, funcao func(int) int) ([]int, error) {
	if len(sliceInteiros) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	result := make([]int, len(sliceInteiros))
	for i, num := range sliceInteiros {
		result[i] = funcao(num)
	}

	return result, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	dobro := func(x int) int { return x * 2 }
	resultado, err := aplicarFuncaoEmSlice(slice, dobro)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Resultado:", resultado)
}
