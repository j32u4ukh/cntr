package cntr

import (
	"bytes"
	"encoding/binary"
)

const (
	// max int32 = 2^31，這裡取小於 max int32 的最大二次冪數
	LIMIT_SIZE int32 = 1073741824
)

type BinaryData struct {
	// 實際數據
	data []byte
	// 數據位元組順序
	order binary.ByteOrder
	// 容器大小(預設為 1024)
	capacity int32
	// 數據實際長度
	length int32
	// 讀寫用的索引值
	index int32
}

func NewBinaryData() *BinaryData {
	bd := &BinaryData{
		order:    binary.LittleEndian,
		capacity: 1024,
		length:   0,
		index:    0,
	}
	bd.data = make([]byte, bd.capacity)
	return bd
}

func LoadBinaryData(data []byte) *BinaryData {
	size := ceilSquare(int32(len(data)))
	bd := NewBinaryData()
	bd.SetCapacity(size)
	bd.AddRawData(data)
	bd.ResetIndex()
	return bd
}

func (b *BinaryData) SetCapacity(capacity int32) {
	b.capacity = ceilSquare(capacity)
	data := make([]byte, b.capacity)
	if b.length > 0 {
		copy(data[:b.length], b.data[:b.length])
	}
	b.data = data
}

func (b *BinaryData) SetOrder(order binary.ByteOrder) {
	b.order = order
}

func (b *BinaryData) GetCapacity() int32 {
	return b.capacity
}

func (b *BinaryData) GetLength() int32 {
	return b.length
}

func (b *BinaryData) ResetIndex() {
	b.index = 0
}

func (b *BinaryData) Clear() {
	b.index = 0
	b.length = 0
}

// ==================================================
// 加入數據
// ==================================================

func (b *BinaryData) AddRawData(v []byte) {
	addDatas(b, v)
}

func (b *BinaryData) AddBoolean(v bool) {
	if v {
		addData(b, 1)
	} else {
		addData(b, 0)
	}
}

func (b *BinaryData) AddInt8(v int8) {
	addNumber(b, v)
}

func (b *BinaryData) AddInt16(v int16) {
	addNumber(b, v)
}

func (b *BinaryData) AddInt32(v int32) {
	addNumber(b, v)
}

func (b *BinaryData) AddInt64(v int64) {
	addNumber(b, v)
}

func (b *BinaryData) AddByte(v byte) {
	addData(b, v)
}

func (b *BinaryData) AddUInt16(v uint16) {
	addNumber(b, v)
}

func (b *BinaryData) AddUInt32(v uint32) {
	addNumber(b, v)
}

func (b *BinaryData) AddUInt64(v uint64) {
	addNumber(b, v)
}

func (b *BinaryData) AddFloat32(v float32) {
	addNumber(b, v)
}

func (b *BinaryData) AddFloat64(v float64) {
	addNumber(b, v)
}

func (b *BinaryData) AddString(v string) {
	b.AddByteArray([]byte(v))
}

func (b *BinaryData) AddByteArray(v []byte) {
	length := int32(len(v))
	b.AddInt32(length)
	addDatas(b, v)
}

func (b *BinaryData) AddMapStringString(data map[string]string) {
	length := int32(len(data))
	b.AddInt32(length)
	for k, v := range data {
		b.AddString(k)
		b.AddString(v)
	}
}

func (b *BinaryData) AddMapStringByteArray(data map[string][]byte) {
	length := int32(len(data))
	b.AddInt32(length)
	for k, v := range data {
		b.AddString(k)
		b.AddByteArray(v)
	}
}

// ==================================================
// 插入數據(目前只能插在最前面)
// ==================================================

func (b *BinaryData) InsertInt32(v int32) {
	insertNumber(b, v, 4)
}

// ==================================================
// 取出數據
// type 1: GetXXX -> ... -> GetXXX -> 數據讀完(b.index == b.length) -> Clear(清空內存數據)
// type 2: GetXXX -> ... -> GetXXX -> 讀取部分數據 -> ResetIndex(重置讀寫索引值) -> GetXXX(再次從最前面開始讀取)
// ==================================================

// 取出全部的數據
func (b BinaryData) GetData() []byte {
	result := make([]byte, b.length)
	copy(result, b.data[:b.length])
	return result
}

