package q3

import (
	"fmt"
	"sort"
)

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, fmt.Errorf("lista vazia")
	}
	if len(numbers) == 1 {
		return numbers[0], numbers[0], float64(numbers[0]), nil
	}
	var maior int
	var menor int
	sort.Ints(numbers)
	menor = (numbers[0])
	maior = (numbers[len(numbers)-1])
	if numbers[0] < 0 {
		maior = numbers[0]
		menor = (numbers[len(numbers)-1])
	}
	var soma int
	for i := 1; i < len(numbers)-1; i++ {
		soma += numbers[i]
	}
	media := float64(soma) / float64(len(numbers)-2)

	return maior, menor, media, nil
}
