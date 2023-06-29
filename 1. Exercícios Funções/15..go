package main

import (
	"errors"
	"fmt"
)

func verificarPresenca(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("O slice está vazio")
	}

	for _, num := range slice {
		if num == numero {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	num := 3
	slice := []int{1, 2, 3, 4, 5}
	presente, err := verificarPresenca(num, slice)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Está presente:", presente)
}
