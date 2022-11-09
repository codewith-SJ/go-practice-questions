package bubblesort

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// Time: O(n) in case array is sorted | O(n^2) average and worst case
// Space: O(1) since we're not using any extra space
func BubbleSort(arr []int) []int {

	isSorted := false

	elementsLeft := len(arr)

	for !isSorted {
		isSorted = true

		for i := 0; i < elementsLeft; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
				isSorted = false
			}
		}
		elementsLeft--
	}
	return arr
}
