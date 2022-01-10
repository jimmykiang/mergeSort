package multithreadImproved

import (
	"github.com/jimmykiang/mergesort/sorttest"
	"testing"
)

func TestMergesortImproved(t *testing.T) {
	sorttest.Test(MergeSort, 1000, t)
}

func BenchmarkMergeSortImproved(b *testing.B) {
	sorttest.Benchmark(MergeSort, 16_000_000, b)
}
