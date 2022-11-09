package twonumbersum

// O(n) time | O(n) space

func calculate(arr []int, targetsum int) []int {
	nums := make(map[int]int)
	for i, num := range arr {
		if j, ok := nums[targetsum-num]; ok {
			return []int{num, j}
		}
		nums[num] = i
	}
	return []int{}
}
