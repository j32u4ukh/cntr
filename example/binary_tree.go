package main

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
	"github.com/spf13/cobra"
)

func registerBinaryTreeCommands(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "binary-tree",
		Short: "Examples for cntr.BinaryTree.",
	}
	cmd.AddCommand(
		// go run . binary-tree int-bst
		&cobra.Command{
			Use:   "int-bst",
			Short: "Insert ints (auto key + WithKey), then Search by value.",
			Run: func(*cobra.Command, []string) {
				runBinaryTreeIntBSTExample()
			},
		},
		// go run . binary-tree custom-compare
		&cobra.Command{
			Use:   "custom-compare",
			Short: "BST ordered by Value string with custom Compare.",
			Run: func(*cobra.Command, []string) {
				runBinaryTreeCustomCompareExample()
			},
		},
		// go run . binary-tree traverse
		&cobra.Command{
			Use:   "traverse",
			Short: "Manual TreeNode tree: preorder/inorder/postorder/level-order KvPair lists.",
			Run: func(*cobra.Command, []string) {
				runBinaryTreeTraverseExample()
			},
		},
	)
	root.AddCommand(cmd)
}

func runBinaryTreeTraverseExample() {
	root := &cntr.TreeNode{Key: "root", Value: 0}
	left := &cntr.TreeNode{Key: "L", Value: 1}
	right := &cntr.TreeNode{Key: "R", Value: 2}
	root.Left = left
	root.Right = right
	left.Left = &cntr.TreeNode{Key: "LL", Value: 3}
	left.Right = &cntr.TreeNode{Key: "LR", Value: 4}
	right.Left = &cntr.TreeNode{Key: "RL", Value: 5}
	right.Right = &cntr.TreeNode{Key: "RR", Value: 6}

	printPairs := func(label string, pairs []*cntr.KvPair[string]) {
		fmt.Printf("%s:", label)
		for _, p := range pairs {
			fmt.Printf(" (%s,%v)", p.Key, p.Value)
		}
		fmt.Println()
	}
	printPairs("preorder", cntr.TraversePreorder(root))
	printPairs("inorder", cntr.TraverseInorder(root))
	printPairs("postorder", cntr.TraversePostorder(root))
	printPairs("level-order", cntr.TraverseLevelOrder(root))
}

func runBinaryTreeIntBSTExample() {
	t := &cntr.BinaryTree{}

	t.InsertWithKey("rootish", 20)
	t.Insert(10)
	t.Insert(30)

	fmt.Printf("Search(30): %v\n", t.Search(30))
	fmt.Printf("Search(15): %v\n", t.Search(15))
	fmt.Printf("Root Key=%q Value=%v\n", t.Root.Key, t.Root.Value)
}

func runBinaryTreeCustomCompareExample() {
	tree := cntr.BinaryTree{
		Compare: func(a, b any) int {
			sa := a.(string)
			sb := b.(string)
			if sa < sb {
				return -1
			}
			if sa > sb {
				return 1
			}
			return 0
		},
	}
	tree.InsertWithKey("k-banana", "banana")
	tree.InsertWithKey("k-apple", "apple")
	tree.InsertWithKey("k-pear", "pear")

	fmt.Println("in-order by Value (string):")
	var walk func(*cntr.TreeNode)
	walk = func(n *cntr.TreeNode) {
		if n == nil {
			return
		}
		walk(n.Left)
		fmt.Printf("  Key=%s Value=%s\n", n.Key, n.Value.(string))
		walk(n.Right)
	}
	walk(tree.Root)

	fmt.Printf("Search(\"pear\"): %v\n", tree.Search("pear"))
}
