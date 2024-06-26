package cntr

import (
	"fmt"
	"sync"
)

type void struct{}

var null void

type Set[T Element] struct {
	Elements map[T]void
	mu       sync.Mutex
}

func NewSet[T Element](elements ...T) *Set[T] {
	s := &Set[T]{Elements: map[T]void{}}
	if len(elements) > 0 {
		for _, element := range elements {
			s.Add(element)
		}
	}
	return s
}

// 加入數據，返回是否加入成功
func (s *Set[T]) Add(element T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 數據已存在，加入失敗
	if s.Contains(element) {
		return false
	}
	// 加入成功
	s.Elements[element] = null
	return true
}

// 數據是否存在
func (s *Set[T]) Contains(element T) bool {
	_, ok := s.Elements[element]
	return ok
}

func (s *Set[T]) Length() int {
	return len(s.Elements)
}

// 刪除元素，返回是否刪除成功
func (s *Set[T]) Remove(element T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.Contains(element) {
		delete(s.Elements, element)
	}
	// 元素不存在，刪除失敗
	return false
}

func (s *Set[T]) ToSlice() []T {
	elements := []T{}
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
	for e := range s.Elements {
		delete(s.Elements, e)
	}
}

func (s *Set[T]) Clone() *Set[T] {
	elements := s.ToSlice()
	clone := NewSet(elements...)
	return clone
}


func (s *Set[T]) String() string {
	elements := s.ToSlice()
	return fmt.Sprintf("Set %+v", elements)
}