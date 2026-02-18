package main

import "fmt"

//Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	data := []int{1, 2, 3, 4, 5}
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
	res := reverseList(&list)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}

// Решение рекурсивное
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// Решение итеративное
func reverseListIterative(head *ListNode) *ListNode {
	cur := head
	var res *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = res
		res = cur
		cur = next
	}
	return res
}

// Решение в котором переворачивается только значение а ссылки в связном списке прежние
func reverseListVal(head *ListNode) *ListNode {
	var i []int
	current := head
	res := head
	for current != nil {
		i = append(i, current.Val)
		current = current.Next
	}
	for j := len(i) - 1; j >= 0; j-- {
		head.Val = i[j]
		head = head.Next
	}

	return res
}

// Решение с помощью мапы
func reverseList_Map(head *ListNode) *ListNode {
	m := make(map[int]*ListNode)
	var i int
	for ; head != nil; i++ {
		m[i] = head
		head = head.Next
	}
	i--
	if i == -1 {
		return nil
	}
	list := ListNode{
		Val:  m[i].Val,
		Next: nil,
	}
	current := &list
	for ; i-1 >= 0; i-- {
		current.Val = m[i].Val
		current.Next = m[i-1]
		current = current.Next
	}
	current.Next = nil
	return &list
}
