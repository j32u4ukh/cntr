package cntr

import (
	"fmt"
	"sync"
)

type Bivalue[K1 Element, K2 Element, V any] struct {
	key1  K1
	key2  K2
	value V
}

func NewBivalue[K1 Element, K2 Element, V any](key1 K1, key2 K2, value V) *Bivalue[K1, K2, V] {
	v := &Bivalue[K1, K2, V]{
		key1:  key1,
		key2:  key2,
		value: value,
	}
	return v
}

type BikeyMap[K1 Element, K2 Element, V any] struct {
	dict1 map[K1]int64
	dict2 map[K2]int64
	dict3 map[int64]*Bivalue[K1, K2, V]
	mu    sync.Mutex
	index int64
}

func NewBikeyMap[K1 Element, K2 Element, V any]() *BikeyMap[K1, K2, V] {
	bm := &BikeyMap[K1, K2, V]{
		dict1: make(map[K1]int64),
		dict2: make(map[K2]int64),
		dict3: make(map[int64]*Bivalue[K1, K2, V]),
		index: 0,
	}
	return bm
}

func (m *BikeyMap[K1, K2, V]) Add(key1 K1, key2 K2, value V) error {
	var ok bool
	if _, ok = m.dict1[key1]; ok {
		return fmt.Errorf("Key1 %+v has been exists.", key1)
	}
	if _, ok = m.dict2[key2]; ok {
		return fmt.Errorf("Key2 %+v has been exists.", key2)
	}
	m.Set(key1, key2, value)
	return nil
}

func (m *BikeyMap[K1, K2, V]) Set(key1 K1, key2 K2, value V) {
	m.dict1[key1] = m.index
	m.dict2[key2] = m.index
	m.dict3[m.index] = NewBivalue[K1, K2, V](key1, key2, value)
	m.index++
}

// 根據兩個對應的 Key 來取值，也可確定兩個 Key 是一組的。
func (m *BikeyMap[K1, K2, V]) GetByKeys(key1 K1, key2 K2) (V, bool) {
	var index1, index2 int64
	var ok bool
	var v V
	if index1, ok = m.dict1[key1]; ok {
		if index2, ok = m.dict2[key2]; ok && (index1 == index2) {
			if bv, ok := m.dict3[index1]; ok {
				return bv.value, true
			}
		}
	}
	return v, false
}

func (m *BikeyMap[K1, K2, V]) GetByKey1(key1 K1) (V, bool) {
	if index, ok := m.dict1[key1]; ok {
		if bv, ok := m.dict3[index]; ok {
			return bv.value, true
		}
	}
	var v V
	return v, false
}

func (m *BikeyMap[K1, K2, V]) GetByKey2(key2 K2) (V, bool) {
	if index, ok := m.dict2[key2]; ok {
		if bv, ok := m.dict3[index]; ok {
			return bv.value, true
		}
	}
	var v V
	return v, false
}

func (m *BikeyMap[K1, K2, V]) ContainKeys(key1 K1, key2 K2) bool {
	ok := m.ContainKey1(key1)
	if !ok {
		return false
	}
	return m.ContainKey2(key2)
}

func (m *BikeyMap[K1, K2, V]) ContainKey1(key1 K1) bool {
	if _, ok := m.dict1[key1]; ok {
		return true
	}
	return false
}

func (m *BikeyMap[K1, K2, V]) ContainKey2(key2 K2) bool {
	if _, ok := m.dict2[key2]; ok {
		return true
	}
	return false
}

func (m *BikeyMap[K1, K2, V]) UpdateByKey1(key1 K1, value *Bivalue[K1, K2, V]) {
	if index, ok := m.dict1[key1]; ok {
		if origin, ok := m.dict3[index]; ok {
			m.update(index, origin, value)
		}
	}
}

func (m *BikeyMap[K1, K2, V]) UpdateByKey2(key2 K2, value *Bivalue[K1, K2, V]) {
	if index, ok := m.dict2[key2]; ok {
		if origin, ok := m.dict3[index]; ok {
			m.update(index, origin, value)
		}
	}
}

func (m *BikeyMap[K1, K2, V]) update(index int64, origin *Bivalue[K1, K2, V], value *Bivalue[K1, K2, V]) {
	if value.key1 != origin.key1 {
		delete(m.dict1, origin.key1)
		m.dict1[value.key1] = index
	}
	if value.key2 != origin.key2 {
		delete(m.dict2, origin.key2)
		m.dict2[value.key2] = index
	}
	m.dict3[index] = value
}

func (m *BikeyMap[K1, K2, V]) DelByKey1(key1 K1) {
	if index, ok := m.dict1[key1]; ok {
		if bv, ok := m.dict3[index]; ok {
			m.del(index, bv)
		}
	}
}

func (m *BikeyMap[K1, K2, V]) DelByKey2(key2 K2) {
	if index, ok := m.dict2[key2]; ok {
		if bv, ok := m.dict3[index]; ok {
			m.del(index, bv)
		}
	}
}

func (m *BikeyMap[K1, K2, V]) del(index int64, bv *Bivalue[K1, K2, V]) {
	delete(m.dict1, bv.key1)
	delete(m.dict2, bv.key2)
	delete(m.dict3, index)
}

func (m *BikeyMap[K1, K2, V]) GetIterator() *Iterator[*Bivalue[K1, K2, V]] {
	values := []*Bivalue[K1, K2, V]{}
	for _, bv := range m.dict3 {
		values = append(values, bv)
	}
	return NewIterator(values)
}
