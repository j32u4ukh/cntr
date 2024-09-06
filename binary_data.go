package cntr

import (
	"bytes"
	"encoding/binary"
)

type BinaryData struct {
	// 實際數據
	buffer bytes.Buffer
	// 數據位元組順序
	order binary.ByteOrder
	// 讀取數據用
	reader *bytes.Reader
}

func NewBinaryData() *BinaryData {
	b := &BinaryData{
		order:  binary.LittleEndian,
		reader: nil,
	}
	return b
}

func LoadBinaryData(data []byte) *BinaryData {
	b := NewBinaryData()
	b.buffer.Write(data)
	return b
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
// 加入數據
// ==================================================

func (b *BinaryData) AddRawData(v []byte) {
	b.buffer.Write(v)
}

func (b *BinaryData) AddBoolean(v bool) {
	binary.Write(&b.buffer, b.order, v)
}

func (b *BinaryData) AddInt8(v int8) {
	binary.Write(&b.buffer, b.order, v)
}

func (b *BinaryData) AddInt16(v int16) {
	binary.Write(&b.buffer, b.order, v)
}

func (b *BinaryData) AddInt32(v int32) {
	binary.Write(&b.buffer, b.order, v)
}

func (b *BinaryData) AddInt64(v int64) {
	binary.Write(&b.buffer, b.order, v)
}

func (b *BinaryData) AddByte(v byte) {
	b.buffer.WriteByte(v)
}

func (b *BinaryData) AddUInt16(v uint16) {
	binary.Write(&b.buffer, b.order, v)
}

func (b *BinaryData) AddUInt32(v uint32) {
	binary.Write(&b.buffer, b.order, v)
}

func (b *BinaryData) AddUInt64(v uint64) {
	binary.Write(&b.buffer, b.order, v)
}

func (b *BinaryData) AddFloat32(v float32) {
	binary.Write(&b.buffer, b.order, v)
}

func (b *BinaryData) AddFloat64(v float64) {
	binary.Write(&b.buffer, b.order, v)
}

func (b *BinaryData) AddString(v string) {
	b.AddByteArray([]byte(v))
}

func (b *BinaryData) AddByteArray(v []byte) {
	length := uint32(len(v))
	b.AddUInt32(length)
	b.buffer.Write(v)
}

func (b *BinaryData) AddMapStringString(data map[string]string) {
	length := uint32(len(data))
	b.AddUInt32(length)
	for k, v := range data {
		b.AddString(k)
		b.AddString(v)
	}
}

func (b *BinaryData) AddMapStringByteArray(data map[string][]byte) {
	length := uint32(len(data))
	b.AddUInt32(length)
	for k, v := range data {
		b.AddString(k)
		b.AddByteArray(v)
	}
}

// ==================================================
// 插入數據(目前只能插在最前面)
// ==================================================

func (b *BinaryData) InsertInt32(v int32) {
	insertNumber(b, v)
}

func (b *BinaryData) InsertUInt32(v uint32) {
	insertNumber(b, v)
}

// 取出全部的數據
func (b BinaryData) GetData() []byte {
	return b.buffer.Bytes()
}

func (b *BinaryData) PopBoolean() bool {
	boolean := b.PopByte()
	return boolean == 1
}

func (b *BinaryData) PopInt8() int8 {
	return popNumber[int8](b)
}

func (b *BinaryData) PopInt16() int16 {
	return popNumber[int16](b)
}

func (b *BinaryData) PopInt32() int32 {
	return popNumber[int32](b)
}

func (b *BinaryData) PopInt64() int64 {
	return popNumber[int64](b)
}

func (b *BinaryData) PopByte() byte {
	return popNumber[uint8](b)
}

func (b *BinaryData) PopUInt16() uint16 {
	return popNumber[uint16](b)
}

func (b *BinaryData) PopUInt32() uint32 {
	return popNumber[uint32](b)
}

func (b *BinaryData) PopUInt64() uint64 {
	return popNumber[uint64](b)
}

func (b *BinaryData) PopFloat32() float32 {
	return popNumber[float32](b)
}

func (b *BinaryData) PopFloat64() float64 {
	return popNumber[float64](b)
}

func (b *BinaryData) PopMapStringString() map[string]string {
	result := map[string]string{}
	length := b.PopInt32()
	var key, value string
	for i := int32(0); i < length; i++ {
		key = b.PopString()
		value = b.PopString()
		result[key] = value
	}
	return result
}

func (b *BinaryData) PopMapStringByteArray() map[string][]byte {
	result := map[string][]byte{}
	length := b.PopInt32()
	var key string
	var value []byte
	for i := int32(0); i < length; i++ {
		key = b.PopString()
		value = b.PopByteArray()
		result[key] = value
	}
	return result
}

func (b *BinaryData) PopString() string {
	result := b.PopByteArray()
	return string(result)
}

func (b *BinaryData) PopByteArray() []byte {
	length := b.PopUInt32()
	result := make([]byte, length)
	binary.Read(&b.buffer, b.order, result)
	return result
}

// ==================================================
// Tools
// ==================================================

func insertNumber[T NumberX](b *BinaryData, v T) {
	data := b.buffer.Bytes()
	binary.Write(&b.buffer, b.order, v)
	b.AddRawData(data)
}

func popNumber[T NumberX](b *BinaryData) T {
	var v T
	binary.Read(&b.buffer, b.order, &v)
	return v
}
