package piscine

func BinaryAddition(a int, b int) []int {
	sum := a + b

	var binaries []int

	for sum > 0 {
		binaries = append(binaries, sum%2)
		sum = sum / 2
	}

	return binaries
}
