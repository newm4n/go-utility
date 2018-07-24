package go_utility

func ArrayContainString(arr []string, component string) bool {
	for _, v := range arr {
		if v == component {
			return true
		}
	}
	return false
}
