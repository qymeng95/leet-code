package offer

import "fmt"

type BinaryTree struct {
	Val   interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

var i = -1

func (node *BinaryTree) Create(data []interface{}) *BinaryTree {
	if node == nil {
		return nil
	}
	i = i + 1
	// 创建二叉树结点
	var bin *BinaryTree
	if i >= len(data) {
		return nil
	} else {
		bin = new(BinaryTree)
		bin.Val = data[i]
		bin.Left = bin.Create(data)
		bin.Right = bin.Create(data)
	}
	return bin
}

func (node *BinaryTree) New(index int, data []interface{}) *BinaryTree {
	if node == nil {
		return nil
	}
	bin := &BinaryTree{data[index], nil, nil}
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

func NewBinary() {
	node := new(BinaryTree)
	node = node.New(0, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(node)
	//node.PrevSort()
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
