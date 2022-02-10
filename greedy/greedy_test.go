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
