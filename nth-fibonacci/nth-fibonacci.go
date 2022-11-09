package fib

// Iterative method
// Time: O(n) | Space: O(1)
func iterativeFib(n int) int {
	lastTwoFib := []int{0, 1}

	count := 3

	for count <= n {
		next := lastTwoFib[0] + lastTwoFib[1]
		lastTwoFib[0] = lastTwoFib[1]
		lastTwoFib[1] = next
		count++
	}

	if n <= 1 {
		return lastTwoFib[0]
	} else {
		return lastTwoFib[1]
	}
}

// Naive Recursive approach
// Time: O(2^n) since in each step, we're going to recursively call the function twice
// Space: O(n)
func naiveRecursive(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	return naiveRecursive(n-1) + naiveRecursive(n-2)
}

// Memoized Recusrive function keeps track of what we have calculated in the previous step and passes it in the recursive call
// Time: O(n)
// Space: O(n)
func memoizedRecursive(n int) int {
	memo := map[int]int{
		1: 0,
		2: 1,
	}

	return solve(n, memo)
}

func solve(n int, memo map[int]int) int {
	if val, foundVal := memo[n]; foundVal {
		return val
	}

	memo[n] = solve(n-1, memo) + solve(n-2, memo)

	return memo[n]
}
