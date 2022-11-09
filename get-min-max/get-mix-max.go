package getminmax

import "fmt"

type Pair struct {
	min int
	max int
}

func (minmax Pair) getMinMax(arr []int, n int) Pair {
	var i int
	if n == 1 {
		minmax.max = arr[0]
		minmax.min = arr[0]
		return minmax
	}
	if arr[0] > arr[1] {
		minmax.max = arr[0]
		minmax.min = arr[1]
	} else {
		minmax.max = arr[1]
		minmax.min = arr[0]
	}

	for i = 2; i < n; i++ {
		if arr[i] > minmax.max {
			minmax.max = arr[i]
		} else if arr[i] < minmax.min {
			minmax.min = arr[i]
		}
	}
	return minmax
}

func main() {
	arr := []int{1000, 11, 445, 1, 330, 3000}
	arr_size := 6

	var minmax Pair
	minmax = minmax.getMinMax(arr, arr_size)

	fmt.Printf("Min element is: %v\n", minmax.min)
	fmt.Printf("Max ele is: %v\n", minmax.max)
}