func (b *BinaryData) PopBoolean() bool {
	boolean := b.PopByte()
	return boolean == 1
}

func (b *BinaryData) PopInt8() int8 {
	return popNumber[int8](b, 1)
}

func (b *BinaryData) PopInt16() int16 {
	return popNumber[int16](b, 2)
}

func (b *BinaryData) PopInt32() int32 {
	return popNumber[int32](b, 4)
}

func (b *BinaryData) PopInt64() int64 {
	return popNumber[int64](b, 8)
}

func (b *BinaryData) PopByte() byte {
	result := b.data[b.index]
	b.index += 1
	return result
}

func (b *BinaryData) PopUInt16() uint16 {
	return popNumber[uint16](b, 2)
}

func (b *BinaryData) PopUInt32() uint32 {
	return popNumber[uint32](b, 4)
}

func (b *BinaryData) PopUInt64() uint64 {
	return popNumber[uint64](b, 8)
}

func (b *BinaryData) PopFloat32() float32 {
	return popNumber[float32](b, 4)
}

func (b *BinaryData) PopFloat64() float64 {
	return popNumber[float64](b, 8)
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
	result := string(b.PopByteArray())
	return result
}

func (b *BinaryData) PopByteArray() []byte {
	length := b.PopInt32()
	result := make([]byte, length)
	copy(result, b.data[b.index:b.index+length])
	b.index += length
	return result
}

// ==================================================
// Tools
// ==================================================

func addNumber[T NumberX](b *BinaryData, v T) {
	bs := NumberToBytes(v, b.order)
	addDatas(b, bs)
}

func addDatas(b *BinaryData, bs []byte) {
	length := int32(len(bs))
	// 若新增數據後將超出容量
	if b.index+length >= b.capacity {
		// 更新容器大小
		b.SetCapacity(b.index + length)
	}
	// 寫入數據
	copy(b.data[b.index:b.index+length], bs)
	// 更新容器屬性
	b.index += length
	b.length += length
}

func addData(b *BinaryData, data byte) {
	if b.index == b.capacity {
		// 更新容器大小
		// fmt.Printf("[TransData] addData | 更新容器屬性 | b.index: %d, b.capacity: %d\n", b.index, b.capacity)
		b.SetCapacity(b.capacity + 1)
	}

	// 寫入數據
	b.data[b.index] = data

	// 更新容器屬性
	b.index += 1
	b.length += 1
}

// ==================================================
// 插入數據(目前只能插在最前面)
// ==================================================

func insertNumber[T NumberX](b *BinaryData, v T, bit int32) {
	// 將原始數據往後平移 bit 個 byte
	copy(b.data[bit:b.length+bit], b.data[:b.length])
	// 將數據寫入最前面的 bit 個 byte
	copy(b.data[:bit], NumberToBytes(v, b.order))
	// 更新長度
	b.length += bit
	// 新的讀寫索引值指向最後的位置
	b.index = b.length
}

func popNumber[T NumberX](b *BinaryData, bit byte) T {
	result := BytesToNumber[T](b.data[b.index:b.index+int32(bit)], b.order)
	b.index += int32(bit)
	return result
}

// 返回大於等於 value 但不大於 LIMIT_SIZE 的二次冪數
func ceilSquare(value int32) int32 {
	if value >= LIMIT_SIZE {
		return LIMIT_SIZE
	}
	temp := value - 1
	temp |= temp >> 1
	temp |= temp >> 2
	temp |= temp >> 4
	temp |= temp >> 8
	temp |= temp >> 16
	if temp < 0 {
		return 1
	} else {
		return temp + 1
	}
}

// binary.ByteOrder
// - binary.BigEndian    7 -> [0 0 0 7]
// - binary.LittleEndian 7 -> [7 0 0 0]

// ===== 轉 byte 陣列 =====
// 數字 轉 byte 陣列
func NumberToBytes[T NumberX](v T, order binary.ByteOrder) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, order, v)
	return bytesBuffer.Bytes()
}

// ===== byte 陣列轉回原始數值 =====

func BytesToNumber[T NumberX](b []byte, order binary.ByteOrder) T {
	var result T
	buffer := bytes.NewBuffer(b)
	binary.Read(buffer, order, &result)
	return result
}
