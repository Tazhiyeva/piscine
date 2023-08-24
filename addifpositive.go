package piscine

func AddIfPositive(a int, b int) int {
	if isPositive(a) && isPositive(b) {
		return a + b
	}
	return 0
}

func isPositive(a int) bool {
	return a >= 0
}
