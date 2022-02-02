package tree

import "fmt"

// PreorderTraversal performs a pre order traversal of tree
// starting from node pointer `node`.
//
//  1. Visit the root.
//
//  2. Traverse the left subtree, i.e., call Preorder(left-subtree)
//
//   3. Traverse the right subtree, i.e., call Preorder(right-subtree)
func PreorderTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf(" ->%d", node.val)
	PreorderTraversal(node.left)
	PreorderTraversal(node.right)
}
