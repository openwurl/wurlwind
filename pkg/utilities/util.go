package utilities

// SliceContainsString searches a string for a slice and confirms if it exists
func SliceContainsString(searchPattern string, sliceToSearch []string) bool {
	for _, sliceObject := range sliceToSearch {
		if sliceObject == searchPattern {
			return true
		}
	}
	return false
}
