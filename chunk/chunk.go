package main

import "github.com/01-edu/z01"

func main() {
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 1)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 2)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 6)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 7)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 8)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, -1)

}

func Chunk(slice []int, size int) {
	if size > len(slice) {
		z01.PrintRune('[')
		z01.PrintRune(']')
		z01.PrintRune('\n')
	} else if size == 0 || size < 0 {
		z01.PrintRune('\n')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('[')
		z01.PrintRune('[')
		for i := 0; i < len(slice); i++ {
			printNbr(slice[i])
			if (i+1)%size != 0 && i != len(slice)-1 {
				z01.PrintRune(' ')
			}
			if (i+1)%size == 0 && i != len(slice)-1 {
				z01.PrintRune(']')
				z01.PrintRune(' ')
				z01.PrintRune('[')
			}
		}
		z01.PrintRune(']')
		z01.PrintRune(']')
		z01.PrintRune('\n')
	}

}

func printNbr(num int) {
	var digits []rune
	if num == 0 {
		z01.PrintRune('0')
	}

	for num > 0 {
		digits = append(digits, rune('0'+num%10))
		num /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}

}
