package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	text := os.Args[1]

	formattedText := split(text)

	for i := 0; i < len(formattedText); i++ {
		printStr(formattedText[i])
		if i != len(formattedText)-1 {
			printStr("---")
		}
	}
	printStr("\n")

}

func split(text string) []string {
	letters := []rune(text)
	var result []string
	var temp []string

	startIndex := 0
	for i := startIndex; i < len(letters); i++ {
		if letters[i] == ' ' {
			temp = append(temp, string(letters[startIndex:i]))
			startIndex = i + 1
		}
	}
	temp = append(temp, string(letters[startIndex:]))

	for i := 0; i < len(temp); i++ {
		if temp[i] != "" {
			result = append(result, temp[i])
		}
	}

	return result
}

func printStr(str string) {
	for i := 0; i < len(str); i++ {
		z01.PrintRune(rune(str[i]))
	}
}
