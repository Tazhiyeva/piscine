package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var newArgs []string

	if len(args) < 2 {
		z01.PrintRune('\n')
		return
	}

	for i := 1; i < len(args); i++ {
		newArgs = append(newArgs, args[i])
	}
	newArgs = append(newArgs, args[0])

	for i := 0; i < len(newArgs); i++ {
		runes := []rune(newArgs[i])
		if len(runes) > 1 {
			for j := 0; j < len(runes); j++ {
				z01.PrintRune(runes[j])

			}
		} else {
			z01.PrintRune(runes[0])

		}

		if i != len(newArgs)-1 {
			z01.PrintRune(' ')

		}

	}
	z01.PrintRune('\n')

}
