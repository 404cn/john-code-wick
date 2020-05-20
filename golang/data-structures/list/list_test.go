package list

import (
	"testing"
)


func TestNewList(t *testing.T) {
	list := NewList()
	if list.Length != 0 || list.Head != nil {
		t.Errorf("List's length: %v, node: %v\n", list.Length, list.Head)
	}
}

func TestNewNode(t *testing.T) {
	node := NewNode("test")
	if node.Value != "test" || node.Next != nil {
		t.Errorf("Node's value: %v, next: %v\n", node.Value, node.Next)
	}
}
