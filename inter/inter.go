package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	result := findCommonChars(str1, str2)
	print(result + "\n")

}

func findCommonChars(str1, str2 string) string {
	seen := make(map[rune]bool)
	result := ""

	for i := 0; i < len(str1); i++ {
		char := rune(str1[i])
		if !seen[char] && contains(str2, char) {
			result += string(char)
			seen[char] = true
		}
	}
	return result
}

func contains(str string, r rune) bool {
	for _, char := range str {
		if char == r {
			return true
		}
	}
	return false
}

func print(str string) {
	for i := 0; i < len(str); i++ {
		z01.PrintRune(rune(str[i]))
	}
}
