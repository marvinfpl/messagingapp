package commons

// used for chat repository
func Contains(array[]uint, in uint) bool {
	for _, k := range array {
		if k == in {
			return true
		}
	}
	return false
}