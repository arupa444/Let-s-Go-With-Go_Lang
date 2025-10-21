package main

import (
	"strings"
	"testing"
)

// Approach 1: Using strings.ContainsRune
func countVowelsContains(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, ch := range s {
		if strings.ContainsRune(vowels, ch) {
			count++
		}
	}
	return count
}

// Approach 2: Using switch
func countVowelsSwitch(s string) int {
	count := 0
	for _, ch := range s {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u',
			'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}

func BenchmarkContains(b *testing.B) {
	s := "GoLang is a fast and fun language for building great systems!"
	for i := 0; i < b.N; i++ {
		countVowelsContains(s)
	}
}

func BenchmarkSwitch(b *testing.B) {
	s := "GoLang is a fast and fun language for building great systems!"
	for i := 0; i < b.N; i++ {
		countVowelsSwitch(s)
	}
}
