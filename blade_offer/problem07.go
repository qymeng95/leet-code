package blade_offer

import (
	"fmt"
	"github.com/badgerodon/collections/queue"
)

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}
	tree := &TreeNode{preorder[0], nil, nil}

	i := 0
	for ; i < len(inorder) && inorder[i] != preorder[0]; i++ {
	}

	tree.Left = buildTree(preorder[1:i+1], inorder[0:i])

	tree.Right = buildTree(preorder[1+i:], inorder[i+1:])

	return tree
}

func LevelOrderTraversal(tree *TreeNode) {
	T := tree
	q := queue.New()
	q.Enqueue(T)
	for q.Len() != 0 {
		T = (q.Dequeue()).(*TreeNode)
		fmt.Println(T.Val)

		if T.Left != nil {
			q.Enqueue(T.Left)
		}
		if T.Right != nil {
			q.Enqueue(T.Right)
		}
	}
}

func TestBuildTree() {
	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	tree2 := buildTree([]int{1, 2}, []int{1, 2})
	fmt.Println(*tree)
	LevelOrderTraversal(tree)
	LevelOrderTraversal(tree2)
}
