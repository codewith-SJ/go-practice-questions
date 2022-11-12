package smallestdifference

import (
	"math"
	"sort"
)

// Time: O(N · log(N) + M · log(M))
// Space: O(log(N) + log(M)) or O(N + M) -> Depends on the sorting algorithm
func smallestdifference(arr1 []int, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)

	smallestAbsDiff := math.Inf(1)

	smallestpair := []int{}

	p1, p2 := 0, 0

	for p1 < len(arr1) && p2 < len(arr2) {
		firstNum := arr1[p1]
		secondNum := arr2[p2]

		currAbsDiff := math.Abs(float64(firstNum) - float64(secondNum))

		if currAbsDiff < smallestAbsDiff {
			smallestAbsDiff = currAbsDiff
			smallestpair[0] = firstNum
			smallestpair[1] = secondNum
		}

		if firstNum == secondNum {
			break
		}

		if firstNum <= secondNum {
			p1++
		} else {
			p2++
		}
	}
	return smallestpair
}
