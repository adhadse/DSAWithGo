package tree

import "fmt"

func InorderTraversal(root *Node) {
	if root == nil {
		return
	}
	InorderTraversal(root.left)   // First recur on left child
	fmt.Printf(" ->%d", root.val) // Check current node
	InorderTraversal(root.right)  // Now recur on right child
}
