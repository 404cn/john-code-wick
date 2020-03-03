package trees

import "fmt"

// Node ...
// binary search tree's node
type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

// NewNode ...
// create a new node with value
func NewNode(i int) *Node {
	return &Node{Value: i}
}

// BSTs ...
// binnary search tree
type BSTs struct {
	// bst's size
	Size int
	// bst's rooted node
	Head *Node
}

// NewBst ...
// return bst with a node, if node is nil then return default bst
func NewBst(n *Node) *BSTs {
	if n == nil {
		return &BSTs{}
	}
	return &BSTs{Head: n, Size: 1}
}

// Insert ...
// 插值
func (bst *BSTs) Insert(value int) {
	node := NewNode(value)
	if bst.GetNodeCount() == 0 {
		bst.Head = node
		bst.Size++
		return
	}

	ptr := bst.Head
	for {
		if value < ptr.Value {
			if ptr.Left == nil {
				ptr.Left = node
				node.Parent = ptr
				break
			} else {
				ptr = ptr.Left
			}
		} else {
			if ptr.Right == nil {
				ptr.Right = node
				node.Parent = ptr
				break
			} else {
				ptr = ptr.Right
			}
		}
	}

	bst.Size++
}

// GetNodeCount ...
// 查找树上的节点数
func (bst *BSTs) GetNodeCount() int {
	return bst.Size
}

// PrintValues ...
// 从小到大打印树中的节点
func (bst *BSTs) PrintValues() {
	fmt.Printf("%v ", bst.Head.Value)
}

// DeleteTree ...
// 删除树
func (bst *BSTs) DeleteTree()

// IsInTree ...
// 如果值存在于树中则返回true
func (bst *BSTs) IsInTree()

// GetHeight ...
// 返回节点所在的高度(如果只有一个节点, 那么高度则为1)
func (bst *BSTs) GetHeight()

// GetMin ...
// 返回树上的最小值
func (bst *BSTs) GetMin()

// GetMax ...
// 返回树上的最大值
func (bst *BSTs) GetMax()

// IsBinarySearchTree ...
// 是否为binnary search tree
func (bst *BSTs) IsBinarySearchTree()

// DeleteValue ...
// 根据值删除树上的节点
func (bst *BSTs) DeleteValue()

// GetSuccessor ...
// 返回给定值的后继
func (bst *BSTs) GetSuccessor()
