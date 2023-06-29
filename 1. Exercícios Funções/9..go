package main

import (
	"errors"
	"fmt"
	"strings"
)

func obterPalavras(str string) ([]string, error) {
	if len(str) == 0 {
		return nil, errors.New("A string está vazia")
	}

	palavras := strings.Split(str, " ")
	return palavras, nil
}

func main() {
	str := "Olá, mundo! Bem-vindo ao Golang."
	palavras, err := obterPalavras(str)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Palavras:", palavras)
}
