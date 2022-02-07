package graph

// contains check if slice contains element
// or not
func contains(slice []interface{}, element interface{}) bool {
	for _, val := range slice {
		if val == element {
			return true
		}
	}
	return false
}
