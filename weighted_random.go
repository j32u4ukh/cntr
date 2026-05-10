package cntr

import (
	"sort"
)

const weightedRandomBufferSize = 64

// WeightedRandom 加權隨機選擇器，權重與累加前綴皆為 T（須可排序、可加總）。
type WeightedRandom[T Number | NumberX] struct {
	tree        *BinaryTree
	totalWeight T
	rnd         *BufferedRandom[T]
}

// NewWeightedRandom 建立選擇器；常見用法：NewWeightedRandom[int]()。
func NewWeightedRandom[T Number | NumberX]() *WeightedRandom[T] {
	s := &WeightedRandom[T]{
		tree: &BinaryTree{},
		rnd:  NewBufferedRandom[T](weightedRandomBufferSize),
	}
	s.tree.Compare = func(a, b any) int {
		return compareNum(a.(T), b.(T))
	}
	return s
}

// Init 使用字典初始化（僅納入權重 > 0；鍵依字串排序以固定累加順序）
func (s *WeightedRandom[T]) Init(weights map[string]T) {
	var pairs []IndexedItem
	z := zeroNum[T]()
	for k, w := range weights {
		if compareNum(w, z) > 0 {
			pairs = append(pairs, IndexedItem{Item: k, Value: w})
		}
	}
	if len(pairs) == 0 {
		s.totalWeight = z
		s.tree.Root = nil
		return
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Item.(string) < pairs[j].Item.(string)
	})

	keys := make([]string, len(pairs))
	prefixes := make([]T, len(pairs))
	sum := z
	for i := range pairs {
		keys[i] = pairs[i].Item.(string)
		w := pairs[i].Value.(T)
		sum = addNum(sum, w)
		prefixes[i] = sum
	}
	s.totalWeight = sum

	s.tree.Root = nil
	s.buildBalancedTree(keys, prefixes)

	s.rnd.Init(z, s.totalWeight)
}

func (s *WeightedRandom[T]) buildBalancedTree(keys []string, prefixes []T) {
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
func maxPrefixLessThan[T Number | NumberX](root *TreeNode, prefix T) *TreeNode {
	var cand *TreeNode
	for n := root; n != nil; {
		nv := n.Value.(T)
		if compareNum(nv, prefix) < 0 {
			cand = n
			n = n.Right
		} else {
			n = n.Left
		}
	}
	return cand
}

// Select 執行加權隨機抽取（目標值來自 BufferedRandom[T]，區間 [0, totalWeight)，與 Init 時一致）。
func (s *WeightedRandom[T]) Select() string {
	z := zeroNum[T]()
	if compareNum(s.totalWeight, z) <= 0 || s.tree.Root == nil {
		return ""
	}
	target := s.rnd.Next()
	curr := s.tree.Root
	for curr != nil {
		prefix := curr.Value.(T)
		lowerBound := z
		if pred := maxPrefixLessThan(s.tree.Root, prefix); pred != nil {
			lowerBound = pred.Value.(T)
		}
		if compareNum(target, lowerBound) >= 0 && compareNum(target, prefix) < 0 {
			return curr.Key
		}
		if compareNum(target, lowerBound) < 0 {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return ""
}

func zeroNum[T Number | NumberX]() T {
	var z T
	return z
}

func compareNum[T Number | NumberX](a, b T) int {
	switch aa := any(a).(type) {
	case int:
		bb := any(b).(int)
		switch {
		case aa < bb:
			return -1
		case aa > bb:
			return 1
		default:
			return 0
		}
	case int8:
		bb := any(b).(int8)
		switch {
		case aa < bb:
			return -1
		case aa > bb:
			return 1
		default:
			return 0
		}
	case int16:
		bb := any(b).(int16)
		switch {
		case aa < bb:
			return -1
		case aa > bb:
			return 1
		default:
			return 0
		}
	case int32:
		bb := any(b).(int32)
		switch {
		case aa < bb:
			return -1
		case aa > bb:
			return 1
		default:
			return 0
		}
	case int64:
		bb := any(b).(int64)
		switch {
		case aa < bb:
			return -1
		case aa > bb:
			return 1
		default:
			return 0
		}
	case uint:
		bb := any(b).(uint)
		switch {
		case aa < bb:
			return -1
		case aa > bb:
			return 1
		default:
			return 0
		}
	case uint8:
		bb := any(b).(uint8)
		switch {
		case aa < bb:
			return -1
		case aa > bb:
			return 1
		default:
			return 0
		}
	case uint16:
		bb := any(b).(uint16)
		switch {
		case aa < bb:
			return -1
		case aa > bb:
			return 1
		default:
			return 0
		}
	case uint32:
		bb := any(b).(uint32)
		switch {
		case aa < bb:
			return -1
		case aa > bb:
			return 1
		default:
			return 0
		}
	case uint64:
		bb := any(b).(uint64)
		switch {
		case aa < bb:
			return -1
		case aa > bb:
			return 1
		default:
			return 0
		}
	case float32:
		bb := any(b).(float32)
		switch {
		case aa < bb:
			return -1
		case aa > bb:
			return 1
		default:
			return 0
		}
	case float64:
		bb := any(b).(float64)
		switch {
		case aa < bb:
			return -1
		case aa > bb:
			return 1
		default:
			return 0
		}
	default:
		return 0
	}
}

func addNum[T Number | NumberX](a, b T) T {
	switch aa := any(a).(type) {
	case int:
		return any(aa + any(b).(int)).(T)
	case int8:
		return any(aa + any(b).(int8)).(T)
	case int16:
		return any(aa + any(b).(int16)).(T)
	case int32:
		return any(aa + any(b).(int32)).(T)
	case int64:
		return any(aa + any(b).(int64)).(T)
	case uint:
		return any(aa + any(b).(uint)).(T)
	case uint8:
		return any(aa + any(b).(uint8)).(T)
	case uint16:
		return any(aa + any(b).(uint16)).(T)
	case uint32:
		return any(aa + any(b).(uint32)).(T)
	case uint64:
		return any(aa + any(b).(uint64)).(T)
	case float32:
		return any(aa + any(b).(float32)).(T)
	case float64:
		return any(aa + any(b).(float64)).(T)
	default:
		return zeroNum[T]()
	}
}
