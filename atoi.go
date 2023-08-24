package piscine

func Atoi(s string) int {
	sign := 1
	number := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '+' && i == 0 {
			continue
		}
		if s[i] == '-' && i == 0 {
			sign = -1
			continue
		}

		if isNumeric(s[i]) {
			number = number*10 + int(s[i]-'0')
		}

		if !isNumeric(s[i]) {
			return 0
		}

	}

	return number * sign

}

func isNumeric(b byte) bool {
	return b >= '0' && b <= '9'
}
