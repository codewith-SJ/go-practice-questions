package threelargestnums

import "math"

// Time: O(n) | Space: O(1)
func findThreeLargest(arr []int) []int {
	threeLargest := []int{math.MinInt64, math.MinInt64, math.MinInt64}
	for i := range arr {
		updateLargest(threeLargest, i)
	}
	return threeLargest
}

func updateLargest(arr []int, i int) {
	if arr[2] == math.MinInt64 || i > arr[2] {
		shiftAndUpdate(arr, i, 2)
	} else if arr[1] == math.MinInt64 || i > arr[1] {
		shiftAndUpdate(arr, i, 1)
	} else if arr[0] == math.MinInt64 || i > arr[0] {
		shiftAndUpdate(arr, i, 0)
	}
}

func shiftAndUpdate(arr []int, numToUpdate int, id int) {
	/*
		[0, 1, 2]
		[x, y, z] where 2 is the `id`. So when this function runs, the whole array becomes: x [y, z, numToUpdate]
	*/
	for i := 0; i <= id; i++ {
		if i == id {
			arr[i] = numToUpdate
		} else {
			arr[i] = arr[i+1]
		}
	}
}
