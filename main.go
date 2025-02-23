package main

var cache = make(map[string]int)

func generateKey(a, b string) string {
	return a + "|" + b
}

func calculate(a, b string) int {
	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}
	if (a[0] == b[0]) {
		return LevenshteinDistance(a[1:], b[1:])
	}
	atail := a[1:]
	btail := b[1:]
	first := LevenshteinDistance(atail, b)
	second := LevenshteinDistance(a, btail)
	third := LevenshteinDistance(atail, btail)
	return min(first, second, third) + 1
}

func LevenshteinDistance(a, b string) int {
	hash := generateKey(a, b)
	value, exists := cache[hash]
	if exists {
		return value
	}
	result := calculate(a, b)
	cache[hash] = result
	return result
}
