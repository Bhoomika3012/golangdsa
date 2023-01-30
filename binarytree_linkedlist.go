package main

import "fmt"

func main() {
	fmt.Println("Binary Tree")
	T := BinaryTree{}
	left := BinaryTree{}
	right := BinaryTree{}

	left.makeTree(20, nil, nil)
	right.makeTree(30, nil, nil)
	T.makeTree(10, &left, &right)

	T.preorder(&T)
	fmt.Println("")
	T.inorder(&T)
	fmt.Println("")
	T.postorder(&T)

}

type TreeNode struct {
	element int
	left    *BinaryTree
	right   *BinaryTree
}

type BinaryTree struct {
	root *TreeNode
}

func (B *BinaryTree) makeTree(ele int, left *BinaryTree, right *BinaryTree) {
	B.root = &TreeNode{ele, left, right}
}

func (B BinaryTree) preorder(troot *BinaryTree) {
	if troot != nil {
		fmt.Print(troot.root.element, " ")
		B.preorder(troot.root.left)
		B.preorder(troot.root.right)
	}
}

func (B BinaryTree) inorder(troot *BinaryTree) {
	if troot != nil {
		B.preorder(troot.root.left)
		fmt.Print(troot.root.element, " ")
		B.preorder(troot.root.right)
	}
}

func (B BinaryTree) postorder(troot *BinaryTree) {
	if troot != nil {
		B.preorder(troot.root.left)
		B.preorder(troot.root.right)
		fmt.Print(troot.root.element, " ")
	}
}
