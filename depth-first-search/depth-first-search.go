package depthfirstsearch

type Node struct {
	Name     string
	Children []*Node
}

// Time: O(V+E) - We traverse every vertex. At every vertex, we traverse through the edges of the vertex as well (by calling DepthFirstSearch)
// Space: O(V) - We store the vertices in the call stack

func (n *Node) DepthFirstSearch(arr []string) []string {
	arr = append(arr, n.Name)

	for _, child := range n.Children {
		arr = child.DepthFirstSearch(arr)
	}

	return arr
}
