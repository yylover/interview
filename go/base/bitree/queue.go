package main


type TreeNode struct {
	Val int
	Left  *TreeNode;
	Right *TreeNode;
}

type LinkNode struct {
	Prev *LinkNode
	Next *LinkNode
	Data *TreeNode
}

type QueueList struct {
	Len int
	Head *LinkNode
	Tail *LinkNode
}

func NewQueueList() *QueueList {
	return &QueueList{
		Len:  0,
		Head: nil,
		Tail: nil,
	}
}

func (ql *QueueList) Emtpy() bool {
	return ql.Len == 0
}

func (ql *QueueList) PushBack(data *TreeNode) {
	node := &LinkNode{nil, nil, data}
	ql.Len++
	if ql.Head == nil {
		ql.Head = node
		ql.Tail = ql.Head
		return
	}

	ql.Tail.Next = node
	node.Prev = ql.Tail
	ql.Tail = node
}

func (ql *QueueList) PopBack() *TreeNode {
	if ql.Head == nil {
		return nil
	}

	node := ql.Tail
	ql.Tail = node.Prev
	ql.Tail.Next = nil
	ql.Len--
	return node.Data
}

func (ql *QueueList) Pop() *TreeNode {
	if ql.Head == nil {
		return nil
	}

	ql.Len--
	node := ql.Head
	ql.Head = node.Next
	if ql.Head != nil {
		ql.Head.Prev = nil
	}

	return node.Data
}


