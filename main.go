package main

func LevenshteinDistance(a, b string) int {
	rows, cols := len(a), len(b)
	previosRow := make([]int, cols + 1)
	currentRow := make([]int, cols + 1)
	for index := range cols {
		previosRow[index] = index
	}
	for i := 1; i < rows + 1; i++ {
		currentRow[0] = i
		for j := 1; j < cols + 1; j++ {
			first := currentRow[j - 1] + 1
			second := previosRow[j] + 1
			third := previosRow[j - 1]
			if a[i-1] != b[j-1] {
				third += 1
			}
			currentRow[j] = min(first, second, third)
		}
		copy(previosRow, currentRow)
	}
	return currentRow[cols]
}
