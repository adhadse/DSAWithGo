package recursion

import (
	"fmt"
	"testing"
)

func TestTailRecursive(t *testing.T) {
	fmt.Println(TailRecursive(10, 1))
}

func TestNonTailRecursive(t *testing.T) {
	fmt.Println(NonTailRecursive(10))
}

func TestRodCutting(t *testing.T) {
	pricesPerLength := []int{1, 5, 8, 9, 10, 17, 17, 20}
	if val := RodCutting(pricesPerLength, len(pricesPerLength)); val != 22 {
		t.Errorf("Function returned value other than 22: %d", val)
	}
}

// TowerOfHanoi(3, 'a', 'b', 'c')
// Move disk 1 from a -> b
// Move disk 2 from a -> c
// Move disk 1 from b -> c
// Move disk 3 from a -> b
// Move disk 1 from c -> a
// Move disk 2 from c -> b
// Move disk 1 from a -> b
func TestTowerOfHanoi(t *testing.T) {
	TowerOfHanoi(3, "a", "b", "c")
}

func TestJosephus(t *testing.T) {
	if val := Josephus(41, 2); val != 19 {
		t.Errorf("Function returned value other than 19 for input (41, 2): %d", val)
	}
}
