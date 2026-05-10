package cntr_test

import (
	"fmt"

	"github.com/j32u4ukh/cntr"
)

func ExampleBinaryTree_Insert() {
	tree := cntr.BinaryTree{}
	tree.Insert(7)
	tree.InsertWithKey("named", 3)
	fmt.Println(tree.Search(7), tree.Search(3))
	// Output:
	// true true
}

func Example_binaryTree_stringValues() {
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
	tree.InsertWithKey("z", "zebra")
	tree.InsertWithKey("a", "apple")
	fmt.Println(tree.Search("apple"))
	// Output:
	// true
}
