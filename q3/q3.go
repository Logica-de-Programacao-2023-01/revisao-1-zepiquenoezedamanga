package q3

import (
	"fmt"
	"sort"
)

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, fmt.Errorf("lista vazia")
	}
	var maior int
	var menor int
	sort.Ints(numbers)
	menor = (numbers[0])
	maior = (numbers[len(numbers)-1])
	var soma int
	maiornegativo := (numbers[len(numbers)-1])
	menornegativo := (numbers[0])
	for _, valor := range numbers {
		if valor < 0 {
			menor = maiornegativo
			maior = menornegativo

		}
		soma += valor
	}
	media := float64(soma) / float64(len(numbers))

	return maior, menor, media, nil
}
