package main

import (
	"fmt"
	"go-repeat/internal/file"
	"go-repeat/internal/text"
)

func main() {
	content, err := file.ReadFile("internal/testdata/input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	toAllUpper := text.ToUpperCase(content)
	fmt.Println(toAllUpper)
	fmt.Println(text.TolowerCase(toAllUpper))

	fmt.Println(content)
	fmt.Println(text.CountWords(content))
	fmt.Println(text.CapitalizeWords(content))

	err = file.WriteFile("internal/testdata/output.txt", "Hello World")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
}
