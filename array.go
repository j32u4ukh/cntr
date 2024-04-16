package cntr

import "fmt"

type Int interface {
	int | int8 | int16 | int32 | int64
}

type UInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type Element interface {
	Int | UInt | Float | string
}

type Array[T Element] struct {
	Elements []T
}

func NewArray[T Element](elements ...T) *Array[T] {
	a := &Array[T]{Elements: elements}
	return a
}

func (a *Array[T]) Append(v interface{}) {
	element, ok := v.(T)
	if ok {
		a.Elements = append(a.Elements, element)
	}
}

func (a *Array[T]) Add(v any) {
	element, ok := v.(T)
	if ok {
		a.Elements = append(a.Elements, element)
	}
}

func (a *Array[T]) Contains(v interface{}) bool {
	idx := a.Find(v)
	return idx != -1
}

func (a *Array[T]) Iter() func() (T, bool) {
	fmt.Println("此函式即將棄用")
	index := 0
	return func() (val T, ok bool) {
		if index >= a.Length() {
			return
		}
		val, ok = a.Elements[index], true
		index++
		return
	}
}

func (a *Array[T]) Get(index int) (any, bool){
	if 0 <= index && index < a.Length(){
		return a.Elements[index], true
	}
	return nil, false
}

func (a *Array[T]) Set(index int, value T){
	if 0 <= index && index < a.Length(){
		a.Elements[index] = value
	}
}

func (a *Array[T]) GetRange(startIndex int, endIndex int) []T{
	length := a.Length()
	if startIndex > endIndex{
		temp := startIndex
		startIndex = endIndex
		endIndex = temp
	}
	startIndex = ModifyIndex(startIndex, length)
	endIndex = ModifyIndex(endIndex, length)
	results := a.Elements[startIndex: endIndex]
	values := make([]T, len(results))
	copy(values, results)
	return values
}

func (a *Array[T]) SetRange(startIndex int, endIndex int, values []T) {
	length := a.Length()
	if startIndex > endIndex{
		temp := startIndex
		startIndex = endIndex
		endIndex = temp
	}
	startIndex = ModifyIndex(startIndex, length)
	endIndex = ModifyIndex(endIndex, length)
	length = endIndex - startIndex - 1
	copy(a.Elements[startIndex: endIndex], values[:length])
}

func (a *Array[T]) GetIterator() *Iterator[T] {
	elements := make([]T, a.Length())
	copy(elements, a.Elements)
	return NewIterator(elements)
}

func (a *Array[T]) Length() int {
	return len(a.Elements)
}

func (a *Array[T]) Find(v interface{}) int {
	for i, e := range a.Elements {
		if e == v.(T) {
			return i
		}
	}
	return -1
}

func (a *Array[T]) Remove(v interface{}) interface{} {
	idx := a.Find(v)
	if idx == -1 {
		return nil
	}
	a.Elements = append(a.Elements[:idx], a.Elements[idx+1:]...)
	return v
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

func ModifyIndex(index int, length int)int{
	if index < 0{
		return 0
	}else if index >= length{
		return length - 1
	} else{
		return index
	}
}