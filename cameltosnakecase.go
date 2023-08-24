package piscine

func CamelToSnakeCase(str string) string {
	var snakeString string

	if !isCamelCase(str) {
		return str
	}
	startIndex := 0
	for i := 1; i < len(str); i++ {
		if isUpper(str[i]) {
			snakeString = snakeString + string(str[startIndex:i])
			if i != len(str)-1 {
				snakeString = snakeString + "_"
			}
			startIndex = i
		}

	}
	snakeString = snakeString + str[startIndex:]
	return snakeString

}

func isCamelCase(str string) bool {
	if isUpper(str[len(str)-1]) {
		return false
	}
	for i := 0; i < len(str)-1; i++ {
		if isUpper(str[i]) && isUpper(str[i+1]) {
			return false
		}
		if !isAlpha(str[i]) {
			return false
		}
	}
	return true
}

func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
