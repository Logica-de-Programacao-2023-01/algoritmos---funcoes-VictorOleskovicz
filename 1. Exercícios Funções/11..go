package main

import (
	"errors"
	"fmt"
	"math"
)

func verificarNumeroPrimo(num int) (bool, error) {
	if num < 0 {
		return false, errors.New("O número é negativo")
	}

	if num < 2 {
		return false, nil
	}

	limit := int(math.Sqrt(float64(num))) + 1
	for i := 2; i < limit; i++ {
		if num%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	num := 17
	primo, err := verificarNumeroPrimo(num)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("É primo:", primo)
}
