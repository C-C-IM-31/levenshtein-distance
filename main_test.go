package main

import (
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	for _, tc := range []struct {
		name   string
		a, b   string
		wanted int
	}{
		{name: "empty", a: "", b: "", wanted: 0},
		{name: "equal", a: "abc", b: "abc", wanted: 0},
		{name: "insertions", a: "ab", b: "abc", wanted: 1},
		{name: "insertions-2", a: "ab", b: "dabc", wanted: 2},
		{name: "deletions", a: "abc", b: "ab", wanted: 1},
		{name: "editions", a: "abc", b: "adc", wanted: 1},
		{name: "sentence", a: "cat and dog", b: "dog and cat", wanted: 6},
	} {
		// START OMIT
		t.Run(tc.name, func(t *testing.T) {
			if got := LevenshteinDistance(tc.a, tc.b); got != tc.wanted {
				t.Errorf("wanted %d, got %d", tc.wanted, got)
			}
			if got := LevenshteinDistance(tc.b, tc.a); got != tc.wanted {
				t.Errorf("inverted - wanted %d, got %d", tc.wanted, got)
			}
		})
		// END OMIT
	}
}

func BenchmarkLevenshteinDistance(b *testing.B) {
	BenchTestRun(b, LevenshteinDistance, Iterations)
}
