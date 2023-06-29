package main

import (
	"errors"
	"fmt"
	"strings"
)

func concatenarComVirgulas(sliceStrings []string) (string, error) {
	if len(sliceStrings) == 0 {
		return "", errors.New("O slice está vazio")
	}

	result := strings.Join(sliceStrings, ",")
	return result, nil
}

func main() {
	slice := []string{"Hello", "World", "Golang"}
	concatenado, err := concatenarComVirgulas(slice)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Concatenado:", concatenado)
}
