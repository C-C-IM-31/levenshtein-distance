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

	dp := make([][]int, aLen+1)
	for i := range dp {
		dp[i] = make([]int, bLen+1)
	}

	for i := 1; i <= aLen; i++ {
		dp[i][0] = i
		for j := 1; j <= bLen; j++ {
			dp[0][j] = j
			dpTempPrevious := dp[i-1]
			dpTemp := dp[i]
			if a[i-1] == b[j-1] {
				dpTemp[j] = dpTempPrevious[j-1]
			} else {
				dpTemp[j] = 1 + min(
					dpTemp[j-1],
					dpTempPrevious[j],
					dpTempPrevious[j-1],
				)
			}
		}
	}

	return dp[aLen][bLen]
}
