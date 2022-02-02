package tree

import "fmt"

// PostorderTraversal performs a post order traversal of tree
// starting from node pointer `node`.
//
// 1. Traverse the left subtree, i.e., call Postorder(left-subtree)
//
// 2. Traverse the right subtree, i.e., call Postorder(right-subtree)
//
// 3. Visit the root.
func PostorderTraversal(node *Node) {
	if node == nil {
		return
	}
	PostorderTraversal(node.left)
	PostorderTraversal(node.right)
	fmt.Printf(" ->%d", node.val)
}
