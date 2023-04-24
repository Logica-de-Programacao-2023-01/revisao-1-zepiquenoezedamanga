package bonus

import (
	"fmt"
	"math"
)
Um dia, três melhores amigos - Pedro, Vanessa e Tônia - decidiram formar uma equipe e participar de concursos de programação. Os participantes geralmente recebem vários problemas durante esses concursos. Muito antes do início, os amigos decidiram que implementariam um problema somente se pelo menos dois deles tivessem certeza da solução. Caso contrário, os amigos não escreveriam a solução do problema.

Este concurso oferece n problemas aos participantes. Para cada problema, sabemos qual amigo tem certeza da solução. Você receberá uma matriz booleana de n linhas e 3 colunas, em que a i-ésima linha representa as opiniões de Pedro, Vanessa e Tônia, respectivamente, sobre o i-ésimo problema. O valor "true" indica que o amigo tem certeza da solução, e "false" indica o contrário.

Ajude os amigos a encontrar o número de problemas para os quais eles escreverão uma solução.


func CalculateDamage(characterLevel, enemyLevel int) (int, error) {

	if characterLevel <= 0 || enemyLevel <= 0 {
		return 0, fmt.Errorf("level inválido ")
	}

	var damage int

	if characterLevel > 100 {
		damage = characterLevel * 20
	} else if enemyLevel > 100 {
		damage = characterLevel * 2
	} else {
		switch {
		case characterLevel > enemyLevel:
			damage = characterLevel * 10
		case characterLevel < enemyLevel:
			damage = characterLevel * 5
		default:
			damage = characterLevel * 7
		}

		levelDiff := int(math.Abs(float64(characterLevel - enemyLevel)))

		if levelDiff > 20 {
			damage *= 5
		} else if levelDiff < 20 {
			damage *= 2
		}
	}

	return damage, nil
}
