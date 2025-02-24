package main

func get(matrix [][]int, row, col int) int {
	if row == 0 && col == 0 {
		return 0
	}
	if row == 0 && col > 0 {
		return col
	}
	if col == 0 && row > 0 {
		return row
	}
	return matrix[row-1][col-1]
}

func LevenshteinDistanceFullMatrix(a, b string) int {
	rows, cols := len(a), len(b)
	matrix := make([][]int, rows)
	for i := 1; i < rows+1; i++ {
		row := make([]int, cols)
		matrix[i-1] = row
		for j := 1; j < cols+1; j++ {
			first := get(matrix, i, j-1) + 1
			second := get(matrix, i-1, j) + 1
			third := get(matrix, i-1, j-1)
			if a[i-1] != b[j-1] {
				third += 1
			}
			row[j-1] = min(first, second, third)
		}
	}
	return get(matrix, rows, cols)
}
