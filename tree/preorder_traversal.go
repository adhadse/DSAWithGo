package tree

import "fmt"

// PreorderTraversal
func PreorderTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf(" ->%d", node.val) // Check current node
	PreorderTraversal(node.left)  // then recur on left subtree
	PreorderTraversal(node.right) // at last recur on right subtree
}
