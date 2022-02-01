package tree

import "fmt"

// InorderTraversal
func InorderTraversal(node *Node) {
	if node == nil {
		return
	}
	InorderTraversal(node.left)   // First recur on left child
	fmt.Printf(" ->%d", node.val) // Check current node
	InorderTraversal(node.right)  // Now recur on right child
}
