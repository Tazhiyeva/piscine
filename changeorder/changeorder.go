package main

import "fmt"

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func Changeorder(node *NodeAddL) *NodeAddL {
	if node == nil || node.Next == nil || node.Next.Next == nil {
		return nil
	}

	oddHead := node
	evenHead := node.Next
	oddCurrent := oddHead
	evenCurrent := evenHead

	current := evenHead.Next
	odd := true

	for current != nil {
		if odd {
			oddCurrent.Next = current
			oddCurrent = oddCurrent.Next
		} else {
			evenCurrent.Next = current
			evenCurrent = evenCurrent.Next
		}

		current = current.Next
		odd = !odd
	}
	oddCurrent.Next = evenHead
	evenCurrent.Next = nil

	return oddHead
}

func main() {
	num1 := &NodeAddL{Num: 1}
	num1 = pushBack(num1, 2)
	num1 = pushBack(num1, 3)
	num1 = pushBack(num1, 4)
	num1 = pushBack(num1, 5)

	result := Changeorder(num1)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func pushBack(node *NodeAddL, num int) *NodeAddL {
	newNode := &NodeAddL{Num: num}
	if node == nil {
		return newNode
	}

	current := node
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return node
}
