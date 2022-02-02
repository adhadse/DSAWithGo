package tree

import "fmt"

// InorderTraversal performs a in order traversal of tree
// starting from node pointer `node`.
//
// 1. Traverse the left subtree, i.e., call Inorder(left-subtree)
//
// 2. Visit the root.
//
// 3. Traverse the right subtree, i.e., call Inorder(right-subtree)
func InorderTraversal(node *Node) {
	if node == nil {
		return
	}
	InorderTraversal(node.left)
	fmt.Printf(" ->%d", node.val)
	InorderTraversal(node.right)
}
