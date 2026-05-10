package test

import (
	"reflect"
	"testing"

	"github.com/j32u4ukh/cntr"
)

func TestBinaryTree_Insert_Search_intOrder(t *testing.T) {
	tree := &cntr.BinaryTree{}
	tree.InsertWithKey("mid", 20)
	tree.Insert(10)
	tree.Insert(30)

	if !tree.Search(20) {
		t.Fatal("expected Search(20) true")
	}
	if !tree.Search(30) || !tree.Search(10) {
		t.Fatal("expected Search leaf values")
	}
	if tree.Search(99) || tree.Search(-1) {
		t.Fatal("expected missing values false")
	}
}

func TestBinaryTree_Insert_autoKeyUsesCounter(t *testing.T) {
	tree := &cntr.BinaryTree{}

	tree.Insert(5)
	tree.Insert(6)

	if tree.Root.Key != "1" || tree.Root.Right == nil || tree.Root.Right.Key != "2" {
		t.Fatalf("unexpected structure: root Key=%q right Key=%q", tree.Root.Key, tree.Root.Right.Key)
	}

	var keys []string
	var walk func(*cntr.TreeNode)
	walk = func(n *cntr.TreeNode) {
		if n == nil {
			return
		}
		walk(n.Left)
		keys = append(keys, n.Key)
		walk(n.Right)
	}
	walk(tree.Root)
	if keys[0] != "1" || keys[1] != "2" {
		t.Fatalf("want in-order keys 1 then 2, got %+v", keys)
	}
}

func TestBinaryTree_customCompare_StringValue(t *testing.T) {
	tree := cntr.BinaryTree{
		Compare: func(a, b any) int {
			sa, sb := a.(string), b.(string)
			if sa < sb {
				return -1
			}
			if sa > sb {
				return 1
			}
			return 0
		},
	}
	tree.InsertWithKey("first", "mango")
	tree.InsertWithKey("second", "lime")
	tree.InsertWithKey("third", "papaya")

	if !tree.Search("lime") || !tree.Search("mango") {
		t.Fatal("expected string values searchable")
	}
	if tree.Search("zz") {
		t.Fatal("expected missing string false")
	}
}

func TestDefaultIntCompare(t *testing.T) {
	if cntr.DefaultIntCompare(1, 2) >= 0 {
		t.Fatal("1 < 2")
	}
	if cntr.DefaultIntCompare(5, 5) != 0 {
		t.Fatal("equal")
	}
	if cntr.DefaultIntCompare(10, 3) <= 0 {
		t.Fatal("10 > 3")
	}
}

func pairInts(pairs []*cntr.KvPair[string]) []int {
	out := make([]int, 0, len(pairs))
	for _, p := range pairs {
		out = append(out, p.Value.(int))
	}
	return out
}

func TestTraverse_orders_intBST(t *testing.T) {
	tree := &cntr.BinaryTree{}
	tree.InsertWithKey("b", 2)
	tree.InsertWithKey("a", 1)
	tree.InsertWithKey("c", 3)

	preorder := pairInts(cntr.TraversePreorder(tree.Root))
	inorder := pairInts(cntr.TraverseInorder(tree.Root))
	postorder := pairInts(cntr.TraversePostorder(tree.Root))

	if !reflect.DeepEqual(preorder, []int{2, 1, 3}) {
		t.Fatalf("preorder want [2 1 3], got %v", preorder)
	}
	if !reflect.DeepEqual(inorder, []int{1, 2, 3}) {
		t.Fatalf("inorder want [1 2 3], got %v", inorder)
	}
	if !reflect.DeepEqual(postorder, []int{1, 3, 2}) {
		t.Fatalf("postorder want [1 3 2], got %v", postorder)
	}
}

func TestTraverse_nilRoot_returnsNil(t *testing.T) {
	if cntr.TraversePreorder(nil) != nil {
		t.Fatal("nil root -> nil")
	}
	if cntr.TraverseInorder(nil) != nil {
		t.Fatal("nil root -> nil")
	}
	if cntr.TraversePostorder(nil) != nil {
		t.Fatal("nil root -> nil")
	}
	if cntr.TraverseLevelOrder(nil) != nil {
		t.Fatal("nil root -> nil")
	}
}

func TestTraverseLevelOrder_order(t *testing.T) {
	//        1
	//       / \
	//      2   3
	//     /
	//    4
	root := &cntr.TreeNode{Value: 1}
	root.Left = &cntr.TreeNode{Value: 2}
	root.Right = &cntr.TreeNode{Value: 3}
	root.Left.Left = &cntr.TreeNode{Value: 4}

	got := pairInts(cntr.TraverseLevelOrder(root))
	if !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Fatalf("level-order want [1 2 3 4], got %v", got)
	}
}

func TestBinaryTree_Search_duplicateValue_goesRight(t *testing.T) {
	tree := &cntr.BinaryTree{}
	tree.InsertWithKey("first", 1)
	tree.InsertWithKey("second", 1)

	if tree.Root.Right == nil || tree.Root.Right.Key != "second" {
		t.Fatal("BST tie should hang right child")
	}
	if !tree.Search(1) {
		t.Fatal("value 1 should exist somewhere")
	}
}
