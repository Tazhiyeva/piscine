package main

import (
	"fmt"
	"os"
)

func hiddenp(s1, s2 string) int {
	indexS1 := 0

	for _, char := range s2 {
		if indexS1 < len(s1) && char == rune(s1[indexS1]) {
			indexS1++
		}

		if indexS1 == len(s1) {
			return 1
		}
	}

	return 0
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	result := hiddenp(s1, s2)
	fmt.Println(result)
}
