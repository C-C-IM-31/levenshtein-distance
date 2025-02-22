package main

func LevenshteinDistance(a, b string) int {
	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}

	x := LevenshteinDistance(a, b[1:]) + 1
	y := LevenshteinDistance(a[1:], b) + 1
	z := LevenshteinDistance(a[1:], b[1:])
	if a[0] != b[0] {
		z++
	}

	return min(x, y, z)
}
