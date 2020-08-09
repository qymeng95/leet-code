package data_struct

import (
	"fmt"
	"github.com/badgerodon/collections/queue"
	"github.com/badgerodon/collections/stack"
)

type BinTree struct {
	Val   int
	Left  *BinTree
	Right *BinTree
}

var i = -1

func (node *BinTree) Create(data []int) *BinTree {
	if node == nil {
		return nil
	}
	i = i + 1
	// 创建二叉树结点
	var bin *BinTree
	if i >= len(data) {
		return nil
	} else {
		bin = new(BinTree)
		bin.Val = data[i]
		bin.Left = bin.Create(data)
		bin.Right = bin.Create(data)
	}
	return bin
}

func (node *BinTree) New(index int, data []int) *BinTree {
	if node == nil {
		return nil
	}
	bin := &BinTree{data[index], nil, nil}
	// 设置完全二叉树左节点  其特征是深度 *2+1为左节点  +2为右节点
	if index < len(data) && 2*index+1 < len(data) {
		fmt.Println()
		bin.Left = bin.New(index*2+1, data)
	}
	if i < len(data) && 2*index+2 < len(data) {
		bin.Right = bin.New(index*2+2, data)
	}
	return bin

}

func PreOrderTraversal(tree *BinTree) {
	if nil != tree {
		fmt.Println(tree.Val)
		PreOrderTraversal(tree.Left)
		PreOrderTraversal(tree.Right)
	}
}

func InOrderTraversal(tree *BinTree) {
	if nil != tree {
		InOrderTraversal(tree.Left)
		fmt.Println(tree.Val)
		InOrderTraversal(tree.Right)
	}
}

func PostOrderTraversal(tree *BinTree) {
	if nil != tree {
		PostOrderTraversal(tree.Left)
		PostOrderTraversal(tree.Right)
		fmt.Println(tree.Val)
	}
}

func InOrderTraversal2(tree *BinTree) {
	T := tree
	s := stack.New()
	for T != nil || s.Len() != 0 {
		for T != nil {
			s.Push(T)
			T = T.Left
		}
		if s.Len() != 0 {
			T = (s.Pop()).(*BinTree)
			fmt.Println(T.Val)
			T = T.Right
		}
	}
}

func LevelOrderTraversal(tree *BinTree) {
	T := tree
	q := queue.New()
	q.Enqueue(T)
	for q.Len() != 0 {
		T = (q.Dequeue()).(*BinTree)
		fmt.Println(T.Val)

		if T.Left != nil {
			q.Enqueue(T.Left)
		}
		if T.Right != nil {
			q.Enqueue(T.Right)
		}
	}
}

func NewBinary() {
	node := new(BinTree)
	node = node.New(0, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(node)

	InOrderTraversal2(node)
	LevelOrderTraversal(node)
}

//
//func TestBuildTree() {
//	n := &ListNode{1, &ListNode{2, nil}}
//	res := ReversePrint(n)
//	res2 := ReversePrint2(n)
//
//	Println(res)
//	Println(res2)
//}
