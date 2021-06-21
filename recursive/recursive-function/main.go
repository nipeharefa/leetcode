package main

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	// fmt.Println("Current Head ", head.Val)

	currHead := head

	nextHead := head.Next
	currHead.Next = nextHead.Next
	// head.Next = c

	head = nextHead
	head.Next = currHead

	// b, _ := json.Marshal(head)
	// fmt.Println(string(b))
	currHead.Next = swapPairs(currHead.Next)
	return head
}

func main() {

	// node := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 			Next: &ListNode{
	// 				Val: 4,
	// 			},
	// 		},
	// 	},
	// }

	// node := &ListNode{
	// 	Val: 1,
	// }

	head := swapPairs(nil)
	// head = swapPairs(head)

	b, _ := json.Marshal(head)
	fmt.Println(string(b))
	// swapPairs(head.Next)
	// fmt.Println(head.Next.Next.Next)
	// fmt.Println(head.Val, head.Next.Val, head.Next.Next.Val)
}
