package utils

func SanitizeForPrintf(str string) string {
	var result string
	for i := range len(str) {
		if string(str[i]) == "\"" && string(str[i-1]) != "\\" {
			result += "\\\""
		} else {
			result += string(str[i])
		}
	}

	return result
}
