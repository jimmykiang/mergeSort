package multithreadNaive

import (
	"runtime"
	"sync"
)

var semaphoreChannel chan struct{}

// MergeSort performs the merge sort algorithm taking advantage of multiple processors.
// Naive but intuitive implementation.
func MergeSort(inputSlice []int64) {
	threads := runtime.NumCPU() - 1
	semaphoreChannel = make(chan struct{}, threads)
	defer close(semaphoreChannel)
	mergesort(inputSlice)
}

func mergesort(src []int64) {
	if len(src) <= 1 {
		return
	}

	mid := len(src) / 2

	left := make([]int64, mid)
	right := make([]int64, len(src)-mid)
	copy(left, src[:mid])
	copy(right, src[mid:])

	wg := sync.WaitGroup{}

	select {
	case semaphoreChannel <- struct{}{}:
		wg.Add(1)
		go func() {
			mergesort(left)
			<-semaphoreChannel
			wg.Done()
		}()
	default:
		// if the goRoutine cant be created then do the processing in this same thread.
		mergesort(left)
	}

	mergesort(right)
	wg.Wait()
	merge(src, left, right)
}

func merge(result, left, right []int64) {
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
}
