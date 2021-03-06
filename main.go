package main

import (
	"mywork/dataStructure/tree"
	"fmt"
)

func main() {
	a := tree.NewBinTreeNode(1)
	tree1 := tree.NewBinaryTree(a)
	a.SetLChild(tree.NewBinTreeNode(2))
	a.SetRChild(tree.NewBinTreeNode(5))
	a.GetLChild().SetRChild(tree.NewBinTreeNode(3))
	a.GetLChild().GetRChild().SetLChild(tree.NewBinTreeNode(4))
	a.GetRChild().SetLChild(tree.NewBinTreeNode(6))
	a.GetRChild().SetRChild(tree.NewBinTreeNode(7))
	a.GetRChild().GetLChild().SetRChild(tree.NewBinTreeNode(9))
	a.GetRChild().GetRChild().SetRChild(tree.NewBinTreeNode(8))

	node2 := a.GetLChild()
	node9 := a.GetRChild().GetLChild().GetRChild()

	fmt.Println("结点2是叶子结点吗? ", node2.IsLeaf())
	fmt.Println("结点9是叶子结点吗? ", node9.IsLeaf())

	fmt.Println("这棵树的结点总数：", tree1.GetSize())

	l := tree1.InOrder()//中序遍历
	for e := l.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*tree.BinTreeNode)
		fmt.Printf("data:%v\n", *obj)
	}
	result := tree1.Find(6)
	fmt.Printf("结点6的父节点数据：%v\t结点6的右孩子节点数据: %v\n", result.GetParent().GetData(), result.GetRChild().GetData())
}