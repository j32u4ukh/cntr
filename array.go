package cntr

import "math/rand"

// 返回陣列中任一元素
func RandomElement[T Element](elements []T) T {
	length := len(elements)
	index := RandInt(length)
	return elements[index]
}

func ShuffleArray[T any](arr []T) {
	withRand(func(randGenerator *rand.Rand) struct{} {
		for i := len(arr) - 1; i > 0; i-- {
			j := randGenerator.Intn(i + 1)
			arr[i], arr[j] = arr[j], arr[i]
		}
		return struct{}{}
	})
}

func ShuffleIndexs(length int) [][]int {
	indexes := [][]int{}
	withRand(func(randGenerator *rand.Rand) struct{} {
		for i := length - 1; i > 0; i-- {
			j := randGenerator.Intn(i + 1)
			indexes = append(indexes, []int{i, j})
		}
		return struct{}{}
	})
	return indexes
}

func ModifyIndex(index int, length int) int {
	if index < 0 {
		return 0
	} else if index >= length {
		return length - 1
	} else {
		return index
	}
}

type Array[T Element] struct {
	Elements []T
}

func NewArray[T Element](elements ...T) *Array[T] {
	a := &Array[T]{Elements: elements}
	return a
}

func (a *Array[T]) Append(v T) {
	a.Elements = append(a.Elements, v)
}

func (a *Array[T]) Contains(v T) bool {
	idx := a.Find(v)
	return idx != -1
}

func (a *Array[T]) Get(index int) (T, bool) {
	if 0 <= index && index < a.Length() {
		return a.Elements[index], true
	}
	var none T
	return none, false
}

func (a *Array[T]) Set(index int, value T) {
	if 0 <= index && index < a.Length() {
		a.Elements[index] = value
	}
}

func (a *Array[T]) GetRange(startIndex int, endIndex int) []T {
	length := a.Length()
	if startIndex > endIndex {
		temp := startIndex
		startIndex = endIndex
		endIndex = temp
	}
	startIndex = ModifyIndex(startIndex, length)
	endIndex = ModifyIndex(endIndex, length)
	results := a.Elements[startIndex:endIndex]
	values := make([]T, len(results))
	copy(values, results)
	return values
}

func (a *Array[T]) SetRange(startIndex int, endIndex int, values []T) {
	length := a.Length()
	if startIndex > endIndex {
		temp := startIndex
		startIndex = endIndex
		endIndex = temp
	}
	startIndex = ModifyIndex(startIndex, length)
	endIndex = ModifyIndex(endIndex, length)
	length = endIndex - startIndex - 1
	copy(a.Elements[startIndex:endIndex], values[:length])
}

func (a *Array[T]) GetIterator() *Iterator[T] {
	elements := make([]T, a.Length())
	copy(elements, a.Elements)
	return NewIterator(elements)
}

func (a *Array[T]) Length() int {
	return len(a.Elements)
}

func (a *Array[T]) Find(v any) int {
	for i, e := range a.Elements {
		if e == v.(T) {
			return i
		}
	}
	return -1
}

func (a *Array[T]) Remove(v T) (T, bool) {
	idx := a.Find(v)
	if idx == -1 {
		var none T
		return none, false
	}
	a.Elements = append(a.Elements[:idx], a.Elements[idx+1:]...)
	return v, true
}

func (a *Array[T]) IsEquals(other *Array[T], isStrict bool) bool {
	if a.Length() != other.Length() {
		return false
	}
	var index int
	for i, element := range a.Elements {
		if other.Elements[i] != element {
			return false
		}
		index = other.Find(element)
		if index == -1 {
			return false
		}
		// 嚴格相等: 元素與順序都需相同
		if isStrict && (i != index) {
			return false
		}
	}
	return true
}

func (a *Array[T]) Clear() {
	a.Elements = []T{}
}

func (a *Array[T]) Clone() *Array[T] {
	elements := make([]T, a.Length())
	copy(elements, a.Elements)
	clone := NewArray(elements...)
	return clone
}
