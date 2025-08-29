package cntr

import (
	"sync"

	"github.com/pkg/errors"
)

type NestedMap[K1 Element, K2 Element, V Element] struct {
	dict map[K1]map[K2]V
	mu   sync.Mutex
}

func NewNestedMap[K1 Element, K2 Element, V Element]() *NestedMap[K1, K2, V] {
	nm := &NestedMap[K1, K2, V]{
		dict: make(map[K1]map[K2]V),
	}
	return nm
}

func (nm *NestedMap[K1, K2, V]) Set(k1 K1, k2 K2, v V) {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	var ok bool
	if _, ok = nm.dict[k1]; !ok {
		nm.dict[k1] = make(map[K2]V)
	}
	nm.dict[k1][k2] = v
}

func (nm *NestedMap[K1, K2, V]) Get(k1 K1, k2 K2) (V, error) {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	var v V
	var ok bool
	if _, ok = nm.dict[k1]; !ok {
		return v, errors.Errorf("Not found Key1: %+v.", k1)
	}
	if _, ok = nm.dict[k1][k2]; !ok {
		return v, errors.Errorf("Not found Key2: %+v.", k2)
	}
	return nm.dict[k1][k2], nil
}

func (nm *NestedMap[K1, K2, V]) GetByK1(k1 K1) (map[K2]V, error) {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	var value map[K2]V
	var ok bool
	if value, ok = nm.dict[k1]; !ok {
		return nil, errors.Errorf("Not found Key1: %+v.", k1)
	}
	return value, nil
}

func (nm *NestedMap[K1, K2, V]) Del(k1 K1, k2 K2) (V, error) {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	var v V
	var ok bool
	if _, ok = nm.dict[k1]; !ok {
		return v, errors.Errorf("Not found Key1: %+v.", k1)
	}
	if _, ok = nm.dict[k1][k2]; !ok {
		return v, errors.Errorf("Not found Key2: %+v.", k2)
	}
	v = nm.dict[k1][k2]
	delete(nm.dict[k1], k2)
	return v, nil
}

func (nm *NestedMap[K1, K2, V]) DelByK1(k1 K1) (map[K2]V, error) {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	var v map[K2]V
	var ok bool
	if v, ok = nm.dict[k1]; !ok {
		return v, errors.Errorf("Not found Key1: %+v.", k1)
	}
	delete(nm.dict, k1)
	return v, nil
}
