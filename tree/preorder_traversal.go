// Copyright (C) 2022   Anurag Dhadse. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
