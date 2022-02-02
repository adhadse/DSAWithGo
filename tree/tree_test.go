package tree

import (
	"testing"
)

// Result in  ->4 ->2 ->5 ->1 ->3
func TestInorderTraversal(t *testing.T) {
	root := &Node{1, nil, nil}
	root.left = &Node{2, nil, nil}
	root.right = &Node{3, nil, nil}
	root.left.left = &Node{4, nil, nil}
	root.left.right = &Node{5, nil, nil}
	InorderTraversal(root)
}
func BenchmarkInorderTraversal(b *testing.B) {
	root := &Node{1, nil, nil}
	root.left = &Node{2, nil, nil}
	root.right = &Node{3, nil, nil}
	root.left.left = &Node{4, nil, nil}
	root.left.right = &Node{5, nil, nil}
	for i := 0; i <= b.N; i++ {
		InorderTraversal(root)
	}
}

// Result in  ->1 ->2 ->4 ->5 ->3
func TestPreorderTraversal(t *testing.T) {
	root := &Node{1, nil, nil}
	root.left = &Node{2, nil, nil}
	root.right = &Node{3, nil, nil}
	root.left.left = &Node{4, nil, nil}
	root.left.right = &Node{5, nil, nil}
	PreorderTraversal(root)
}

func BenchmarkPreorderTraversal(b *testing.B) {
	root := &Node{1, nil, nil}
	root.left = &Node{2, nil, nil}
	root.right = &Node{3, nil, nil}
	root.left.left = &Node{4, nil, nil}
	root.left.right = &Node{5, nil, nil}
	for i := 0; i <= b.N; i++ {
		PreorderTraversal(root)
	}
}

// Result in  ->4 ->5 ->2 ->3 ->1
func TestPostorderTraversal(t *testing.T) {
	root := &Node{1, nil, nil}
	root.left = &Node{2, nil, nil}
	root.right = &Node{3, nil, nil}
	root.left.left = &Node{4, nil, nil}
	root.left.right = &Node{5, nil, nil}
	PostorderTraversal(root)
}

func BenchmarkPostorderTraversal(b *testing.B) {
	root := &Node{1, nil, nil}
	root.left = &Node{2, nil, nil}
	root.right = &Node{3, nil, nil}
	root.left.left = &Node{4, nil, nil}
	root.left.right = &Node{5, nil, nil}
	for i := 0; i <= b.N; i++ {
		PostorderTraversal(root)
	}
}
