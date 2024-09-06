package cntr

type IndexedItem struct {
	Index int
	Item  any
	Value any
}

func NewIndexedItem(index int, item any) *IndexedItem {
	return &IndexedItem{
		Index: index,
		Item:  item,
	}
}

func (ii *IndexedItem) GetItem() any {
	return ii.Item
}

func (ii *IndexedItem) SetValue(value any) {
	ii.Value = value
}

func (ii *IndexedItem) GetValue() any {
	return ii.Value
}
