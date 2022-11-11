package palindromechecking

import "strings"

// Idea - 1: Reverse the input string and check if its reversed string is equal to it or not. Here, storing the reversed string as a string itself
func palindromechecking(s string) bool {
	var reverse strings.Builder

	for i := len(s) - 1; i >= 0; i-- {
		reverse.WriteByte(s[i])
	}

	return reverse.String() == s
}

// Idea - 2: Same as above but storing the reversed string as an array

func palindromecheckingIdea2(s string) bool {
	var strArr = []string{}

	for i := len(s) - 1; i >= 0; i-- {
		strArr = append(strArr, string(s[i]))
	}
	return strings.Join(strArr, "") == s
}

// Idea - 3: (Important) Two pointer alternative -> A palindrome string is a mirror image from left and right. If we start from the centre and then go outwards
// both ways then we should encounter the same characters on both sides

// Time: O(n); n is length of string
// Space: O(1)
func palindromecheckingIdea3(s string) bool {
	leftLoc := 0
	rightLoc := len(s) - 1

	for leftLoc < rightLoc {
		if leftLoc != rightLoc {
			return false
		}
		leftLoc++
		rightLoc--
	}
	return true
}

// Same as above but recursive
// Time: O(n)
// Space: O(n)
func palindromecheckingIdea4(s string, left int) bool {
	// left = 0
	right := len(s) - 1 - left

	if left >= right {
		return true
	}

	return s[left] == s[right] && palindromecheckingIdea4(s, left+1)
}

// Tail call optimization
/*
The Idea 4 can be optimized by converting the function into a tail recursive function by changing the recursion to the end of the function itself.
This way the recursion takes constant space
*/

// Time: O(n)
// Space: O(1) if tail optimization is supported else O(n)
func palindromecheckingIdea5(s string, left int) bool {
	right := len(s) - 1 - left
	if left >= right {
		return true
	}
	if s[left] != s[right] {
		return false
	}

	return palindromecheckingIdea5(s, left+1)
}
