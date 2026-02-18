package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	data := []int{3, 2, 0, -4}
	list := ListNode{
		Val:  data[0],
		Next: nil,
	}
	current := &list
	for i := 1; i < len(data); i++ {
		current.Next = &ListNode{
			Val: data[i],
		}
		current = current.Next
	}
	fmt.Println(hasCycle(&list))
}

func hasCycle2(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	for head != nil && head.Next != nil {
		if _, exist := m[head]; !exist {
			m[head] = true
		} else {
			return true
		}
		head = head.Next
	}
	return false
}

func hasCycle(head *ListNode) bool {
	var rabbit = head
	for head != nil && head.Next != nil && rabbit != nil && rabbit.Next.Next != nil {
		head = head.Next
		rabbit = rabbit.Next.Next
		if head == rabbit {
			return true
		}
	}
	return false
}
