package main

import (
	"errors"
	"fmt"
	"strings"
)

func substituirVogaisPorAsterisco(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("A string está vazia")
	}

	vogais := "aeiouAEIOU"
	result := ""
	for _, char := range str {
		if strings.ContainsRune(vogais, char) {
			result += "*"
		} else {
			result += string(char)
		}
	}

	return result, nil
}

func main() {
	str := "Hello, World!"
	novaStr, err := substituirVogaisPorAsterisco(str)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Nova string:", novaStr)
}
