package sort

// InsertionSort operates similarly to sorting a hand of cards
// compare current index against all previous index values
// and when we find the correct index, we insert it there;
// in the worst case (when arr is sorted in reverse order), it takes O(N ^ 2) time
func InsertionSort(arr []int) (sorted []int) {
	sorted = make([]int, len(arr))
	_ = copy(sorted, arr)
	if len(sorted) < 2 {
		return sorted
	}
	// first item is always trivially sorted
	for i := 1; i < len(sorted); i++ {
		cur := i
		for cur > 0 && sorted[cur-1] > sorted[cur] {
			sorted[cur], sorted[cur-1] = sorted[cur-1], sorted[cur]
			cur--
		}
	}
	return sorted
}
