package cntr

import (
	"fmt"
	"sync"
)

type Void struct{}

var NULL Void

// 統一的 Set 結構，支援兩種模式：
// 1. 基本類型（Element）：直接作為 map key
// 2. 複雜類型（IHashable）：使用 hash 作為 key
type Set[T Element] struct {
	Elements map[T]Void
	mu       sync.RWMutex // 統一使用 RWMutex 提升讀取效能
}

// 為基本類型創建 Set
func NewSet[T Element](elements ...T) *Set[T] {
	s := &Set[T]{Elements: make(map[T]Void)}
	if len(elements) > 0 {
		s.Add(elements...)
	}
	return s
}

// 為 IHashable 類型創建 Set（使用 hash 作為 key）
type HashableSet[T IHashable[T]] struct {
	Elements map[string]T
	mu       sync.RWMutex
}

func NewHashableSet[T IHashable[T]](elements ...T) *HashableSet[T] {
	s := &HashableSet[T]{Elements: make(map[string]T)}
	if len(elements) > 0 {
		s.Add(elements...)
	}
	return s
}

// Set 的方法
func (s *Set[T]) Add(elements ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, element := range elements {
		if !s.containsUnsafe(element) {
			s.Elements[element] = NULL
		}
	}
}

func (s *Set[T]) Contains(element T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.containsUnsafe(element)
}

func (s *Set[T]) containsUnsafe(element T) bool {
	_, ok := s.Elements[element]
	return ok
}

func (s *Set[T]) Length() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.Elements)
}

func (s *Set[T]) Remove(element T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.containsUnsafe(element) {
		delete(s.Elements, element)
		return true
	}
	return false
}

func (s *Set[T]) ToSlice() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	elements := make([]T, 0, len(s.Elements))
	for e := range s.Elements {
		elements = append(elements, e)
	}
	return elements
}

func (s *Set[T]) GetIterator() *Iterator[T] {
	elements := s.ToSlice()
	return NewIterator(elements)
}

func (s *Set[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Elements = make(map[T]Void)
}

func (s *Set[T]) Clone() *Set[T] {
	elements := s.ToSlice()
	return NewSet[T](elements...)
}

func (s *Set[T]) String() string {
	elements := s.ToSlice()
	return fmt.Sprintf("Set %+v", elements)
}

// HashableSet 的方法
func (s *HashableSet[T]) Add(elements ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, element := range elements {
		s.Elements[element.GetHash()] = element
	}
}

func (s *HashableSet[T]) Contains(element T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.Elements[element.GetHash()]
	return ok
}

func (s *HashableSet[T]) Length() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.Elements)
}

func (s *HashableSet[T]) Remove(element T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	key := element.GetHash()
	if _, ok := s.Elements[key]; ok {
		delete(s.Elements, key)
		return true
	}
	return false
}

func (s *HashableSet[T]) ToSlice() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	elements := make([]T, 0, len(s.Elements))
	for _, e := range s.Elements {
		elements = append(elements, e)
	}
	return elements
}

func (s *HashableSet[T]) GetIterator() *Iterator[T] {
	elements := s.ToSlice()
	return NewIterator(elements)
}

func (s *HashableSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Elements = make(map[string]T)
}

func (s *HashableSet[T]) Clone() *HashableSet[T] {
	elements := s.ToSlice()
	return NewHashableSet(elements...)
}

func (s *HashableSet[T]) String() string {
	elements := s.ToSlice()
	return fmt.Sprintf("HashableSet %+v", elements)
}
