package main

import (
	"fmt"
	"testing"
)

func TestLevenshteinDistanceMemoised(t *testing.T) {
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
			if got := LevenshteinDistanceMemoised(tc.a, tc.b); got != tc.wanted {
				t.Errorf("wanted %d, got %d", tc.wanted, got)
			}
			if got := LevenshteinDistanceMemoised(tc.b, tc.a); got != tc.wanted {
				t.Errorf("inverted - wanted %d, got %d", tc.wanted, got)
			}
		})
		// END OMIT
	}
}

func BenchmarkLevenshteinDistanceMemoised(b *testing.B) {
	b.Run("short", func(b *testing.B) {
		for b.Loop() {
			if got := LevenshteinDistanceMemoised("abc", "def"); got == 0 {
				b.Error("non-zero distance")
			}
		}
	})
	b.Run("medium", func(b *testing.B) {
		for b.Loop() {
			if got := LevenshteinDistanceMemoised("abcedf", "defabc"); got == 0 {
				b.Error("non-zero distance")
			}
		}
	})
	b.Run("longer", func(b *testing.B) {
		for b.Loop() {
			if got := LevenshteinDistanceMemoised("abcedfabcedf", "vbcedfabcedf"); got == 0 {
				b.Error("non-zero distance")
			}
		}
	})
}

func ExampleLevenshteinDistanceMemoised() {
	fmt.Println(LevenshteinDistanceMemoised("abc", "adc"))
}

func main() {
	ExampleLevenshteinDistanceMemoised()
}
