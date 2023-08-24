package main

import (
	"fmt"
)

func main() {
	// 3 -> 1 -> 5
	num1 := &NodeAddL{Num: 5}
	num1 = pushFront(num1, 1)
	num1 = pushFront(num1, 3)

	// 5 -> 9 -> 2
	num2 := &NodeAddL{Num: 2}
	num2 = pushFront(num2, 9)
	num2 = pushFront(num2, 5)

	// 9 -> 0 -> 7
	result := AddLinkedNumbers(num1, num2)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func pushFront(node *NodeAddL, num int) *NodeAddL {
	newNode := &NodeAddL{Num: num, Next: node}
	return newNode
}
func AddLinkedNumbers(num1, num2 *NodeAddL) *NodeAddL {
	sum1 := 0
	sum2 := 0

	for num1 != nil || num2 != nil {
		if num1 != nil {
			sum1 = sum1*10 + num1.Num
			num1 = num1.Next
		}
		if num2 != nil {
			sum2 = sum2*10 + num2.Num
			num2 = num2.Next
		}
	}

	result := sum1 + sum2

	var newList *NodeAddL = nil
	for result > 0 {
		newList = pushFront(newList, result%10)
		result = result / 10
	}

	return newList

}
