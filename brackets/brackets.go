package main

func main() {

}

func isBracket(ch rune) bool {
	brackets := "{}()[]"

	for _, b := range brackets {
		if ch == b {
			return true
		}
	}
	return false
}

func isMatchingPair(opening, closing rune) bool {
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	return pairs[opening] == closing

}

func isCorrectlyBracketed(expression string) bool {
	stack := []rune{}

	for _, ch := range expression {
		if isBracket(ch) {
			if len(stack) == 0 {
				stack = append(stack, ch)
			} else {
				top := stack[len(stack)-1]
				if isMatchingPair(top, ch) {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, ch)
				}
			}

		}

	}
	return len(stack) == 0

}
