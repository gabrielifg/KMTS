// Arquivo topsis/matrix/matrix.go
package matrix

// Normalize é a função que normaliza uma matriz de decisão
func Normalize(matrix [][]float64) [][]float64 {
	// Obter o número de linhas e colunas da matriz
	rows := len(matrix)
	cols := len(matrix[0])

	// Criar uma matriz vazia para armazenar a matriz normalizada
	normalized := make([][]float64, rows)
	for i := range normalized {
		normalized[i] = make([]float64, cols)
	}

	// Para cada coluna da matriz, calcular o valor máximo
	for j := 0; j < cols; j++ {
		max := 0.0
		for i := 0; i < rows; i++ {
			if matrix[i][j] > max {
				max = matrix[i][j]
			}
		}

		// Dividir cada elemento da coluna pelo valor máximo, obtendo a matriz normalizada
		for i := 0; i < rows; i++ {
			normalized[i][j] = matrix[i][j] / max
		}
	}

	// Retornar a matriz normalizada
	return normalized
}

// Multiply é a função que multiplica uma matriz por um vetor de pesos
func Multiply(matrix [][]float64, weights []float64) [][]float64 {
	// Obter o número de linhas e colunas da matriz
	rows := len(matrix)
	cols := len(matrix[0])

	// Verificar se o número de colunas da matriz é igual ao número de elementos do vetor de pesos
	if cols != len(weights) {
		panic("O número de colunas da matriz deve ser igual ao número de elementos do vetor de pesos")
	}

	// Criar uma matriz vazia para armazenar a matriz ponderada
	weighted := make([][]float64, rows)
	for i := range weighted {
		weighted[i] = make([]float64, cols)
	}

	// Multiplicar cada elemento da matriz pelo peso correspondente, obtendo a matriz ponderada
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			weighted[i][j] = matrix[i][j] * weights[j]
		}
	}

	// Retornar a matriz ponderada
	return weighted
}

// Sort é a função que ordena um vetor de índices de proximidade relativa e retorna um vetor de posições das alternativas, em ordem decrescente
func Sort(indices []float64) []int {
	// Obter o número de elementos do vetor
	size := len(indices)

	// Criar um vetor para armazenar as posições das alternativas
	positions := make([]int, size)

	// Inicializar o vetor de posições com os valores de 1 a size
	for i := 0; i < size; i++ {
		positions[i] = i + 1
	}

	// Ordenar o vetor de índices e o vetor de posições, usando o algoritmo de ordenação por seleção
	for i := 0; i < size-1; i++ {
		max := i
		for j := i + 1; j < size; j++ {
			if indices[j] > indices[max] {
				max = j
			}
		}
		indices[i], indices[max] = indices[max], indices[i]
		positions[i], positions[max] = positions[max], positions[i]
	}

	// Retornar o vetor de posições das alternativas, ordenado pelo índice de proximidade relativa
	return positions
}
