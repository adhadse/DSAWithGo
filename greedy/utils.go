package greedy

func InsertSortCost(list []KnapsackItem) []KnapsackItem {
	for i := 1; i < len(list); i++ {
		key := list[i].cost
		j := i - 1
		for j >= 0 && key < list[j].cost {
			list[j+1] = list[j]
			j--
		}
		list[j+1].cost = key
	}
	return list
}
