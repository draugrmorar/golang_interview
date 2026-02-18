package main

import (
	"fmt"
	"slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	sl := make([]int, 0)
	for i := 0; i < len(lists); i++ {
		current := lists[i]
		for current != nil {
			sl = append(sl, current.Val)
			current = current.Next
		}
	}
	slices.Sort(sl)
	if len(sl) == 0 {
		return nil
	}
	res := ListNode{Val: sl[0]}
	current := &res
	for i := 1; i < len(sl); i++ {
		current.Next = &ListNode{
			Val: sl[i],
		}
		current = current.Next
	}
	return &res
}

func main() {
	data := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	lists := make([]*ListNode, 3)
	for i := 0; i < len(data); i++ {
		lists[i] = &ListNode{
			Val:  data[i][0],
			Next: nil,
		}
		current := lists[i]
		for j := 1; j < len(data[i]); j++ {
			current.Next = &ListNode{
				Val: data[i][j],
			}
			current = current.Next
		}
	}

	res := mergeKLists(lists)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}
