package main

import (
	"errors"
	"fmt"
	"math"
)

func numerosPrimosAte(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("O número é negativo")
	}

	primos := make([]int, 0)

	for i := 2; i <= numero; i++ {
		if isPrimo(i) {
			primos = append(primos, i)
		}
	}

	return primos, nil
}

func isPrimo(numero int) bool {
	if numero < 2 {
		return false
	}

	limite := int(math.Sqrt(float64(numero)))
	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	num := 10
	primos, err := numerosPrimosAte(num)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Números primos:", primos)
}
