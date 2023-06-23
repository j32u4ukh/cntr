package cntr

type Iterator[V any] struct {
	datas  []V
	index  int
	length int
}

func NewIterator[V any](datas []V) *Iterator[V] {
	it := &Iterator[V]{
		datas:  datas,
		index:  -1,
		length: len(datas),
	}
	return it
}

func (it *Iterator[V]) HasNext() bool {
	it.index += 1
	return it.index < it.length
}

func (it *Iterator[V]) Next() V {
	return it.datas[it.index]
}
