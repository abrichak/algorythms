package main

import "testing"

func BenchmarkSortingBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareArray(bestCase)
		sortInsertion(&src)
	}
}

func BenchmarkSortingAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareArray(averageCase)
		sortInsertion(&src)
	}
}

func BenchmarkSortingWorst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := prepareArray(worstCase)
		sortInsertion(&src)
	}
}

func BenchmarkOnlyPrepareArrayAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prepareArray(averageCase)
	}
}
