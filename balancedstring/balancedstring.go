package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	fmt.Println(BalancedString(args[0]))

}

func BalancedString(str string) int {
	count := 0
	countC := 0
	countD := 0

	for i := 0; i < len(str); i++ {
		if str[i] == 'C' {
			countC++
		} else if str[i] == 'D' {
			countD++
		}

		if countC == countD {
			count++
			countC = 0
			countD = 0
		}

	}
	return count

}
