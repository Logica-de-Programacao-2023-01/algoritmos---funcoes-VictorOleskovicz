package main

import (
	"errors"
	"fmt"
)

func obterPares(sliceInteiros []int) ([]int, error) {
	if len(sliceInteiros) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	result := []int{}
	for _, num := range sliceInteiros {
		if num%2 == 0 {
			result = append(result, num)
		}
	}

	return result, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	pares, err := obterPares(slice)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Pares:", pares)
}
