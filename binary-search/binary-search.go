package binarysearch

import (
	"sort"
)

func IterativebinarySearch(arr []int, target int) int {
	sort.Ints(arr)
	low := 0
	high := len(arr)

	for low <= high {
		mid := low + (high-low)/2

		if mid == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func RecursiveBinarySearch(arr []int, low int, high int, target int) int {
	if high >= low {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			return RecursiveBinarySearch(arr, low, mid-1, target)
		}
		return RecursiveBinarySearch(arr, mid+1, high, target)
	}
	return -1
}
