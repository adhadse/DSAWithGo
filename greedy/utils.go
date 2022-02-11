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

func InsertSortFrequency(list []*HuffmanNode) []*HuffmanNode {
	for i := 1; i < len(list); i++ {
		item := list[i]
		key := list[i].frequency
		j := i - 1
		for j >= 0 && key < list[j].frequency {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = item
	}
	return list
}

func remove(slice []*HuffmanNode, s *HuffmanNode) []*HuffmanNode {
	for i := range slice {
		if slice[i] == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
