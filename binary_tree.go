package cntr

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

// BinaryTree 二元樹
type BinaryTree struct {
	Root    *TreeNode
	Compare Compare // 若為 nil，Insert／Search 使用 DefaultIntCompare
}

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

func (t *BinaryTree) cmp() Compare {
	if t.Compare != nil {
		return t.Compare
	}
	return DefaultIntCompare
}

// Insert 插入節點（依 Compare 的 BST 規則；未注入 Compare 時等同舊版 int 比較）
func (t *BinaryTree) Insert(value any) {
	t.Root = insertNode(t.Root, value, t.cmp())
}

func insertNode(node *TreeNode, value any, cmp Compare) *TreeNode {
	if node == nil {
		return &TreeNode{
			Value: value,
		}
	}

	switch c := cmp(value, node.Value); {
	case c < 0:
		node.Left = insertNode(node.Left, value, cmp)
	default:
		node.Right = insertNode(node.Right, value, cmp)
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
