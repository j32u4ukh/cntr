package cntr

import (
	"strconv"
	"sync/atomic"
)

// TreeNode 二元樹節點
type TreeNode struct {
	Key   string
	Value any
	Left  *TreeNode
	Right *TreeNode
}

// Compare 與 cmp.Compare 相同語意：a < b 回傳負數，相等回傳 0，a > b 回傳正數。
// 插入與搜尋皆使用此函式決定左／右子樹與是否命中。
type Compare func(a, b any) int

// DefaultIntCompare 沿用舊版整數二元搜尋樹行為：a、b 皆須為 int，否則 panic。
func DefaultIntCompare(a, b any) int {
	ai := a.(int)
	bi := b.(int)
	if ai < bi {
		return -1
	}
	if ai > bi {
		return 1
	}
	return 0
}

// BinaryTree 二元樹
type BinaryTree struct {
	Root    *TreeNode
	Compare Compare // 若為 nil，Insert／InsertWithKey／Search 使用 DefaultIntCompare（比較 Value）

	autoSeq int64 // Insert 自動鍵：自 1 起遞增，以十進位字串填入 TreeNode.Key
}

// Insert 插入節點（依 Compare 比較 Value；未注入 Compare 時視 Value 為 int）。
// TreeNode.Key 自動遞增（"1"、"2"、…）。
func (t *BinaryTree) Insert(value any) {
	n := atomic.AddInt64(&t.autoSeq, 1)
	t.Root = insertNode(t.Root, strconv.FormatInt(n, 10), value, t.cmp())
}

// InsertWithKey 插入節點並設定 TreeNode.Key。
func (t *BinaryTree) InsertWithKey(key string, value any) {
	t.Root = insertNode(t.Root, key, value, t.cmp())
}

func insertNode(node *TreeNode, key string, value any, cmp Compare) *TreeNode {
	if node == nil {
		return &TreeNode{
			Key:   key,
			Value: value,
		}
	}

	switch c := cmp(value, node.Value); {
	case c < 0:
		node.Left = insertNode(node.Left, key, value, cmp)
	default:
		node.Right = insertNode(node.Right, key, value, cmp)
	}

	return node
}

// Search 搜尋節點
func (t *BinaryTree) Search(value any) bool {
	return searchNode(t.Root, value, t.cmp())
}

func searchNode(node *TreeNode, value any, cmp Compare) bool {
	if node == nil {
		return false
	}

	switch c := cmp(value, node.Value); {
	case c == 0:
		return true
	case c < 0:
		return searchNode(node.Left, value, cmp)
	default:
		return searchNode(node.Right, value, cmp)
	}
}

func (t *BinaryTree) cmp() Compare {
	if t.Compare != nil {
		return t.Compare
	}
	return DefaultIntCompare
}
