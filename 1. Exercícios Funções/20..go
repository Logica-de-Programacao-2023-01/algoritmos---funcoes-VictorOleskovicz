package main

import (
	"errors"
	"fmt"
)

func stringsComMaisDe5Caracteres(sliceStrings []string) ([]string, error) {
	if len(sliceStrings) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	result := make([]string, 0)
	for _, str := range sliceStrings {
		if len(str) > 5 {
			result = append(result, str)
		}
	}

	return result, nil
}

func main() {
	slice := []string{"Golang", "OpenAI", "Hello", "World"}
	stringsFiltradas, err := stringsComMaisDe5Caracteres(slice)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Strings filtradas:", stringsFiltradas)
}
