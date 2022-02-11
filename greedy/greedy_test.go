package greedy

import "testing"

func TestSolveFractionalKnapsack(t *testing.T) {
	weights := []int{10, 40, 20, 30}
	values := []int{60, 40, 100, 120}
	knapsackCapacity := 50
	if val := SolveFractionalKnapsack(weights, values, knapsackCapacity); val != 240 {
		t.Errorf("Fractional Knapsack returned total value other than 240: %d", val)
	}
}

// Should result in
// f -> 00
// c -> 0100
// d -> 0101
// a -> 01100
// b -> 01101
// e -> 0111
func TestHuffmanCode(t *testing.T) {
	characters := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	frequencies := []int{5, 9, 12, 13, 16, 45}

	HuffmanNodes := HuffmanCode(characters, frequencies)
	PrintHuffmanNodes(HuffmanNodes[0], "")
}
