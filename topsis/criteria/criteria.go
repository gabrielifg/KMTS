package criteria

import (
	"math"
)

// Ideal é a função que determina a solução ideal positiva e a solução ideal negativa de uma matriz ponderada
func Ideal(weighted [][]float64) ([]float64, []float64) {
	// Obter o número de linhas e colunas da matriz
	rows := len(weighted)
	cols := len(weighted[0])

	// Criar dois vetores vazios para armazenar a solução ideal positiva e a solução ideal negativa
	positive := make([]float64, cols)
	negative := make([]float64, cols)

	// Para cada coluna da matriz, determinar o maior e o menor valor
	for j := 0; j < cols; j++ {
		max := 0.0
		min := math.MaxFloat64
		for i := 0; i < rows; i++ {
			if weighted[i][j] > max {
				max = weighted[i][j]
			}
			if weighted[i][j] < min {
				min = weighted[i][j]
			}
		}

		// Atribuir o maior valor ao vetor da solução ideal positiva e o menor valor ao vetor da solução ideal negativa
		positive[j] = max
		negative[j] = min
	}

	// Retornar os vetores da solução ideal positiva e da solução ideal negativa
	return positive, negative
}

// Proximity é a função que calcula o índice de proximidade relativa de cada alternativa à solução ideal positiva
func Proximity(weighted [][]float64, positive []float64, negative []float64) []float64 {
	// Obter o número de linhas da matriz
	rows := len(weighted)

	// Criar um vetor vazio para armazenar os índices de proximidade relativa
	indices := make([]float64, rows)

	// Para cada linha da matriz, calcular a distância euclidiana à solução ideal positiva e à solução ideal negativa
	for i := 0; i < rows; i++ {
		dplus := distance(weighted[i], positive)
		dminus := distance(weighted[i], negative)

		// Calcular o índice de proximidade relativa, usando a fórmula: C_i^* = d_i^- / (d_i^- + d_i^+)
		indices[i] = dminus / (dminus + dplus)
	}

	// Retornar o vetor de índices de proximidade relativa
	return indices
}

// Distance é a função que calcula a distância euclidiana entre dois vetores
func Distance(v1 []float64, v2 []float64) float64 {
	// Verificar se os vetores têm o mesmo tamanho
	if len(v1) != len(v2) {
		panic("Os vetores devem ter o mesmo tamanho")
	}

	// Inicializar a variável que armazena a soma dos quadrados das diferenças entre os elementos dos vetores
	sum := 0.0

	// Para cada elemento dos vetores, calcular a diferença, elevar ao quadrado e somar ao acumulador
	for i := 0; i < len(v1); i++ {
		sum += math.Pow(v1[i]-v2[i], 2)
	}

	// Retornar a raiz quadrada da soma dos quadrados das diferenças
	return math.Sqrt(sum)
}
