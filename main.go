package main

func LevenshteinDistance(a, b string) int {
	aLen, bLen := len(a), len(b)

	prevRow := make([]int, bLen+1)
	currRow := make([]int, bLen+1)

	for index := range prevRow {
		prevRow[index] = index
	}

	for i := 1; i <= aLen; i++ {
		currRow[0] = i
		for j := 1; j <= bLen; j++ {
			prevRowTemp := prevRow[j-1]
			if a[i-1] == b[j-1] {
				currRow[j] = prevRowTemp
			} else {
				currRow[j] = min(currRow[j-1]+1, prevRow[j]+1, prevRowTemp+1)
			}
		}

		copy(prevRow, currRow)
	}

	return currRow[bLen]
}
