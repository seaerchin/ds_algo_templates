package sort

// CountingSort is a sorting algorithm that sorts integers; counting sort is NOT a comparison sort
func CountingSort(arr []int, largest int) []int {
	// we need largest to know the size of our counting array
	counts := make([]int, largest)
	// initiate the counts of integers in our array 
	for _, v := range arr {
		counts[v]++
	}
	// form a prefix array
	pref := make([]int, largest)
	for i := 1; i < largest; i++ {
		pref[i] = pref[i - 1] + counts[i]
	}
	res := make([]int, len(arr))
	for j := len(arr); j > 0; j-- {
		res[pref[arr[j]]] = arr[j]
		pref[arr[j]]--
	}
	return res
}