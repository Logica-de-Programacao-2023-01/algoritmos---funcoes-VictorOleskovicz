package main

import (
	"fmt"
	"sort"
)

func encontrarSegundoMaior(sliceInteiros []int) (int, error) {
	if len(sliceInteiros) < 2 {
		return 0, fmt.Errorf("O slice deve conter pelo menos dois elementos")
	}

	sort.Ints(sliceInteiros)
	return sliceInteiros[len(sliceInteiros)-2], nil
}

func main() {
	slice := []int{5, 2, 9, 1, 7}
	segundoMaior, err := encontrarSegundoMaior(slice)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Segundo Maior:", segundoMaior)
}
