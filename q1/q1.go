package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	if currentPurchase <= 0 {
		return 0, fmt.Errorf("valor da compra invalido")
	}

	var desconto float64
	var soma float64
	var media float64

	for _, valor := range purchaseHistory {
		soma += valor

	}
	media = float64(soma) / float64(len(purchaseHistory))

	if soma > 1000.00 {
		desconto += 0.10
		currentPurchase = desconto

	} else if soma <= 1000.00 {
		desconto += 0.05
		currentPurchase = desconto

	} else if soma <= 500.00 {
		desconto += 0.02
		currentPurchase = desconto

	} else if soma == 0 {
		desconto += 0.10
		currentPurchase = desconto

	} else if media > 1000.00 {
		desconto += 0.20
		currentPurchase = desconto
	}

	return currentPurchase, nil
}

