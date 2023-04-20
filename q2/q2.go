package main

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	if len(text) == 0 {
		return 0, fmt.Errorf("texto vazio")
	}

	palavras := strings.Split(text, " ")
	soma := 0.
	for x := 0; x < len(palavras); x++ {
		soma += float64(len(palavras[x]))
	}

	media := soma / float64(len(palavras))

	fmt.Printf("%f", media)

	return 0, nil
}
