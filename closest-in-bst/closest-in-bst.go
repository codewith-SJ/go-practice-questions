package closestinbst

type BST struct {
	Val   int
	Left  *BST
	Right *BST
}

// Time: O(log(n)); n is the number of nodes | Worst case: O(n) when the tree only has one branch
// Space: O(1)
func (tree *BST) findClosestValue(element int) int {
	currentNode := tree
	closestNode := tree.Val

	for currentNode != nil {
		if absoluteDiff(currentNode.Val, element) < absoluteDiff(closestNode, element) {
			closestNode = currentNode.Val
		}

		if element < currentNode.Val {
			currentNode = currentNode.Left
		} else if element > currentNode.Val {
			currentNode = currentNode.Right
		} else {
			break
		}
	}
	return closestNode
}

func absoluteDiff(x int, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

/*
	Average: O(log(n)) time | O(1) space

Worst: O(n) time | O(1) space
*/
func findClosestValueInBstHelperInterative(tree *BST, target int, closest int) int {
	currentNode := tree
	for currentNode != nil {
		if absoluteDiff(currentNode.Val, target) < absoluteDiff(closest, target) {
			closest = currentNode.Val
		}
		if target < currentNode.Val {
			closest = currentNode.Left.Val
		} else if target > currentNode.Val {
			closest = currentNode.Right.Val
		} else {
			break
		}
	}
	return closest
}
