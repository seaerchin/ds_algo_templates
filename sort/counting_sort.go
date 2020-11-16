package sort

// CountingSort is a sorting algorithm that sorts integers; counting sort is NOT a comparison sort
func CountingSort(arr []int, largest int) []int {
	// we need largest to know the size of our counting array
	counts := make([]int, largest + 1)
	// initiate the counts of integers in our array 
	for _, v := range arr {
		counts[v]++
	}
	// form a prefix array
	for i := 1; i < largest + 1; i++ {
		counts[i] = counts[i - 1] + counts[i]
	}
	res := make([]int, len(arr))
	for j:= 0; j < len(arr); j++ {
		res[counts[arr[j]] - 1] = arr[j]
		counts[arr[j]]--
	}
	return res
}