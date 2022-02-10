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
		currentWeight := knapsackItems[i].weight
		currentValue := knapsackItems[i].value
		if knapsackCapacity-currentWeight >= 0 {
			knapsackCapacity -= currentWeight
			totalValue += currentValue
		} else {
			fraction := knapsackCapacity / currentWeight
			totalValue += currentValue * fraction
			knapsackCapacity = knapsackCapacity - (currentWeight * fraction)
			break
		}
	}
	return totalValue
}
