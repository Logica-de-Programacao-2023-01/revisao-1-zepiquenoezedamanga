package q3

import (
	"fmt"
)

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, fmt.Errorf("lista vazia")
	}
	if len(numbers) == 1 {
		return numbers[0], numbers[0], float64(numbers[0]), nil
	}
	maior := -9999999999
	menor := 99999999999
	var soma int
	var contagem int
	for i := 1; i < len(numbers)-1; i++ {
		soma += numbers[i]
		contagem++
		if numbers[i] > maior {
			maior = numbers[i]
		}
		if numbers[i] < menor {
			menor = numbers[i]
		}
	}
	media := float64(soma) / float64(len(numbers)-2)

	return maior, menor, media, nil
}
