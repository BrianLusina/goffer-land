// Package binarytrees contains types and struct forTrees
package binarytrees

// BinaryTreeNode represent a BinaryTreeNode in a BinarySearchTree
type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// NewBinaryTree creates a new Binary Tree Node
func (t *BinaryTreeNode) NewBinaryTree(data int) BinaryTreeNode {
	return BinaryTreeNode{
		Data: data,
	}
}

func (t *BinaryTreeNode) inorderTraversalIteravely() (result []int) {
	stack := []*BinaryTreeNode{}
	current := t

	for current != nil || len(stack) != 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		result = append(result, current.Data)
		current = current.Right
	}
	return
}

func (t *BinaryTreeNode) inorderTraversalRecurse(root *BinaryTreeNode) (result []int) {
	if root != nil {
		if root.Left != nil {
			t.inorderTraversalRecurse(t.Left)
		}

		result = append(result, root.Data)

		if root.Right != nil {
			t.inorderTraversalRecurse(t.Right)
		}
	}

	return
}

func (t *BinaryTreeNode) inOrderMorrisTraversal() (result []int) {
	current := t
	var pre *BinaryTreeNode

	for current != nil {
		if current.Left == nil {
			// add the current value of the node
			result = append(result, current.Data)
			// # Move to next right node
			current = current.Right
		} else {
			// # we have a left subtree
			pre = current.Left

			// # find rightmost
			for pre.Right != nil {
				pre = pre.Right
			}

			// # put current after the pre node
			pre.Right = current
			// # store current node
			temp := current
			// # move current to top of new tree
			current = current.Left
			// # original current left be None, avoid infinite loops
			temp.Left = nil
		}
	}

	return
}

// IsValidBst checks if a binary tree is valid
func (t *BinaryTreeNode) IsValidBst() bool {
	panic("Implement me!")
}

// PreOrderTraversal of a binary tree, returns values of each node
func (t *BinaryTreeNode) PreOrderTraversal() (values []int) {

	if t == nil {
		return
	}

	stack := []*BinaryTreeNode{}
	current := t

	for current != nil || len(stack) != 0 {
		for current != nil {
			values = append(values, t.Data)
			stack = append(stack, t)
			current = t.Left
		}
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		current = current.Right
	}

	return
}

// PostorderTraversal of a binary tree, returns values of each node starting with left most subtree, uses 2 stacks
// to keep track of values of nodes and pops them from one stack adding them to the other
func (t *BinaryTreeNode) PostOrderTraversal() (values []int) {
	stackOne := []*BinaryTreeNode{}
	stackTwo := []*BinaryTreeNode{}

	if t == nil {
		return
	}

	stackOne = append(stackOne, t)

	for len(stackOne) != 0 {
		node := stackOne[len(stackOne)-1]
		stackOne = stackOne[:len(stackOne)-1]
		stackTwo = append(stackTwo, node)

		if node.Left != nil {
			stackOne = append(stackOne, node.Left)
		}

		if node.Right != nil {
			stackOne = append(stackOne, node.Right)
		}
	}

	for len(stackTwo) != 0 {
		node := stackTwo[len(stackTwo)-1]
		stackTwo = stackTwo[:len(stackTwo)-1]
		values = append(values, node.Data)
	}

	return
}

// SearchNode searches for a value in a BST by walking either left or right of the tree given the value is either
// less than or greater than current node respectively. This uses a recursive approach to find a node in a Tree
// if found, returns the node which is the subtree with that value if not found, returns nil
func (t *BinaryTreeNode) SearchNode(node *BinaryTreeNode, val int) *BinaryTreeNode {
	if node == nil {
		return nil
	}

	if val == node.Data {
		return node
	}

	if val < node.Data && node.Left != nil {
		return t.SearchNode(node.Left, val)
	}

	if val > node.Data && node.Right != nil {
		return t.SearchNode(node.Right, val)
	}

	return nil
}

// InsertNode inserts a BinaryTreeNode into the BST. Inserts it left if the val is less than the current root
// inserts it right if the val is greater than the current root. This operation is repeated recursively
func (root *BinaryTreeNode) InsertNode(val int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{
			Data: val,
		}
	}

	if val < root.Data && root.Left != nil {
		root.InsertNode(val)
	} else if val <= root.Data {
		root.Left = &BinaryTreeNode{
			Data: val,
		}
	} else if val > root.Data && root.Right != nil {
		root.InsertNode(val)
	} else {
		root.Right = &BinaryTreeNode{
			Data: val,
		}
	}
	return root
}
