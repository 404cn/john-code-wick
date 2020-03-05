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
	Root *Node
}

// NewBst ...
// return bst with a node, if node is nil then return default bst
func NewBst(n *Node) *BSTs {
	if n == nil {
		return &BSTs{}
	}
	return &BSTs{Root: n, Size: 1}
}

// Insert ...
// 插值
func (bst *BSTs) Insert(value int) {
	insertNode(bst.Root, value)
	bst.Size++
}

func insertNode(node *Node, value int) {
	if node == nil {
		node.Value = value
		return
	}

	if value < node.Value {
		insertNode(node.Left, value)
	} else {
		insertNode(node.Right, value)
	}
}

// GetNodeCount ...
// 查找树上的节点数
func (bst *BSTs) GetNodeCount() int {
	return bst.Size
}

// PrintValues ...
// 从小到大打印树中的节点
func (bst *BSTs) PrintValues() {
	printValues(bst.Root)
}

func printValues(node *Node) {
	if node != nil {
		printValues(node.Left)
		fmt.Printf("%v ", node.Value)
		printValues(node.Right)
	}
}

// IsInTree ...
// 如果值存在于树中则返回true
func (bst *BSTs) IsInTree(value int) bool {
	return isInTree(bst.Root, value)
}

func isInTree(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	} else {
		return map[bool]bool{true: isInTree(node.Left, value), false: isInTree(node.Right, value)}[value < node.Value]
	}
}

// GetHeight 返回节点所在的高度(如果只有一个节点, 那么高度则为1)
func GetHeight(node *Node) int {
	if node.Left == nil && node.Right == nil {
		return 1
	}

	l := GetHeight(node.Left) + 1
	r := GetHeight(node.Right) + 1

	return map[bool]int{true: l, false: r}[l >= r]
}

// GetMin ...
// 返回树上的最小值
func GetMin(node *Node) *Node {
	curr := node

	for curr.Left != nil {
		curr = curr.Left
	}

	return curr
}

// GetMax ...
// 返回树上的最大值
func GetMax(node *Node) *Node {
	curr := node

	for curr.Right != nil {
		curr = curr.Right
	}

	return curr
}

// IsBinarySearchTree ...
// 是否为binnary search tree
func IsBinarySearchTree(root *Node) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left == nil || root.Right == nil {
		return false
	}

	if root.Left.Value >= root.Value || root.Right.Value < root.Value {
		return false
	}

	return IsBinarySearchTree(root.Left) && IsBinarySearchTree(root.Right)
}

// DeleteValue ...
// 根据值删除树上的节点
func DeleteValue(value int, root *Node) *Node {
	// TODO parent node
	if root == nil {
		return nil
	}

	if value < root.Value {
		root.Left = DeleteValue(value, root.Left)
	} else if value > root.Value {
		root.Right = DeleteValue(value, root.Right)
	} else {
		if root.Left == nil {
			root = root.Right
		} else if root.Right == nil {
			root = root.Left
		} else {
			// 用后继节点替换要删除的节点
			// 因为后继节点是比要删除的节点的值大的最小的节点
			d := GetSuccessor(root)
			root.Value = d.Value
			root.Right = DeleteValue(d.Value, root.Right)
		}
	}

	return root
}

// GetSuccessor ...
// 返回给定值的后继
func GetSuccessor(node *Node) *Node {
	if node.Right != nil {
		return GetMin(node.Right)
	}

	y := node.Parent
	for y != nil && node == y.Right {
		node, y = y, y.Parent
	}

	return y
}
