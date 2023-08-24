package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	runes := []rune(args[0])
	for i := 0; i < len(runes); i++ {
		if runes[0] == '-' {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}

	}
	number := RunesToInt(runes)
	result := AddPrimeSum(number)
	PrintInt(result)

}

func AddPrimeSum(num int) int {
	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum = sum + i
		}
	}

	return sum
}

func isPrime(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func RunesToInt(runes []rune) int {
	result := 0
	for i := 0; i < len(runes); i++ {
		result = result*10 + int(runes[i]-'0')
	}
	return result
}

func PrintInt(n int) {
	var result []rune

	for n > 0 {
		result = append(result, rune('0'+n%10))
		n = n / 10
	}

	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i])

	}
	z01.PrintRune('\n')

}
