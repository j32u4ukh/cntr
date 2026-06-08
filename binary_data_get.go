package cntr

import (
	"io"

	"github.com/pkg/errors"
)

// ==================================================
// 統一標量讀取接頭
// ==================================================

// PopScalar 統一處理所有單一數字讀取，依 sample 型態自動選擇 Varint / Uvarint / IEEE 754 解碼。
// 調用範例：v, err := b.PopScalar(int32(0)); id := v.(int32)
// 新程式碼建議直接使用此接頭；下方 PopInt32、PopFloat64 等為向後相容的薄包裝。
func (b *BinaryData) PopScalar(sample any) (any, error) {
	val, err := b.PopNumber(sample)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to decode scalar number")
	}
	return val, nil
}

// popLength 供陣列 / Map 協議使用，先解開變長長度前綴。
func (b *BinaryData) popLength() (uint32, error) {
	val, err := b.PopScalar(uint32(0))
	if err != nil {
		return 0, err
	}
	return val.(uint32), nil
}

// 取出全部的數據
func (b BinaryData) GetData() []byte {
	return b.buffer.Bytes()
}

// FetchByteArray 依指定長度從緩衝區讀取原始位元組；回傳獨立切片，避免與內部 buffer 共享記憶體。
func (b *BinaryData) FetchByteArray(length uint32) ([]byte, error) {
	// 如果長度為 0，直接返回空 slice
	if length == 0 {
		return []byte{}, nil
	}
	buf := b.buffer.Next(int(length))
	if len(buf) < int(length) {
		return nil, io.EOF
	}
	result := make([]byte, length)
	copy(result, buf)
	return result, nil
}

// PopBoolean 讀取寫入端壓入的變長 uint8（0 或 1）。
func (b *BinaryData) PopBoolean() (bool, error) {
	val, err := b.PopScalar(uint8(0))
	if err != nil {
		return false, errors.Wrap(err, "Failed to read bool data")
	}
	return val.(uint8) == 1, nil
}

// 以下為各型態的相容讀取接頭，內部皆委派至 PopScalar。

func (b *BinaryData) PopInt8() (int8, error) {
	val, err := b.PopScalar(int8(0))
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read int8 data")
	}
	return val.(int8), nil
}

func (b *BinaryData) PopInt16() (int16, error) {
	val, err := b.PopScalar(int16(0))
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read int16 data")
	}
	return val.(int16), nil
}

func (b *BinaryData) PopInt32() (int32, error) {
	val, err := b.PopScalar(int32(0))
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read int32 data")
	}
	return val.(int32), nil
}

func (b *BinaryData) PopInt64() (int64, error) {
	val, err := b.PopScalar(int64(0))
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read int64 data")
	}
	return val.(int64), nil
}

func (b *BinaryData) PopByte() (byte, error) {
	val, err := b.PopScalar(uint8(0))
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read uint8 data")
	}
	return val.(uint8), nil
}

func (b *BinaryData) PopUInt16() (uint16, error) {
	val, err := b.PopScalar(uint16(0))
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read uint16 data")
	}
	return val.(uint16), nil
}

func (b *BinaryData) PopUInt32() (uint32, error) {
	val, err := b.PopScalar(uint32(0))
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read uint32 data")
	}
	return val.(uint32), nil
}

func (b *BinaryData) PopUInt64() (uint64, error) {
	val, err := b.PopScalar(uint64(0))
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read uint64 data")
	}
	return val.(uint64), nil
}

func (b *BinaryData) PopFloat32() (float32, error) {
	val, err := b.PopScalar(float32(0))
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read float32 data")
	}
	return val.(float32), nil
}

func (b *BinaryData) PopFloat64() (float64, error) {
	val, err := b.PopScalar(float64(0))
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read float64 data")
	}
	return val.(float64), nil
}

// PopString 透過 PopByteArray 還原 byte 切片後轉回字串。
func (b *BinaryData) PopString() (string, error) {
	result, err := b.PopByteArray()
	if err != nil {
		return "", errors.Wrap(err, "Faield to read string data")
	}
	return string(result), nil
}

// ==================================================
// Pop Array
// 協議：先讀取變長長度前綴，再依序還原各元素。
// ==================================================

