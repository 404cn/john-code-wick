package rbtree

type (
	color bool
	Type  int
)

const (
	BLACK color = true
	RED   color = false
)

type RBTreeNode struct {
	Parent *RBTreeNode
	Left   *RBTreeNode
	Right  *RBTreeNode
	// zero value is false/red
	Color color
	Val   Type
}

// Insert insert new node to red black tree
func (root *RBTreeNode) Insert(node *RBTreeNode) *RBTreeNode {
	insertRecurse(root, node)

	// Repair the tree in case any of the red-black properties have been violated
	insertRepairTree(node)

	// find new root node to return
	root = node
	for root.GetParent() != nil {
		root = root.GetParent()
	}
	return root
}

func insertRecurse(root, node *RBTreeNode) {
	// Recursively descend the tree until a leaf is found.
	if root != nil {
		if node.Val < root.Val {
			if root.Left != nil {
				insertRecurse(root.Left, node)
				// FIXME 学到一手
				return
			} else {
				root.Left = node
			}
		} else {
			if root.Right != nil {
				insertRecurse(root.Right, node)
				return
			} else {
				root.Right = node
			}
		}
	}
}

// TODO
func insertRepairTree(node *RBTreeNode) {
	if node.GetParent() == nil {
		insertCase1(node)
	} else if node.GetParent().Color == BLACK {
		insertCase2(node)
	} else if node.GetUncle() != nil && node.GetUncle().Color == RED {
		insertCase3(node)
	} else {
		insertCase4(node)
	}
}

func insertCase1(node *RBTreeNode) {
	// Case 1: The current node N is at the root of the tree.
	// In this case, it is repainted black to satisfy property 2 (the root is black).
	// Since this adds one black node to every path at once,
	// property 5 (all paths from any given node to its leaf nodes contain the same number of black nodes)
	// is not violated.
	node.Color = BLACK
}

func insertCase2(node *RBTreeNode) {
	// The current node's parent P is black, so property 4 (both children of every
	// red node are black) is not invalidated. In this case, the tree is still valid.
	// Property 5 (all paths from any given node to its leaf nodes contain the same
	// number of black nodes) is not threatened, because the current node N has two
	// black leaf children, but because N is red, the paths through each of its
	// children have the same number of black nodes as the path through the leaf it
	// replaced, which was black, and so this property remains satisfied.
	return
	// because node.Color zero value is red/false.
}

func insertCase3(node *RBTreeNode) {
	// If both the parent P and the uncle U are red, then both of them can be repainted
	// black and the grandparent G becomes red to maintain property 5 (all paths from
	// any given node to its leaf nodes contain the same number of black nodes).
	// Since any path through the parent or uncle must pass through the grandparent,
	// the number of black nodes on these paths has not changed. However,
	// the grandparent G may now violate Property 2 (The root is black) if it is the root
	// or Property 4 (Both children of every red node are black) if it has a red parent.
	// To fix this, the tree's red-black repair procedure is rerun on G.

	// Note that this is a tail-recursive call, so it could be rewritten as a loop.
	// Since this is the only loop, and any rotations occur after this loop,
	// this proves that a constant number of rotations occur.
	node.GetParent().Color = BLACK
	node.GetUncle().Color = BLACK
	node.GetGrandParent().Color = RED
	// if node's grand parent node is not root node, then must check and repair nodes's grand parent node.
	insertRepairTree(node.GetGrandParent())
}

func insertCase4(node *RBTreeNode) {
	// step 1: The parent P is red but the uncle U is black. The ultimate goal will be
	// to rotate the current node into the grandparent position, but this will not work
	// if the current node is on the "inside" of the subtree under G (i.e., if N is the
	// left child of the right child of the grandparent or the right child of the left
	// child of the grandparent). In this case, a left rotation on P that switches the
	// roles of the current node N and its parent P can be performed. The rotation
	// causes some paths (those in the sub-tree labelled "1") to pass through the node
	// N where they did not before. It also causes some paths (those in the sub-tree
	// labelled "3") not to pass through the node P where they did before. However, both
	// of these nodes are red, so property 5 (all paths from any given node to its leaf
	// nodes contain the same number of black nodes) is not violated by the rotation.
	// After this step has been completed, property 4 (both children of every red node
	// are black) is still violated, but now we can resolve this by continuing to step 2.
	p, g := node.GetParent(), node.GetGrandParent()

	if node == p.Right && p == g.Left {
		rotateLeft(p)
		// TODO ?
		node = node.Left
	} else if node == p.Left && p == g.Right {
		rotateRignt(p)
		node = node.Right
	}

	insertCase4Step2(node)
}

func insertCase4Step2(node *RBTreeNode) {

}

func rotateLeft(node *RBTreeNode)
func rotateRignt(node *RBTreeNode)

// GetParent return current node's parent, root's parent node is nil
func (node *RBTreeNode) GetParent() *RBTreeNode {
	return node.Parent
}

// GetGrandParent return current nodes's grand parent, if this node
// is root or root's child, then return nil
func (node *RBTreeNode) GetGrandParent() *RBTreeNode {
	return node.GetParent().GetParent()
}

// GetSibling return current nodes's sibling, no parent means no sibling
func (node *RBTreeNode) GetSibling() *RBTreeNode {
	if node.Parent == nil {
		return nil
	}

	// TODO 相等的节点?
	if node == node.Parent.Left {
		return node.Parent.Right
	} else {
		return node.Parent.Left
	}
}

// GetUncle return current node's uncle node, no parent means no sibling
func (node *RBTreeNode) GetUncle() *RBTreeNode {
	return node.GetParent().GetSibling()
}
