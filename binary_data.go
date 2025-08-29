package cntr

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
)

type BinaryData struct {
	// 實際數據
	buffer bytes.Buffer
	// 數據位元組順序
	order binary.ByteOrder
}

func NewBinaryData() *BinaryData {
	b := &BinaryData{
		order: binary.LittleEndian,
	}
	return b
}

func LoadBinaryData(data []byte) (*BinaryData, error) {
	b := NewBinaryData()
	err := b.AddRawData(data)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load data")
	}
	return b, nil
}

func (b *BinaryData) SetOrder(order binary.ByteOrder) {
	b.order = order
}

func (b *BinaryData) GetCapacity() uint32 {
	return uint32(b.buffer.Cap())
}

func (b *BinaryData) GetLength() uint32 {
	return uint32(b.buffer.Len())
}

func (b *BinaryData) Reset() {
	b.buffer.Reset()
}

// ==================================================
// Tools
// ==================================================

func insertNumber[T NumberX](b *BinaryData, v T) error {
	data := b.GetData()
	fmt.Printf("insertNumber raw data: %+v\n", data)
	b.Reset()
	fmt.Printf("cap: %d, len: %d\n", b.GetCapacity(), b.GetLength())
	Resetdata := b.GetData()
	fmt.Printf("Resetdata raw data: %+v\n", Resetdata)
	err := binary.Write(&b.buffer, b.order, v)
	if err != nil {
		return errors.Wrapf(err, "Failed to insert number: %+v", v)
	}
	fmt.Printf("Write number | cap: %d, len: %d\n", b.GetCapacity(), b.GetLength())
	dataV := b.GetData()
	fmt.Printf("inserted v raw data: %+v\n", dataV)
	err = b.AddRawData(data)
	if err != nil {
		return errors.Wrap(err, "Failed to rewrite original data")
	}
	fmt.Printf("Write RawData | cap: %d, len: %d\n", b.GetCapacity(), b.GetLength())
	data = b.GetData()
	fmt.Printf("inserted raw data: %+v\n", data)
	return nil
}

func popNumber[T NumberX](b *BinaryData) (T, error) {
	var v T
	err := binary.Read(&b.buffer, b.order, &v)
	if err != nil {
		return v, errors.Wrap(err, "Failed to read data")
	}
	return v, nil
}
