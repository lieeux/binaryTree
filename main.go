package main

import "fmt"

type Node struct {
	Id    int
	Name  string
	Left  *Node
	Right *Node
}

func PreOrder(node *Node) { //前序遍历：根-左-右
	if node != nil {
		fmt.Printf("%d-%s\n", node.Id, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

func InfixOrder(node *Node) { //中序遍历：左-根-右
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("%d-%s\n", node.Id, node.Name)
		InfixOrder(node.Right)
	}
}

func PostOrder(node *Node) { //后序遍历：左-右-根
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("%d-%s\n", node.Id, node.Name)
	}
}

func main() {
	root := &Node{
		Id:   1,
		Name: "穆",
	}
	l := &Node{
		Id:   2,
		Name: "阿鲁迪巴",
	}
	r := &Node{
		Id:   3,
		Name: "撒加",
	}
	l2 := &Node{
		Id:   10,
		Name: "修罗",
	}
	r2 := &Node{
		Id:   4,
		Name: "加隆",
	}
	lr := &Node{
		Id:   12,
		Name: "阿布罗狄",
	}
	root.Left = l
	root.Right = r
	r.Right = r2
	l.Left = l2
	l.Right = lr
	PostOrder(root)
}
