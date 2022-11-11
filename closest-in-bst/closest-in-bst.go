package closestinbst

type BST struct {
	Val   int
	Left  *BST
	Right *BST
}

/*
	Average: O(log(n)) time | O(1) space

Worst: O(n) time | O(1) space
*/
func (tree *BST) findClosestValueIterative(element int) int {
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

func (tree *BST) findClosestValueRecursive(element int) int {
	return findClosestValLogic(tree, element, tree.Val)
}

func findClosestValLogic(tree *BST, element int, closest int) int {
	if absoluteDiff(tree.Val, element) < absoluteDiff(closest, element) {
		closest = tree.Val
	}

	if element < tree.Val && tree.Left != nil {
		return findClosestValLogic(tree.Left, element, closest)
	}

	if element > tree.Val && tree.Right != nil {
		return findClosestValLogic(tree.Right, element, closest)
	}

	return closest
}
