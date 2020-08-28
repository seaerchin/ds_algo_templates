package search

// BinarySearch requires that the array is already sorted before hand;
// it uses a pivot (the middle element in the array)
// to split the array into 2 portions and searches based on whether
// the target is greater/lesser than the pivot;
// binary search is efficient and has time complexity of O(log n)
func BinarySearch(arr []int, target int) (index int) {
	var low, high = 0, len(arr) - 1
	for low <= high {
		var mid = (low + high) / 2
		var pivot = arr[mid]
		if pivot == target {
			return mid
		}
		if target > pivot {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
