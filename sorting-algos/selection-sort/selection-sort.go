package main

// Time: O(n^2) since we have two loops
// Space: O(1) since we are not using any extra space
func selectionSort(arr []int) []int {
	start := 0

	for start < len(arr)-1 {
		smallest := start

		for i := start + 1; i < len(arr); i++ {
			if arr[i] < arr[smallest] {
				smallest = i
			}
		}

		if smallest != start {
			swap(arr, start, smallest)
		}
		start++
	}
	return arr
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
