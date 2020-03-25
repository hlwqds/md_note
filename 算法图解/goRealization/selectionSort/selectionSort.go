package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var Lh *ListHead

type ListNode struct {
	next, prev *ListNode
	name       string
	num        int
}
type ListHead struct {
	list *ListNode
}

func InitList() *ListHead {
	return &ListHead{
		list: nil,
	}
}
func (L *ListHead) AddNode(n *ListNode) {

	if L.list != nil {
		L.list.prev = n
	}
	n.next = L.list
	L.list = n
}

func (L *ListHead) DelNode(name string) *ListNode {
	var lp *ListNode
	for lp = L.list; lp != nil; lp = lp.next {

		if strings.Compare(lp.name, name) == 0 {
			if lp.next != nil {
				lp.next.prev = lp.prev
			}
			if lp.prev != nil {
				lp.prev.next = lp.next
			} else {
				L.list = lp.next
			}
			return lp
		}
	}
	return nil
}

func (L *ListHead) Print() {
	for lp := L.list; lp != nil; lp = lp.next {
		fmt.Printf("%p, %v\n", lp, lp)
	}
}

func (L *ListHead) PopMax() *ListNode {
	var lp *ListNode
	var maxNode *ListNode = L.list
	for lp = L.list; lp != nil; lp = lp.next {
		if lp.num > maxNode.num {
			maxNode = lp
		}
	}

	if maxNode != nil {
		if maxNode.next != nil {
			maxNode.next.prev = maxNode.prev
		}
		if maxNode.prev != nil {
			maxNode.prev.next = maxNode.next
		} else {
			L.list = maxNode.next
		}
	}

	return maxNode
}

func init() {
	Lh = InitList()
	for i := 0; i < 100; i++ {
		s := fmt.Sprintf("node_%d", i)
		n := ListNode{
			name: s,
			num:  rand.Int() % 100,
		}
		Lh.AddNode(&n)
	}
	Lh.Print()
}

func selectSort(lh *ListHead) *ListHead {
	tmp := InitList()

	for {
		max := lh.PopMax()
		fmt.Println(max)
		if max != nil {
			tmp.AddNode(max)
		} else {
			break
		}
	}

	return tmp
}

func main() {
	Lh.Print()
	Lh.DelNode("node_99")
	Lh.Print()

	Lh = selectSort(Lh)
	Lh.Print()
}
