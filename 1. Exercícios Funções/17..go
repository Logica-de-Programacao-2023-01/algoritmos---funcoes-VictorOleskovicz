package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func ordenarStrings(sliceStrings []string) (string, error) {
	if len(sliceStrings) == 0 {
		return "", errors.New("O slice está vazio")
	}

	sort.Strings(sliceStrings)
	result := strings.Join(sliceStrings, "")
	return result, nil
}

func main() {
	slice := []string{"banana", "abacaxi", "laranja", "
