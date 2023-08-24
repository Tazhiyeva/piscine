package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AtoiBase("125", "0123456789"))       // Output: 125
	fmt.Println(piscine.AtoiBase("1111101", "01"))           // Output: 125
	fmt.Println(piscine.AtoiBase("7D", "0123456789ABCDEF"))  // Output: 125
	fmt.Println(piscine.AtoiBase("uoi", "choumi"))           // Output: 125
	fmt.Println(piscine.AtoiBase("bbbbbab", "-ab"))          // Output: 0
	fmt.Println(piscine.AtoiBase("abc", "0123456789"))       // Output: 0
	fmt.Println(piscine.AtoiBase("101010", "01"))            // Output: 0
	fmt.Println(piscine.AtoiBase("1F", "0123456789"))        // Output: 0
	fmt.Println(piscine.AtoiBase("GHI", "ABCDEF"))           // Output: 0
	fmt.Println(piscine.AtoiBase("777", "0123456789ABCDEF")) // Output: 0
	fmt.Println(piscine.AtoiBase("XYZ", "0123456789"))       // Output: 0
	fmt.Println(piscine.AtoiBase("321", "012"))              // Output: 0
	fmt.Println(piscine.AtoiBase("123", "XYZ"))              // Output: 0
	fmt.Println(piscine.AtoiBase("ABC", "abc"))              // Output: 0
	fmt.Println(piscine.AtoiBase("12345", ""))

}
