package main

import (
	"errors"
	"fmt"
	"sort"
)

func ordenarSlice(sliceInteiros []int) ([]int, error) {
	if len(sliceInteiros) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	sortedSlice := make([]int, len(sliceInteiros))
	copy(sortedSlice, sliceInteiros)
	sort.Ints(sortedSlice)
	return sortedSlice, nil
}

func main() {
	slice := []int{5, 2, 9, 1, 7}
	ordenado, err := ordenarSlice(slice)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Ordenado:", ordenado)
}
