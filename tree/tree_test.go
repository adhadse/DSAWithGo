package tree

import (
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root := &Node{1, nil, nil}
	root.left = &Node{2, nil, nil}
	root.right = &Node{3, nil, nil}
	root.left.left = &Node{4, nil, nil}
	root.left.right = &Node{5, nil, nil}
	InorderTraversal(root)
}
