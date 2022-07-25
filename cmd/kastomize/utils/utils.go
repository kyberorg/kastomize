package utils

func SearchString(slice []string, str string) bool {
	for _, element := range slice {
		if element == str {
			return true
		}
	}
	return false
}
