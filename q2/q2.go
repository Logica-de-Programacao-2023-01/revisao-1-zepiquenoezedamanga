package q2

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	speciais:= []string{"!","?","@","#","$",".",",","1","2","3","4","5","6","7","8","9"}
	for _, caracter:= range speciais {
		text = strings.ReplaceAll(text,caracter,"")
	}
	palavras := strings.Fields(text)
	if len(palavras) == 0 {
		return 0, fmt.Errorf("texto vazio")
		
	}
	var totalpalavra int
	for _, palavra:= range palavras {
		totalpalavra += len(palavra)
	}
	return float64(totalpalavra) / float64(len(palavras)),nil
}
