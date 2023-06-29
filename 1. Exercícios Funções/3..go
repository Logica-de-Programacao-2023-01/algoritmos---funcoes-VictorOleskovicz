package main

import "fmt"

func concatenarStrings(sliceStrings []string) string {
	result := ""

	for _, str := range sliceStrings {
		result += str
	}

	return result
}

func main() {
	slice := []string{"Hello", " ", "World!"}
	concatenado := concatenarStrings(slice)
	fmt.Println("Concatenado:", concatenado)
}
