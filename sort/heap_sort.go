package sort

// HeapSort asymptotic approximations:
// O(N * lg N)
func HeapSort(src *[]int32) {
	buildMaxHeap(src)

	result := *src
	heapSize := len(result)

	// Loop invariant: at the beginning of every iteration:
	// - the subarray result[0, i-1] is a non-increasing heap made of the i lowest elements of array
	// - the subarray result[i, N-1] is a sorted array consisting of the N-i largest elements of array
	for i := heapSize - 1; i >= 1; i-- {
		tmp := result[i]
		result[i] = result[0]
		result[0] = tmp

		heapSize--
		maxHeapify(src, heapSize, 0)
	}
}

// Build non-increasing heap from a random array
// O(N)
func buildMaxHeap(src *[]int32) {
	result := *src
	heapSize := len(result)

	// Loop invariant: at the beginning of every iteration, every node result[i+1], result[i+2], ..., result[N-1]
	// is a root of a non-increasing heap
	for i := len(result) / 2; i >= 0; i-- {
		maxHeapify(src, heapSize, i)
	}
}

func maxHeapify(src *[]int32, heapSize int, i int) {
	result := *src
	largestIndex := i

	l := left(i)
	r := right(i)

	if l < heapSize && result[l] > result[i] {
		largestIndex = l
	}

	if r < heapSize && result[r] > result[largestIndex] {
		largestIndex = r
	}

	if i != largestIndex {
		tmp := result[i]
		result[i] = result[largestIndex]
		result[largestIndex] = tmp

		maxHeapify(src, heapSize, largestIndex)
	}
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}
