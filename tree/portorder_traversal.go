package tree

import "fmt"

// PostorderTraversal
func PostorderTraversal(node *Node) {
	if node == nil {
		return
	}
	PostorderTraversal(node.left)  // First recur on left subtree
	PostorderTraversal(node.right) // then recur on right subtree
	fmt.Printf(" ->%d", node.val)  // at last Check current node
}
