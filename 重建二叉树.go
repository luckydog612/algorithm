package main

import "fmt"

/**
	题目：输入某个二叉树的前序遍历和中序遍历的结果，请重建该二叉树。
	假设输入的前序遍历和中序遍历的结果中都不含重复的数字。例如：输入前序遍历序列{1,2,4,7,3,5,6,8}
	和中序遍历序列{4,7,2,1,5,3,8,6}，则重建如图所示的二叉树并输出它的头节点。
					1
			2				3
		4				5		6
			7				8(6的子节点)
 */

type BinaryTreeNode struct {
	Value     int
	LeftNode  *BinaryTreeNode
	RightNode *BinaryTreeNode
}

func NewBinaryTreeNode(value int) *BinaryTreeNode {
	return &BinaryTreeNode{Value: value}
}

func CreateTreeConstruct(preorder []int, inorder []int) *BinaryTreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}
	return createTreeConstruct(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func createTreeConstruct(preorder []int, preStart int, preEnd int, inorder []int, inStart int, inEnd int) *BinaryTreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	// 前序遍历序列的第一个数字是根节点的值
	root := NewBinaryTreeNode(preorder[preStart])
	// 找到中序遍历数组中根节点的位置
	var rootIndex int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Value {
			rootIndex = i
			break
		}
	}
	// 计算左子树和右子树的节点数
	leftCount := rootIndex - inStart
	//rightCount := inEnd - rootIndex

	// 左子树递归
	root.LeftNode = createTreeConstruct(preorder, preStart+1, preStart+leftCount, inorder, inStart, rootIndex-1)
	root.RightNode = createTreeConstruct(preorder, preStart+leftCount+1, preEnd, inorder, rootIndex+1, inEnd)
	return root
}

// 前序遍历
func printPreOrder(root *BinaryTreeNode, order []int) []int {
	if root == nil {
		return nil
	}
	order = append(order, root.Value)
	if root.LeftNode != nil {
		order = printPreOrder(root.LeftNode, order)
	}
	if root.RightNode != nil {
		order = printPreOrder(root.RightNode, order)
	}
	return order
}

// 中序遍历
func printInOrder(root *BinaryTreeNode, order []int) []int {
	if root == nil {
		return nil
	}
	if root.LeftNode != nil {
		order = printInOrder(root.LeftNode, order)
	}
	order = append(order, root.Value)
	if root.RightNode != nil {
		order = printInOrder(root.RightNode, order)
	}
	return order
}

func main() {
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	tree := CreateTreeConstruct(preOrder, inOrder)
	fmt.Println("根节点", tree.Value)
	/*   验证   */
	pOrder := make([]int, 0)
	iOrder := make([]int, 0)
	// 前序遍历
	pOrder = printPreOrder(tree, pOrder)
	// 中序遍历
	iOrder = printInOrder(tree, iOrder)
	fmt.Println(pOrder)
	fmt.Println(iOrder)
}
