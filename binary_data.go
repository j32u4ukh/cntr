package cntr

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"

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

func writeVarint(b *BinaryData, v int64) error {
	var buf [binary.MaxVarintLen64]byte
	n := binary.PutVarint(buf[:], v)
	_, err := b.buffer.Write(buf[:n])
	return err
}

func writeUvarint(b *BinaryData, v uint64) error {
	var buf [binary.MaxVarintLen64]byte
	n := binary.PutUvarint(buf[:], v)
	_, err := b.buffer.Write(buf[:n])
	return err
}

func writeFloat32(b *BinaryData, v float32) error {
	var buf [4]byte
	b.order.PutUint32(buf[:], math.Float32bits(v))
	_, err := b.buffer.Write(buf[:])
	return err
}

func writeFloat64(b *BinaryData, v float64) error {
	var buf [8]byte
	b.order.PutUint64(buf[:], math.Float64bits(v))
	_, err := b.buffer.Write(buf[:])
	return err
}

func readFloat32(b *BinaryData) (float32, error) {
	var buf [4]byte
	if _, err := io.ReadFull(&b.buffer, buf[:]); err != nil {
		return 0, err
	}
	return math.Float32frombits(b.order.Uint32(buf[:])), nil
}

func readFloat64(b *BinaryData) (float64, error) {
	var buf [8]byte
	if _, err := io.ReadFull(&b.buffer, buf[:]); err != nil {
		return 0, err
	}
	return math.Float64frombits(b.order.Uint64(buf[:])), nil
}

// 自動根據型態調變長度寫入
func insertNumber[T NumberX](b *BinaryData, v T) error {
	switch val := any(v).(type) {
	case int8:
		return writeVarint(b, int64(val))
	case int16:
		return writeVarint(b, int64(val))
	case int32:
		return writeVarint(b, int64(val))
	case int64:
		return writeVarint(b, val)
	case uint8:
		return writeUvarint(b, uint64(val))
	case uint16:
		return writeUvarint(b, uint64(val))
	case uint32:
		return writeUvarint(b, uint64(val))
	case uint64:
		return writeUvarint(b, val)
	case float32:
		return writeFloat32(b, val)
	case float64:
		return writeFloat64(b, val)
	}
	return nil
}

// 與寫入邏輯對齊的自動長度讀取
func popNumber[T NumberX](b *BinaryData) (T, error) {
	var v T
	switch any(v).(type) {
	case int8:
		val, err := binary.ReadVarint(&b.buffer)
		if err != nil {
			return v, err
		}
		return any(int8(val)).(T), nil
	case int16:
		val, err := binary.ReadVarint(&b.buffer)
		if err != nil {
			return v, err
		}
		return any(int16(val)).(T), nil
	case int32:
		val, err := binary.ReadVarint(&b.buffer)
		if err != nil {
			return v, err
		}
		return any(int32(val)).(T), nil
	case int64:
		val, err := binary.ReadVarint(&b.buffer)
		if err != nil {
			return v, err
		}
		return any(val).(T), nil
	case uint8:
		val, err := binary.ReadUvarint(&b.buffer)
		if err != nil {
			return v, err
		}
		return any(uint8(val)).(T), nil
	case uint16:
		val, err := binary.ReadUvarint(&b.buffer)
		if err != nil {
			return v, err
		}
		return any(uint16(val)).(T), nil
	case uint32:
		val, err := binary.ReadUvarint(&b.buffer)
		if err != nil {
			return v, err
		}
		return any(uint32(val)).(T), nil
	case uint64:
		val, err := binary.ReadUvarint(&b.buffer)
		if err != nil {
			return v, err
		}
		return any(val).(T), nil
	case float32:
		val, err := readFloat32(b)
		if err != nil {
			return v, err
		}
		return any(val).(T), nil
	case float64:
		val, err := readFloat64(b)
		if err != nil {
			return v, err
		}
		return any(val).(T), nil
	}
	return v, nil
}

// InsertNumber 統一標量數字寫入接頭，依實際型態自動選擇編碼方式
func (b *BinaryData) InsertNumber(v any) error {
	switch val := v.(type) {
	case int8:
		return insertNumber(b, val)
	case int16:
		return insertNumber(b, val)
	case int32:
		return insertNumber(b, val)
	case int64:
		return insertNumber(b, val)
	case int:
		return writeVarint(b, int64(val))
	case uint8:
		return insertNumber(b, val)
	case uint16:
		return insertNumber(b, val)
	case uint32:
		return insertNumber(b, val)
	case uint64:
		return insertNumber(b, val)
	case uint:
		return writeUvarint(b, uint64(val))
	case float32:
		return insertNumber(b, val)
	case float64:
		return insertNumber(b, val)
	default:
		return errors.Errorf("unsupported number type: %T", v)
	}
}

// PopNumber 統一標量數字讀取接頭，依 sample 型態自動選擇解碼方式
func (b *BinaryData) PopNumber(sample any) (any, error) {
	switch sample.(type) {
	case int8:
		return popNumber[int8](b)
	case int16:
		return popNumber[int16](b)
	case int32:
		return popNumber[int32](b)
	case int64:
		return popNumber[int64](b)
	case int:
		val, err := binary.ReadVarint(&b.buffer)
		if err != nil {
			return 0, err
		}
		return int(val), nil
	case uint8:
		return popNumber[uint8](b)
	case uint16:
		return popNumber[uint16](b)
	case uint32:
		return popNumber[uint32](b)
	case uint64:
		return popNumber[uint64](b)
	case uint:
		val, err := binary.ReadUvarint(&b.buffer)
		if err != nil {
			return uint(0), err
		}
		return uint(val), nil
	case float32:
		return popNumber[float32](b)
	case float64:
		return popNumber[float64](b)
	default:
		return nil, errors.Errorf("unsupported number type: %T", sample)
	}
}
