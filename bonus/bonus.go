package bonus

import "fmt"

func CalculateDamage(characterLevel, enemyLevel int) (int, error) {
	if characterLevel == enemyLevel {
		fmt.Println(characterLevel * 7)
	} else if characterLevel < enemyLevel {
		fmt.Println(characterLevel * 5)
	} else if enemyLevel > 100 {
		fmt.Println(characterLevel * 2)
	} else if characterLevel > enemyLevel {
		fmt.Printf(characterLevel * 10)
	} else if characterLevel > 100 {
		fmt.Println(characterLevel * 20)
	} else if characterLevel-enemyLevel > 20 {
		fmt.Printf(characterLevel * 5)
	} else if characterLevel-enemyLevel < 20 {
		fmt.Println(characterLevel * 2)
	} else if characterLevel < 0 || enemyLevel < 0 {
		return 0, fmt.Errorf("nível invalido")
	}
	return 0, nil
}