func (b *BinaryData) PopInt32Array() ([]int32, error) {
	length, err := b.popLength()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of int32 array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []int32{}, nil
	}
	result := make([]int32, length)
	for i := uint32(0); i < length; i++ {
		val, err := b.PopScalar(int32(0))
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read data of int32 array")
		}
		result[i] = val.(int32)
	}
	return result, nil
}

func (b *BinaryData) PopInt64Array() ([]int64, error) {
	length, err := b.popLength()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of int64 array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []int64{}, nil
	}
	result := make([]int64, length)
	for i := uint32(0); i < length; i++ {
		val, err := b.PopScalar(int64(0))
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read data of int64 array")
		}
		result[i] = val.(int64)
	}
	return result, nil
}

func (b *BinaryData) PopByteArray() ([]byte, error) {
	length, err := b.popLength()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of byte array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []byte{}, nil
	}
	return b.FetchByteArray(length)
}

func (b *BinaryData) PopUInt32Array() ([]uint32, error) {
	length, err := b.popLength()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of uint32 array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []uint32{}, nil
	}
	result := make([]uint32, length)
	for i := uint32(0); i < length; i++ {
		val, err := b.PopScalar(uint32(0))
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read data of uint32 array")
		}
		result[i] = val.(uint32)
	}
	return result, nil
}

func (b *BinaryData) PopUInt64Array() ([]uint64, error) {
	length, err := b.popLength()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of uint64 array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []uint64{}, nil
	}
	result := make([]uint64, length)
	for i := uint32(0); i < length; i++ {
		val, err := b.PopScalar(uint64(0))
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read data of uint64 array")
		}
		result[i] = val.(uint64)
	}
	return result, nil
}

func (b *BinaryData) PopFloat32Array() ([]float32, error) {
	length, err := b.popLength()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of float32 array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []float32{}, nil
	}
	result := make([]float32, length)
	for i := uint32(0); i < length; i++ {
		val, err := b.PopScalar(float32(0))
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read data of float32 array")
		}
		result[i] = val.(float32)
	}
	return result, nil
}

func (b *BinaryData) PopFloat64Array() ([]float64, error) {
	length, err := b.popLength()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of float64 array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []float64{}, nil
	}
	result := make([]float64, length)
	for i := uint32(0); i < length; i++ {
		val, err := b.PopScalar(float64(0))
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read data of float64 array")
		}
		result[i] = val.(float64)
	}
	return result, nil
}

// ==================================================
// Pop Map
// 協議：先讀取條目數，再依序還原 key + value。
// ==================================================

func (b *BinaryData) PopMapInt64Int64Array() (map[int64][]int64, error) {
	length, err := b.popLength()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of map")
	}
	result := make(map[int64][]int64, length)
	for i := uint32(0); i < length; i++ {
		keyVal, err := b.PopScalar(int64(0))
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read key of map")
		}
		arrVal, err := b.PopInt64Array()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read value of map")
		}
		result[keyVal.(int64)] = arrVal
	}
	return result, nil
}

func (b *BinaryData) PopMapUInt32UInt32Array() (map[uint32][]uint32, error) {
	length, err := b.popLength()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of map")
	}
	result := make(map[uint32][]uint32, length)
	for i := uint32(0); i < length; i++ {
		keyVal, err := b.PopScalar(uint32(0))
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read key of map")
		}
		arrVal, err := b.PopUInt32Array()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read value of map")
		}
		result[keyVal.(uint32)] = arrVal
	}
	return result, nil
}

func (b *BinaryData) PopMapStringString() (map[string]string, error) {
	length, err := b.popLength()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of map")
	}
	result := make(map[string]string, length)
	for i := uint32(0); i < length; i++ {
		key, err := b.PopString()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read key of map")
		}
		value, err := b.PopString()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read value of map")
		}
		result[key] = value
	}
	return result, nil
}

func (b *BinaryData) PopMapStringByteArray() (map[string][]byte, error) {
	length, err := b.popLength()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of map")
	}
	result := make(map[string][]byte, length)
	for i := uint32(0); i < length; i++ {
		key, err := b.PopString()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read key of map")
		}
		value, err := b.PopByteArray()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read value of map")
		}
		result[key] = value
	}
	return result, nil
}
