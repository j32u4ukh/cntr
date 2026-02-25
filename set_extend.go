package cntr

import (
	"sync"
)

type ExtendSet[T IHashable[T]] struct {
	Elements map[string]T
	mu       sync.RWMutex // 使用 RWMutex 提升讀取效能
}

func NewExtendSet[T IHashable[T]](elements ...T) *ExtendSet[T] {
	s := &ExtendSet[T]{Elements: make(map[string]T)}
	if len(elements) > 0 {
		s.Add(elements...)
	}
	return s
}

func (s *ExtendSet[T]) Add(elements ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, element := range elements {
		// 直接賦值即可，map 會處理覆蓋。
		// 如果你想嚴格遵守「若存在則忽略」，目前的邏輯沒問題。
		s.Elements[element.GetHash()] = element
	}
}

func (s *ExtendSet[T]) Contains(element T) bool {
	s.mu.RLock() // 讀取鎖
	defer s.mu.RUnlock()
	_, ok := s.Elements[element.GetHash()]
	return ok
}

func (s *ExtendSet[T]) ToSlice() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	elements := make([]T, 0, len(s.Elements))
	for _, e := range s.Elements {
		elements = append(elements, e)
	}
	return elements
}

func (s *ExtendSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Elements = make(map[string]T) // 高效清空
}

func (s *ExtendSet[T]) Remove(element T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	key := element.GetHash()
	if _, ok := s.Elements[key]; ok {
		delete(s.Elements, key)
		return true
	}
	return false
}
