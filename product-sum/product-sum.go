package productsum

type SpecialArray []interface{}

// Time: O(n) where n is the number of elements in the array including the subarrays
// Space: O(d) where d is the greatest depth of "special" arrays in the array
func ProductSum(arr []interface{}) int {
	return findProductSum(arr, 1)
}

func findProductSum(arr []interface{}, depth int) int {
	sum := 0

	for _, l := range arr {
		if SpecialArray, yes := l.(SpecialArray); yes {
			sum += findProductSum(SpecialArray, depth+1)
		} else if num, yes := l.(int); yes {
			sum += num
		}
	}
	return sum * depth
}
