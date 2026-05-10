package cntr

import (
	"sort"
)

const weightedRandomBufferSize = 64

// WeightedRandom 加權隨機選擇器
type WeightedRandom struct {
	tree        *BinaryTree
	totalWeight int
	rnd         *BufferedRandom[int]
}

func NewWeightedSelector() *WeightedRandom {
	return &WeightedRandom{
		tree: &BinaryTree{},
		rnd:  NewBufferedRandom[int](weightedRandomBufferSize),
	}
}

// Init 使用字典初始化（僅納入權重 > 0；鍵依字串排序以固定累加順序）
func (s *WeightedRandom) Init(weights map[string]int) {
	var pairs []IndexedItem
	for k, w := range weights {
		if w > 0 {
			pairs = append(pairs, IndexedItem{Item: k, Value: w})
		}
	}
	if len(pairs) == 0 {
		s.totalWeight = 0
		s.tree.Root = nil
		return
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Item.(string) < pairs[j].Item.(string)
	})

	keys := make([]string, len(pairs))
	prefixes := make([]int, len(pairs))
	currentSum := 0
	for i := range pairs {
		keys[i] = pairs[i].Item.(string)
		currentSum += pairs[i].Value.(int)
		prefixes[i] = currentSum
	}
	s.totalWeight = currentSum

	s.tree.Root = nil
	s.buildBalancedTree(keys, prefixes)

	s.rnd.Init(0, s.totalWeight)
}

func (s *WeightedRandom) buildBalancedTree(keys []string, prefixes []int) {
	var recruit func(int, int)
	recruit = func(start, end int) {
		if start > end {
			return
		}
		mid := (start + end) / 2
		s.tree.InsertWithKey(keys[mid], prefixes[mid])
		recruit(start, mid-1)
		recruit(mid+1, end)
	}
	recruit(0, len(keys)-1)
}

// maxPrefixLessThan 在依累加權重（Value）建成的 BST 上，找出嚴格小於 prefix 的最大 Value 所對應節點。
func maxPrefixLessThan(root *TreeNode, prefix int) *TreeNode {
	var cand *TreeNode
	for n := root; n != nil; {
		nv := n.Value.(int)
		if nv < prefix {
			cand = n
			n = n.Right
		} else {
			n = n.Left
		}
	}
	return cand
}

// Select 執行加權隨機抽取（目標值來自 BufferedRandom[int]，區間 [0, totalWeight)，與 Init 時一致）。
func (s *WeightedRandom) Select() string {
	if s.totalWeight <= 0 || s.tree.Root == nil {
		return ""
	}
	target := s.rnd.Next()
	// --- 以下維持原本的 BST 搜尋邏輯 ---
	curr := s.tree.Root
	for curr != nil {
		prefix := curr.Value.(int)
		lowerBound := 0
		if pred := maxPrefixLessThan(s.tree.Root, prefix); pred != nil {
			lowerBound = pred.Value.(int)
		}
		if target >= lowerBound && target < prefix {
			return curr.Key
		}
		if target < lowerBound {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return ""
}
