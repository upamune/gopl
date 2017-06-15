package ex03

import (
	"testing"
	"strings"
)

var t = strings.Split(strings.Repeat("a", 100), "")

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1(t)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2(t)
	}
}
