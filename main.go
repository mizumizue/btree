package main

import (
	"btree/btree"
	"fmt"
)

func main() {
	tree := btree.NewBTree()
	tree.Add(5)
	tree.Add(3)
	tree.Add(8)
	tree.Add(2)
	tree.Add(4)
	tree.Add(6)
	tree.Add(9)
	tree.Add(1)
	fmt.Println(tree.Max())
	fmt.Println(tree.Min())
}
