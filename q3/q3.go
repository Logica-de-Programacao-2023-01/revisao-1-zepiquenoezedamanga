package q3

import "fmt"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {

	maior := -9999999999
	menor := 99999999999
	var soma float64
	var contagem float64
	var media float64
	if len(numbers) == 0 {
		return 0, 0, 0, fmt.Errorf("Lista vazia")
	}
	for i := 0; i < len(numbers); i++ {
		soma += float64(numbers[i])
		contagem++
		if numbers[i] > maior {
			maior = numbers[i]
		}
		if numbers[i] < menor {
			menor = numbers[i]
		}
	}
	media = (soma - float64(menor) - float64(maior)) / (contagem - 2)
	return maior, menor, media, nil
}
