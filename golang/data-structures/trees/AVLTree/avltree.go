package avltree

import (
	"math"
)

type (
	Type int
)

type AVLTreeNode struct {
	Value  Type
	Height int
	Left   *AVLTreeNode
	Right  *AVLTreeNode
}

func NewAVLTreeNode(value Type) *AVLTreeNode {
	return &AVLTreeNode{Value: value}
}

// Height 获取节点的高度
func Height(root *AVLTreeNode) int {
	return map[bool]int{true: 0, false: root.Height}[root == nil]
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

// LLRotation 左左失衡的情况下旋转
func LLRotation(root *AVLTreeNode) *AVLTreeNode {
	l := root.Left

	root.Left = l.Right
	l.Right = root

	root.Height = max(Height(root.Left), Height(root.Right)) + 1
	l.Height = max(Height(l.Left), root.Height) + 1

	return l
}

// RRRotation 右右失衡的情况下旋转
func RRRotation(root *AVLTreeNode) *AVLTreeNode {
	l := root.Right

	root.Right = l.Left
	l.Left = root

	root.Height = max(Height(root.Left), Height(root.Right)) + 1
	l.Height = max(root.Height, Height(l.Right)) + 1

	return l
}

// LRRotation 左右失衡的情况下旋转
func LRRotation(root *AVLTreeNode) *AVLTreeNode {
	root.Left = RRRotation(root.Left)
	return LLRotation(root)
}

// RLRotation 右左失衡的情况下旋转
func RLRotation(root *AVLTreeNode) *AVLTreeNode {
	root.Right = LLRotation(root.Right)
	return RRRotation(root)
}

// Insert 插入节点并返回根节点
func Insert(root *AVLTreeNode, value Type) *AVLTreeNode {
	if root == nil {
		root = NewAVLTreeNode(value)
	}

	if value < root.Value {
		Insert(root.Left, value)

		if Height(root.Left)-Height(root.Right) == 2 {
			if value < root.Left.Value {
				root = LLRotation(root)
			} else {
				root = LRRotation(root)
			}
		}
	} else {
		Insert(root.Right, value)

		if Height(root.Right)-Height(root.Left) == 2 {
			if value > root.Right.Value {
				root = RRRotation(root)
			} else {
				root = RLRotation(root)
			}
		}
	}

	root.Height = max(Height(root.Left), Height(root.Right)) + 1

	return root
}

// Search 查找节点
func Search(root *AVLTreeNode, value Type) *AVLTreeNode {
	if root == nil {
		return nil
	}

	if value == root.Value {
		return root
	} else if value < root.Value {
		return Search(root.Left, value)
	} else {
		return Search(root.Right, value)
	}
}

func MaxNode(root *AVLTreeNode) *AVLTreeNode {
	if root == nil {
		return nil
	}

	if root.Right == nil {
		return root
	}

	return MaxNode(root.Right)
}

func MinNode(root *AVLTreeNode) *AVLTreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil {
		return root
	}

	return MaxNode(root.Left)
}

// Delete 删除节点并返回根节点
func Delete(root *AVLTreeNode, node *AVLTreeNode) *AVLTreeNode {
	if root == nil {
		return nil
	}

	if node.Value < root.Value {
		root.Left = Delete(root.Left, node)

		if Height(root.Right)-Height(root.Left) == 2 {
			if Height(root.Right.Left) > Height(root.Right.Right) {
				root = RLRotation(root.Right)
			} else {
				root = RRRotation(root.Right)
			}
		}
	} else if node.Value > root.Value {
		root.Right = Delete(root.Right, node)

		if Height(root.Left)-Height(root.Right) == 2 {
			if Height(root.Left.Left) > Height(root.Left.Right) {
				root = LLRotation(root)
			} else {
				root = LRRotation(root)
			}
		}
	} else {
		// 要删除的节点的左右孩子都非空的情况下
		if root.Left != nil && root.Right != nil {
			if Height(root.Left) > Height(root.Right) {
                // 如果tree的左子树比右子树高；
                // 则(01)找出tree的左子树中的最大节点
                //   (02)将该最大节点的值赋值给tree。
                //   (03)删除该最大节点。
                // 这类似于用"tree的左子树中最大节点"做"tree"的替身；
                // 采用这种方式的好处是：删除"tree的左子树中最大节点"之后，AVL树仍然是平衡的。
				maxNode := MaxNode(root.Left)
				root = maxNode
				Delete(root.Left, maxNode)
            } else {
                // 如果tree的左子树不比右子树高(即它们相等，或右子树比左子树高1)
                // 则(01)找出tree的右子树中的最小节点
                //   (02)将该最小节点的值赋值给tree。
                //   (03)删除该最小节点。
                // 这类似于用"tree的右子树中最小节点"做"tree"的替身；
                // 采用这种方式的好处是：删除"tree的右子树中最小节点"之后，AVL树仍然是平衡的。
				minNode := MinNode(root.Right)
				root = minNode
				Delete(root.Right, minNode)
			}
		} else {
			if root.Left != nil {
				root = root.Left
			} else {
				root = root.Right
			}
		}

	return root
}

// DeleteWithValue 通过值删除节点并返回根节点
func DeleteWithValue(root *AVLTreeNode, value Type) *AVLTreeNode {
	return Delete(root, Search(root, value))
}
