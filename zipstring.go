package piscine

/*

Write a function that takes a string and returns a new string that replaces every character with the number of duplicates and the character itself, deleting the extra duplications.

The letters are from the latin alphabet list only. Any other character, symbols, shall not be tested.


*/

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	zipped := ""
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			zipped += (IntToString(count) + string(s[i-1]))
			count = 1
		}
	}
	zipped += IntToString(count) + string(s[len(s)-1])
	return zipped
}

func IntToString(n int) string {
	if n == 0 {
		return "0"
	}

	var result string
	var digits []rune
	for n > 0 {
		digit := rune(n % 10)
		digits = append(digits, digit)
		n /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		result += string(digits[i] + '0') // Convert rune to character
	}
	return result
}
