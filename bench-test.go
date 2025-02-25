package main

import (
	"strconv"
	"testing"
)

func generateAlphabeticalStrings(length int) (string, string) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	first := alphabet[:length]
	second := "zyxwvutsrqponmlkjihgfedcba"[:length]
	return first, second
}

func BenchTestRun(b *testing.B, res func(string, string) int, iterations int) {
	for i := 0; i <= iterations; i++ {
		first, second := generateAlphabeticalStrings(i)
		b.Run("dist_"+strconv.Itoa(i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				if got := res(first, second); got != i {
					b.Errorf("expected distance %d, got %d", i, got)
				}
			}
		})
	}
}
