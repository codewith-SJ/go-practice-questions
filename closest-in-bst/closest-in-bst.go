package closestinbst

import "math"

type BST struct {
	root *Node
}

type Node struct {
	key   byte
	left  *Node
	right *Node
}

func (t *BST) insert(data byte) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

func (n *Node) insert(data byte) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func findClosestValueInBst(tree *BST, target int) int {
	infVal := math.Inf(1)
	return findClosestValueInBstHelper(tree, target, int(infVal))
}

// Average: O(log(n)) time | O(log(n)) space
// Worst: O(n) time | O(n) space
// Rescurive
func findClosestValueInBstHelper(tree *BST, target int, closest int) int {
	if tree == nil {
		return closest
	}
	if math.Abs(float64(target)-float64(closest)) > math.Abs(float64(target)-float64(tree.root.key)) {
		closest = int(tree.root.key)
	}
	if target < int(tree.root.key) {
		return findClosestValueInBstHelper(tree.root.left, target, closest)
	} else if target > int(tree.root.key) {
		return findClosestValueInBstHelper(tree.root.right, target, closest)
	} else {
		return closest
	}
}

/*
	Average: O(log(n)) time | O(1) space

Worst: O(n) time | O(1) space
*/
func findClosestValueInBstHelperInterative(tree *BST, target int, closest int) int {
	currentNode := tree
	for currentNode != nil {
		if math.Abs(float64(target)-float64(closest)) > math.Abs(float64(target)-float64(currentNode.root.key)) {
			closest = int(currentNode.root.key)
		}
		if target < int(currentNode.root.key) {
			closest = int(currentNode.root.left.key)
		} else if target > int(currentNode.root.key) {
			closest = int(currentNode.root.right.key)
		} else {
			break
		}
	}
	return closest
}
