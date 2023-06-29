package main

import (
	"errors"
	"fmt"
)

func aplicarFuncao(slice []int, funcao func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("O slice está vazio")
	}

	soma := 0
	for _, num := range slice {
		resultado := funcao(num)
		soma += resultado
	}

	return soma, nil
}

// Função exemplo que pode ser passada como parâmetro
func dobrarNumero(numero int) int {
	return numero * 2
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	soma, err := aplicarFuncao(slice, dobrarNumero)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Soma dos resultados:", soma)
}
