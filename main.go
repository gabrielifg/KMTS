package kmts

import (
	"fmt"
	topsis "kmts/topsis/criteria"
)

func main() {
	// Definir os pesos dos critérios
	weights := []float64{0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 1, 1, 1, 1, 1, 1, 1}

	// Definir as alternativas de alocação dos multi-tenancys
	alternatives := [][]float64{
		{8, 4, 16, 8, 100, 50, 1, 1, 1, 1, 1, 1, 1},     // Alternativa 1
		{4, 2, 8, 4, 50, 25, 0, 0, 0, 0, 0, 0, 0},       // Alternativa 2
		{6, 3, 12, 6, 75, 37.5, 1, 0, 1, 0, 1, 0, 1},    // Alternativa 3
		{10, 5, 20, 10, 125, 62.5, 0, 1, 0, 1, 0, 1, 0}, // Alternativa 4
	}

	// Aplicar o algoritmo TOPSIS
	indices := topsis.Apply(alternatives, weights)

	// Mostrar os índices de proximidade relativa de cada alternativa
	fmt.Println("Índices de proximidade relativa:")
	for i, index := range indices {
		fmt.Printf("Alternativa %d: %.4f\n", i+1, index)
	}

	// Ordenar as alternativas de acordo com o índice de proximidade relativa, em ordem decrescente
	ranking := topsis.Rank(indices)

	// Mostrar o ranking das alternativas
	fmt.Println("Ranking das alternativas:")
	for i, position := range ranking {
		fmt.Printf("%dº lugar: Alternativa %d\n", i+1, position)
	}
}
