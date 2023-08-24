package piscine

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	charSet := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' {
			return false
		}
		if charSet[char] {
			return false
		}
		charSet[char] = true
	}
	return true
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}
	baseMap := make(map[byte]int)
	for i, char := range base {
		baseMap[byte(char)] = i
	}
	value := 0
	for i := len(s) - 1; i >= 0; i-- {
		value += baseMap[s[i]] * power(len(base), len(s)-1-i)
	}
	return value
}

func power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
