package main

import "testing"

func BenchmarkSorting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareArray()
		sortInsertion(&src)
	}
}
