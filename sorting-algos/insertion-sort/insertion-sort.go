package insertionsort

// Time: Best case: O(n) if the array is already sorted | O(n^2) since we're running two loops in case array not sorted
// Space: O(1) since we're not using any extra space
func insetionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			swap(arr, j, j-1)
		}
		j--
	}
	return arr
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
