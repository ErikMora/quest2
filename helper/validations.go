package helper

func ValidateSizeAndStatus(size string, list map[int]string) bool {
	for _, value := range list {
		if size == value {
			return true
		}
	}
	return false
}
