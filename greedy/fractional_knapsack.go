package greedy

type KnapsackItem struct {
	weight int
	value  int
	cost   int // value / weight
}

func SolveFractionalKnapsack(weight, value []int, knapsackCapacity int) int {
	var knapsackItems []KnapsackItem
	for i := range weight {
		knapsackItems = append(knapsackItems, KnapsackItem{weight: weight[i],
			value: value[i], cost: value[i] / weight[i]})
	}

	knapsackItems = InsertSortCost(knapsackItems)

	totalValue := 0
	for i := range knapsackItems {
		itemWeight := knapsackItems[i].weight
		itemValue := knapsackItems[i].value
		if knapsackCapacity-itemWeight >= 0 {
			knapsackCapacity -= itemWeight
			totalValue += itemValue
		} else {
			fraction := knapsackCapacity / itemWeight
			totalValue += itemValue * fraction
			knapsackCapacity = knapsackCapacity - (itemWeight * fraction)
			break
		}
	}
	return totalValue
}
