package list

import (
	"errors"
)

var (
	ErrEmpty       = errors.New("List is empty")
	ErrBadIndex    = errors.New("List is out of range or index < 0")
	ErrNodeNotExit = errors.New("Node not exit")
)

// 链表(Linked list): 是一种线性表, 但是并不会按线性的顺序存储数据, 而是在每一个
// 节点里存到下一个节点的指针(Pointer).
// 特点: 节点的链接方向是单向的; 相对于数组来说, 单向链表的随机访问速度慢, 但单向链表
// 删除/添加数据的效率很高.
// Big O: 链表在插入时可以达到O(1), 但查找一个节点或者访问特定编号的节点需要O(n).

// 带头节点的双向循环链表
type List struct {
	// length of linked list.
	Length int
	// point to linked list's next node.
	Head *Node
}

// 链表节点, 数据域与指针域.
type Node struct {
	// store node's value.
	Value interface{}
	// point to prev node.
	Prev *Node
	// point to next node.
	Next *Node
}

// 创建新的链表, 长度为0, 指针域
func NewList() *List {
	return &List{}
}

// 创建新的节点
func NewNode(value interface{}) *Node {
	return &Node{Value: value}
}

// 返回链表的长度
func (l *List) Size() int {
	return l.Length
}

// 如果链表为空返回0
func (l *List) Empty() bool {
	return l.Length == 0
}

// 返回指定位置节点的数据, 从0开始
func (l *List) ValueAt(index int) (interface{}, error) {
	if l.Empty() {
		return nil, ErrEmpty
	}

	if index < 0 || index > l.Size()-1 {
		return nil, ErrBadIndex
	}

	p := l.Head
	// index count from 0 for first
	for i := 0; i <= index; i++ {
		p = p.Next
	}

	return p.Value, nil
}

// 在链表开始除添加节点
func (l *List) PushFront(value interface{}) {
	node := NewNode(value)

	if l.Empty() {
		l.Head.Next = node
		l.Head.Prev = node
		node.Prev = l.Head
		node.Next = l.Head
	} else {
		node.Next = l.Head.Next.Next
		node.Prev = l.Head
		l.Head.Next = node
		l.Head.Next.Next.Prev = node
	}

	l.Length++
}

// remove front item and return it's value
func (l *List) PopFront() (interface{}, error) {
	if l.Empty() {
		return nil, ErrEmpty
	}

	node := l.Head.Next
	l.Head.Next = node.Next
	node.Next.Prev = l.Head

	l.Length--

	return node.Value, nil
}

// adds an item at the end
func (l *List) PushBack(value interface{}) {
	node := NewNode(value)
	p := l.Head

	for i := 0; i < l.Length; i++ {
		p = p.Next
	}

	node.Next = l.Head
	node.Prev = p
	p.Next = node
	l.Head.Prev = node

	l.Length++
}

// removes end item and returns its value
func (l *List) PopBack() (interface{}, error) {
	if l.Empty() {
		return nil, ErrEmpty
	}

	p := l.Head
	for i := 0; i < l.Length; i++ {
		p = p.Next
	}

	p.Prev.Next = l.Head
	l.Head.Prev = p.Prev

	l.Length--

	return p.Value, nil
}

// get value of front item
func (l *List) Front() (interface{}, error) {
	return l.ValueAt(0)
}

// get value of end item
func (l *List) Back() (interface{}, error) {
	return l.ValueAt(l.Length - 1)
}

// insert value at index, so current item at that index is pointed to by new item at index
func (l *List) Insert(index int, value interface{}) error {
	if index < 0 || index > l.Length {
		return ErrBadIndex
	}

	switch index {
	case 0:
		l.PushFront(value)
		return nil
	case l.Length:
		l.PushBack(value)
		return nil
	}

	p := l.Head
	for i := 0; i <= index; i++ {
		p = p.Next
	}

	node := NewNode(value)
	node.Prev = p
	node.Next = p.Next

	p.Next.Prev = node
	p.Next = node

	l.Length++

	return nil
}

// removes node at given index
func (l *List) Erase(index int) error {
	if index < 0 || index > l.Length-1 {
		return ErrBadIndex
	}

	switch index {
	case 0:
		l.PopFront()
		return nil
	case l.Length - 1:
		l.PopBack()
		return nil
	}

	p := l.Head
	for i := 0; i <= index; i++ {
		p = p.Next
	}

	p.Prev.Next = p.Next
	p.Next.Prev = p.Prev

	l.Length--

	return nil
}

// returns the value of the node at nth position from the end of the list
func (l *List) ValueNFromEnd(n int) (interface{}, error) {
	return l.ValueAt(l.Length - n)
}

// reverses the list
func (l *List) Reverse() {
	// reverse head node
	node := l.Head.Next
	l.Head.Next = l.Head.Prev
	l.Head.Next = node

	for i := 0; i < l.Length-1; i++ {
		next := node.Next
		node.Next = node.Prev
		node.Prev = next
		node = node.Next
	}
}

// removes the first item in the list with this value
func (l *List) RemoveValue(value interface{}) error {
	if l.Empty() {
		return ErrEmpty
	}

	p := l.Head

	for i := 0; i < l.Length; i++ {
		p = p.Next
		if p.Value == value {
			p.Prev.Next = p.Next
			p.Next.Prev = p.Prev
			l.Length--
			return nil
		}
	}

	return ErrNodeNotExit
}
