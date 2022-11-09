package branchsums

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// Recursive method DFS
// Time: O(n) since we have n nodes and traversing n node takes n time
// Space: O(n) since space for n nodes is n.
func calculateBranchSums(node *BinaryTree, runningSum int, sums *[]int) {
	if node == nil {
		return
	}

	runningSum += node.Value

	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, runningSum)
		return
	}

	calculateBranchSums(node.Left, runningSum, sums)
	calculateBranchSums(node.Right, runningSum, sums)
}

// Iterative method DFS

// Time: O(n) since we have n nodes and traversing n node takes n time
// Space: O(n) since space for n nodes is n.
type unvisitedNode struct {
	node       *BinaryTree
	runningSum int
}

func BranchSums(root *BinaryTree) []int {
	sums := []int{}
	stack := []unvisitedNode{unvisitedNode{root, 0}}

	for len(stack) > 0 {
		last := len(stack) - 1
		node, runningSum := stack[last].node, stack[last].runningSum
		stack = stack[:last]

		runningSum += node.Value
		if node.Left == nil && node.Right == nil {
			sums = append(sums, runningSum)
			continue
		}
		if node.Right != nil {
			stack = append(stack, unvisitedNode{node.Right, runningSum})
		}

		if node.Left != nil {
			stack = append(stack, unvisitedNode{node.Left, runningSum})
		}
	}
	return sums
}
