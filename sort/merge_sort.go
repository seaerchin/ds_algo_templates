package sort

// MergeSort is a recursive algorithm
// it aims to break the problem down into smaller sub-problems
// and then, solve the sub-problems to solve the overall problem
func MergeSort(arr []int) (sorted []int) {
	// do note that slices in golang are a thin wrapper
	// consisting only of head_ptr, len, cap
	l := len(arr)
	if l <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left, right := MergeSort(arr[:mid]), MergeSort(arr[mid:])
	return merge(left, right)
}

// merge assumes that left and right are both already sorted
func merge(left, right []int) (sorted []int) {
	sorted = make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}
	sorted = append(sorted, left[i:]...)
	sorted = append(sorted, right[j:]...)
	return sorted
}
