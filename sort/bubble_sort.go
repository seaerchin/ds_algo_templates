package sort

// BubbleSort works through comparing adjacent elements
func BubbleSort(arr []int) (sorted []int) {
	sorted = make([]int, len(arr))
	_ = copy(sorted, arr)
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if sorted[j] < sorted[j-1] { 
				sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
			}
		}
	}
	return sorted	
}