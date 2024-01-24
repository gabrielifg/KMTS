package topsis

import (
	"topsis/criteria"
	"topsis/matrix"
)

// Apply é a função que aplica o algoritmo TOPSIS para um conjunto de alternativas e pesos dos critérios
func Apply(alternatives [][]float64, weights []float64) []float64 {
	// Normalizar a matriz de decisão
	normalized := matrix.Normalize(alternatives)

	// Multiplicar a matriz normalizada pelos pesos dos critérios
	weighted := matrix.Multiply(normalized, weights)

	// Determinar a solução ideal positiva e a solução ideal negativa
	positive, negative := criteria.Ideal(weighted)

	// Calcular o índice de proximidade relativa de cada alternativa à solução ideal positiva
	indices := criteria.Proximity(weighted, positive, negative)

	// Retornar o vetor de índices de proximidade relativa
	return indices
}

// Rank é a função que ordena as alternativas de acordo com o índice de proximidade relativa, em ordem decrescente
func Rank(indices []float64) []int {
	// Retornar o vetor de posições das alternativas, ordenado pelo índice de proximidade relativa
	return matrix.Sort(indices)
}
