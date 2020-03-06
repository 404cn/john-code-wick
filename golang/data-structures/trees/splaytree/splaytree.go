package splaytree

import (
	"fmt"
)

type Type int
type SplayTreeNode struct {
	Val    Type
	Left   *SplayTreeNode
	Right  *SplayTreeNode
	Parent *SplayTreeNode
}

func NewNode(val Type) *SplayTreeNode {
	return &SplayTreeNode{Val: val}
}

// Splay 旋转指定值的节点为根节点
func Splay(root, x *SplayTreeNode) {
	if root == nil {
		return
	}

	p := x.Parent
	g := p.Parent

	if x.Parent == root {
		if root.Left == x {
			zig(root)
		} else if root.Right == x {
			zag(root)
		}
	} else if p == g.Left && x == p.Left {
		zig(g)
		zig(p)
	} else if p == g.Right && x == p.Right {
		zag(g)
		zag(p)
	} else if p == g.Left && x == p.Right {
		zig(p)
		zag(g)
	} else if p == g.Right && x == g.Left {
		zag(p)
		zig(g)
	}
}

// 右旋
func zig(root *SplayTreeNode) {
	if root == nil {
		return
	}

	rLeft := root.Left

	if rLeft != nil {
		root.Left = rLeft.Right

		if rLeft.Right != nil {
			rLeft.Right.Parent = root
		}

		rLeft.Parent = root.Parent
	}

	// 如果传入的节点还有父节点
	// 则需要调整父节点的指向
	if root.Parent != nil {
		if root == root.Parent.Left {
			root.Parent.Left = rLeft
		} else {
			root.Parent.Right = rLeft
		}
	}

	if rLeft != nil {
		rLeft.Right = root
	}
	root.Parent = rLeft.Right
}

// 左旋转
func zag(root *SplayTreeNode) {
	if root == nil {
		return
	}

	rRight := root.Right

	if rRight != nil {
		root.Right = rRight.Left

		if rRight.Left != nil {
			rRight.Left.Parent = root
		}

		rRight.Parent = root.Parent
	}

	if root.Parent != nil {
		if root == root.Parent.Left {
			root.Parent.Left = rRight
		} else {
			root.Parent.Right = rRight
		}
	}

	if rRight != nil {
		rRight.Left = root
	}
	root.Parent = rRight.Left
}

// Preorder 前序遍历
func Preorder(root *SplayTreeNode) {
	if root != nil {
		fmt.Println(root.Val)
		Preorder(root.Left)
		Preorder(root.Right)
	}
}

// Inorder 中序遍历
func Inorder(root *SplayTreeNode) {
	if root != nil {
		Preorder(root.Left)
		fmt.Println(root.Val)
		Preorder(root.Right)
	}
}

// Postorder 后序遍历
func Postorder(root *SplayTreeNode) {
	if root != nil {
		Preorder(root.Left)
		Preorder(root.Right)
		fmt.Println(root.Val)
	}
}

// Search 查找节点
func Search(root *SplayTreeNode, val Type) *SplayTreeNode {
	if root == nil {
		return nil
	}

	if val == root.Val {
		return root
	} else if val < root.Val {
		return Search(root.Left, val)
	} else {
		return Search(root.Right, val)
	}
}

// MaxNode 返回最大节点
func MaxNode(root *SplayTreeNode) *SplayTreeNode {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root
	}
	return MaxNode(root.Right)
}

// MinNode 返回最小节点
func MinNode(root *SplayTreeNode) *SplayTreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil {
		return root
	}
	return MinNode(root.Left)
}
