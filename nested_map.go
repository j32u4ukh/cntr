package cntr

import (
	"fmt"
	"sync"
)

type NestedMap[K1 Element, K2 Element, V Element] struct {
	dict map[K1]map[K2][]V
	mu   sync.Mutex
}

func NewNestedMap[K1 Element, K2 Element, V Element]() *NestedMap[K1, K2, V] {
	nm := &NestedMap[K1, K2, V]{
		dict: make(map[K1]map[K2][]V),
	}
	return nm
}

func (nm *NestedMap[K1, K2, V]) Add(k1 K1, k2 K2, vs []V) {
	var ok bool
	if _, ok = nm.dict[k1]; !ok {
		nm.dict[k1] = make(map[K2][]V)
	}
	if _, ok = nm.dict[k1][k2]; !ok {
		nm.dict[k1][k2] = []V{}
	}
	nm.dict[k1][k2] = append(nm.dict[k1][k2], vs...)
}

func (nm *NestedMap[K1, K2, V]) Get(k1 K1, k2 K2) ([]V, error) {
	var ok bool
	if _, ok = nm.dict[k1]; !ok {
		return nil, fmt.Errorf("Not found Key1: %+v.", k1)
	}
	if _, ok = nm.dict[k1][k2]; !ok {
		return nil, fmt.Errorf("Not found Key2: %+v.", k2)
	}
	return nm.dict[k1][k2], nil
}

func (nm *NestedMap[K1, K2, V]) GetByKey1(k1 K1) ([]V, error) {
	var dict map[K2][]V
	var ok bool
	result := []V{}
	if dict, ok = nm.dict[k1]; ok {
		for _, values := range dict {
			result = append(result, values...)
		}
	}
	return result, nil
}

func (nm *NestedMap[K1, K2, V]) GetByKey2(k2 K2) ([]V, error) {
	var dict map[K2][]V
	var values []V
	var ok bool
	result := []V{}
	for _, dict = range nm.dict {
		if values, ok = dict[k2]; ok {
			result = append(result, values...)
		}
	}
	return result, nil
}
