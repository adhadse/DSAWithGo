package greedy

import (
	"fmt"
	"strconv"
)

type HuffmanNode struct {
	symbol      string       // symbol name (character
	frequency   int          // frequency of symbol
	left        *HuffmanNode // HuffmanNode left of current node
	right       *HuffmanNode // HuffmanNode right of current node
	huffmanCode int          // tree direction (0/1)
}

// PrintHuffmanNodes is a utility function to print huffman codes for all
// symbols in the newly created huffman tree
func PrintHuffmanNodes(node *HuffmanNode, val string) {
	// huffman code for current node
	newVal := val + strconv.Itoa(node.huffmanCode)

	// If node is not an edge node
	// then traverse inside it
	if node.left != nil {
		PrintHuffmanNodes(node.left, newVal)
	}
	if node.right != nil {
		PrintHuffmanNodes(node.right, newVal)
	}
	// If node is edge node then
	// display its huffman code
	if !(node.left != nil) && !(node.right != nil) {
		fmt.Printf("%s -> %s\n", string(node.symbol), newVal)
	}
}

func HuffmanCode(characters []rune, frequencies []int) []*HuffmanNode {
	var nodes []*HuffmanNode

	// Converting characters and frequencies into
	// huffman tree nodes
	for i := range characters {
		nodes = append(nodes, &HuffmanNode{frequency: frequencies[i], symbol: string(characters[i])})
	}

	for len(nodes) > 1 {
		// sort all nodes in ascending order
		// based on their frequency
		nodes = InsertSortFrequency(nodes)

		// pick 2 smallest nodes
		left := nodes[0]
		right := nodes[1]

		// assign directional value to these nodes
		left.huffmanCode = 0
		right.huffmanCode = 1

		// combine the 2 smallest nodes to create
		// new node as their parent
		newNode := HuffmanNode{frequency: left.frequency + right.frequency,
			symbol: left.symbol + right.symbol,
			left:   left, right: right}

		// remove the 2 nodes and add their
		// parent as new node among others
		nodes = remove(nodes, nodes[0])
		nodes = remove(nodes, nodes[0])
		nodes = append(nodes, &newNode)
	}
	return nodes
}
