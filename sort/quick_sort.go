package sort

// QuickSort asymptotic approximations:
// O(sqr(N))
// ~(N * lg N)
// o(N * lg N)
func QuickSort(src *[]int32, startIndex int, endIndex int) {
	if startIndex < endIndex {
		internalIndex := partition(src, startIndex, endIndex)

		QuickSort(src, startIndex, internalIndex-1)
		QuickSort(src, internalIndex+1, endIndex)
	}
}

// Split array into 2 arrays: all elements of the first one are not greater than the baseValue, all elements of the second one are not lower than the baseValue
// O(N)
func partition(src *[]int32, startIndex int, endIndex int) int {
	result := *src
	baseValue := result[endIndex]

	i := startIndex - 1

	// Loop invariant: at the beginning of every iteration:
	// - every element of the subarray result[startIndex, i] is lower or equals the baseValue (if i >= startIndex)
	// - every element of the subarray result[i+1, j-1] is greater or equals the baseValue (if i+1 <= j-1)
	// - the last element of the array result[endIndex] always equals the baseValue
	for j := startIndex; j < endIndex; j++ {
		if result[j] <= baseValue {
			i++
			tmp := result[j]
			result[j] = result[i]
			result[i] = tmp
		}
	}

	i++
	result[endIndex] = result[i]
	result[i] = baseValue

	return i
}
