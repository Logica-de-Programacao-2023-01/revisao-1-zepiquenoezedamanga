package q5

import "fmt"

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {

	var temperatura_convertida float64
	switch fromScale {

	case "C":
		switch toScale {
		case "C":
			temperatura_convertida = temp
		case "F":
			temperatura_convertida = temp*9/5 + 32
		case "K":
			temperatura_convertida = temp + 273.15
		default:
			return 0, fmt.Errorf("escala invalida")
		}

	case "F":
		switch toScale {
		case "C":
			temperatura_convertida = (temp - 32) * 5 / 9
		case "F":
			temperatura_convertida = temp
		case "K":
			temperatura_convertida = (temp-32)*5/9 + 273.15
		default:
			return 0, fmt.Errorf("escala invalida")
		}

	case "K":
		switch toScale {
		case "C":
			temperatura_convertida = temp - 273.15
		case "F":
			temperatura_convertida = (temp-273.15)*9/5 + 32
		case "K":
			temperatura_convertida = temp
		default:
			return 0, fmt.Errorf("escala invalida")
		}

	default:
		return 0, fmt.Errorf("escala invalida")
	}

	return temperatura_convertida, nil
}
