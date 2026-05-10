package cntr

import (
	"math/rand"
	"sync"
	"time"
)

// --- 隨機基礎設施 ---

var randPool = sync.Pool{
	New: func() any {
		// 使用更現代的 NewSource
		return rand.New(rand.NewSource(time.Now().UnixNano()))
	},
}

// 快速獲取隨機數生成器，減少閉包分配
func withRand[T any](fn func(*rand.Rand) T) T {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return fn(r)
}

// RandInt 返回 [0, n) 的隨機 int（使用 var 以便測試替換實作）。
var RandInt = func(n int) int {
	if n <= 0 {
		return 0
	}
	return withRand(func(r *rand.Rand) int {
		return r.Intn(n)
	})
}

// DirectRandom 每次以 RandInt(totalWeight) 取得 [0, totalWeight) 的隨機整數。
type DirectRandom struct{}

func (s *DirectRandom) Next(totalWeight int) int {
	return RandInt(totalWeight)
}

// --- 泛型隨機函數 ---

// RandNumber 回傳區間 [min, max) 上的均勻隨機值（max<=min 時回傳 min）。
func RandNumber[T Number | NumberX](min, max T) T {
	if max <= min {
		return min
	}

	var zero T
	switch any(zero).(type) {
	case float32, float64:
		fMin := anyToFloat64(min)
		fMax := anyToFloat64(max)
		res := withRand(func(r *rand.Rand) float64 {
			return fMin + r.Float64()*(fMax-fMin)
		})
		return float64ToAny[T](res)

	default:
		iMin := anyToInt64(min)
		iMax := anyToInt64(max)
		span := iMax - iMin
		res := withRand(func(r *rand.Rand) int64 {
			return r.Int63n(span)
		})
		return int64ToAny[T](iMin + res)
	}
}

func anyToInt64[T any](v T) int64 {
	switch val := any(v).(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	default:
		return 0
	}
}

func anyToFloat64[T any](v T) float64 {
	if f, ok := any(v).(float64); ok {
		return f
	}
	if f, ok := any(v).(float32); ok {
		return float64(f)
	}
	return 0
}

func float64ToAny[T any](v float64) T {
	var zero T
	switch any(zero).(type) {
	case float32:
		return any(float32(v)).(T)
	default:
		return any(v).(T)
	}
}

func int64ToAny[T any](v int64) T {
	var zero T
	switch any(zero).(type) {
	case int:
		return any(int(v)).(T)
	case int8:
		return any(int8(v)).(T)
	case int16:
		return any(int16(v)).(T)
	case int32:
		return any(int32(v)).(T)
	case int64:
		return any(v).(T)
	case uint:
		return any(uint(v)).(T)
	case uint8:
		return any(uint8(v)).(T)
	case uint16:
		return any(uint16(v)).(T)
	case uint32:
		return any(uint32(v)).(T)
	case uint64:
		return any(uint64(v)).(T)
	default:
		return zero
	}
}

// --- 緩衝策略 ---

const (
	StrategyShuffle = iota
	StrategyRefill
)

type BufferedRandom[T Number | NumberX] struct {
	mu          sync.Mutex
	buffer      []T
	minValue    T
	maxValue    T
	cursor      int
	capacity    int
	strategy    int
	maxShuffles int
	remShuffles int
	initialized bool
}

func NewBufferedRandom[T Number | NumberX](n int) *BufferedRandom[T] {
	if n <= 0 {
		n = 100
	}
	return &BufferedRandom[T]{
		buffer:   make([]T, n),
		capacity: n,
		strategy: StrategyRefill,
	}
}

func (s *BufferedRandom[T]) Init(min, max T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.minValue, s.maxValue = min, max
	s.initialized = true
	s.fill()
}

func (s *BufferedRandom[T]) Next() T {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.initialized {
		panic("BufferedRandom: call Init before Next")
	}

	val := s.buffer[s.cursor]
	s.cursor++

	if s.cursor >= s.capacity {
		s.handleBufferEmpty()
	}
	return val
}

func (s *BufferedRandom[T]) handleBufferEmpty() {
	s.cursor = 0
	if s.strategy == StrategyShuffle {
		s.shuffle()
		return
	}

	if s.maxShuffles > 0 && s.remShuffles > 0 {
		s.shuffle()
		s.remShuffles--
	} else {
		s.fill()
		s.remShuffles = s.maxShuffles
	}
}

func (s *BufferedRandom[T]) fill() {
	s.cursor = 0
	for i := 0; i < s.capacity; i++ {
		s.buffer[i] = RandNumber(s.minValue, s.maxValue)
	}
}

func (s *BufferedRandom[T]) shuffle() {
	withRand(func(r *rand.Rand) struct{} {
		r.Shuffle(len(s.buffer), func(i, j int) {
			s.buffer[i], s.buffer[j] = s.buffer[j], s.buffer[i]
		})
		return struct{}{}
	})
}
