package cntr

import (
	"math/rand"
	"sync"
	"time"
)

var randPool sync.Pool = sync.Pool{
	New: func() any {
		source := rand.NewSource(time.Now().UnixNano())
		return rand.New(source)
	},
}

func applyRand(def func(randGenerator *rand.Rand)) {
	randGenerator := randPool.Get().(*rand.Rand)
	defer randPool.Put(randGenerator)
	def(randGenerator)
}

// 返回隨機 int
var RandInt = func(n int) int {
	var value int
	applyRand(func(randGenerator *rand.Rand) {
		value = randGenerator.Intn(n)
	})
	return value
}

// RandomStrategy 定義隨機數獲取策略
type RandomStrategy interface {
	Next(totalWeight int) int
}

// --- 方案一：DirectRandom (即時隨機) ---
type DirectRandom struct{}

func (s *DirectRandom) Next(totalWeight int) int {
	return RandInt(totalWeight)
}

// --- 方案二：BufferedShuffle (緩衝洗牌隨機) ---
type BufferedShuffle struct {
	buffer []int
	cursor int
	n      int
}

func NewBufferedShuffle(n int) *BufferedShuffle {
	return &BufferedShuffle{
		buffer: make([]int, 0, n),
		cursor: 0,
		n:      n,
	}
}

func (s *BufferedShuffle) Next(totalWeight int) int {
	// 如果緩衝區為空或已用完，重新生成並洗牌
	if len(s.buffer) == 0 || s.cursor >= len(s.buffer) {
		s.refill(totalWeight)
	}

	val := s.buffer[s.cursor]
	s.cursor++
	return val
}

func (s *BufferedShuffle) refill(totalWeight int) {
	s.buffer = s.buffer[:0] // 重置列表
	for i := 0; i < s.n; i++ {
		s.buffer = append(s.buffer, RandInt(totalWeight))
	}

	// 打亂緩衝區
	applyRand(func(r *rand.Rand) {
		r.Shuffle(len(s.buffer), func(i, j int) {
			s.buffer[i], s.buffer[j] = s.buffer[j], s.buffer[i]
		})
	})
	s.cursor = 0
}
