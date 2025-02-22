package main

//func LevenshteinDistance(a, b string) int {
//	if len(a) == 0 {
//		return len(b)
//	}
//	if len(b) == 0 {
//		return len(a)
//	}
//
//	x := LevenshteinDistance(a, b[1:]) + 1
//	y := LevenshteinDistance(a[1:], b) + 1
//	z := LevenshteinDistance(a[1:], b[1:])
//	if a[0] != b[0] {
//		z++
//	}
//
//	return min(x, y, z)
//}

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
