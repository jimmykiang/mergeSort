package multithreadImproved

import (
	"sort"
	"sync"
)

func MergeSort(src []int64) {
	temp := make([]int64, len(src))
	mergesort(src, temp)
}

func mergesort(src, temp []int64) {
	// run the original sort.Slice implementation when there are a low number of entries to sort.
	// this is to reduce the overhead of creating goRoutines including the cases when the input slice lenght is 1.
	if len(src) <= 10000 {
		sort.Slice(src, func(i int, j int) bool { return src[i] <= src[j] })
		return
	}

	mid := len(src) / 2

	left, lTemp := src[:mid], temp[:mid]
	right, rTemp := src[mid:], temp[mid:]

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		mergesort(left, lTemp)
		wg.Done()
	}()

	mergesort(right, rTemp)

	wg.Wait()

	merge(src, temp, left, right)
}

func merge(src, result, left, right []int64) {
	var l, r, i int

	for l < len(left) || r < len(right) {
		if l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				result[i] = left[l]
				l++
			} else {
				result[i] = right[r]
				r++
			}
		} else if l < len(left) {
			result[i] = left[l]
			l++
		} else if r < len(right) {
			result[i] = right[r]
			r++
		}
		i++
	}

	// reduce the number of calls to copy slices.
	copy(src, result)
}
