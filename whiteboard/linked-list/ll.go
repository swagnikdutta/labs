package linked_list

import "fmt"

type node struct {
	value int
	next  *node
}

func NewNode(value int) *node {
	node := new(node)
	node.value = value
	node.next = nil
	return node
}

func Start() {
	list := []int{10, 20, 30}

	var prevnode, head *node
	prevnode, head = nil, nil

	for _, item := range list {
		n := NewNode(item)
		if prevnode == nil {
			head = n
			prevnode = n
			continue
		}
		prevnode.next = n
		prevnode = n
	}

	start := head
	for start != nil {
		fmt.Print(start.value, ", ")
		start = start.next
	}
}
