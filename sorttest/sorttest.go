package sorttest

import (
	"math/rand"
	"testing"
	"time"
)

// Test is a helper method to test sort implementations.
func Test(sort func([]int64), elements int, t *testing.T) {
	testSlice := createTestSlice(elements)

	sort(testSlice)

	for i := 0; i < len(testSlice)-1; i++ {
		if testSlice[i] >= testSlice[i+1] {
			// testSlice should be sorted in ascending order
			t.Fail()
		}
	}
}

// Benchmark is a helper method to benchmark the sort function algorithm.
func Benchmark(sort func([]int64), elements int, b *testing.B) {
	shuffled := createTestSlice(elements)
	src := make([]int64, elements)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(src, shuffled)
		b.StartTimer()
		sort(src)
	}
}

// createTestSlice returns a new slice of n random integers.
func createTestSlice(n int) []int64 {
	arr := make([]int64, n)
	rand.Seed(time.Now().Unix())
	for i := range arr {
		arr[i] = rand.Int63()
	}
	return arr
}
