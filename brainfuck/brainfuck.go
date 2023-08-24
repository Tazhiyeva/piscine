package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <Brainfuck code>")
		return
	}
	code := os.Args[1]
	memory := make([]byte, 2048)
	ptr := 0

	loopStack := []int{}

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '.':
			fmt.Printf("%c", memory[ptr])
		case '[':
			if memory[ptr] == 0 {
				loopDepth := 1
				for loopDepth > 0 {
					i++
					if code[i] == '[' {
						loopDepth++
					} else if code[i] == ']' {
						loopDepth--
					}
				}
			} else {
				loopStack = append(loopStack, i)
			}

		case ']':
			if memory[ptr] != 0 {
				i = loopStack[len(loopStack)-1] - 1
			} else {
				loopStack = loopStack[:len(loopStack)-1]
			}
		}

	}
	fmt.Println()
}
