package main

import (
	"errors"
	"fmt"
)

func numerosComuns(slice1 []int, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Um dos slices está vazio")
	}

	numeros := make([]int, 0)
	numerosMap := make(map[int]bool)

	for _, num := range slice1 {
		numerosMap[num] = true
	}

	for _, num := range slice2 {
		if numerosMap[num] {
			numeros = append(numeros, num)
		}
	}

	return numeros, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}
	comuns, err := numerosComuns(slice1, slice2)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Números comuns:", comuns)
}
