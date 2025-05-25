package cntr

import (
	"encoding/binary"

	"github.com/pkg/errors"
)

// 取出全部的數據
func (b BinaryData) GetData() []byte {
	return b.buffer.Bytes()
}

// 讀取 byte 陣列
func (b *BinaryData) FetchByteArray(length uint32) ([]byte, error) {
	// 如果長度為 0，直接返回空 slice
	if length == 0 {
		return []byte{}, nil
	}
	result := make([]byte, length)
	err := binary.Read(&b.buffer, b.order, result)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read byte array")
	}
	return result, nil
}

func (b *BinaryData) PopBoolean() (bool, error) {
	boolean, err := b.PopByte()
	if err != nil {
		return false, errors.Wrap(err, "Failed to read bool data")
	}
	return boolean == 1, nil
}

func (b *BinaryData) PopInt8() (int8, error) {
	value, err := popNumber[int8](b)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read int8 data")
	}
	return value, nil
}

func (b *BinaryData) PopInt16() (int16, error) {
	value, err := popNumber[int16](b)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read int16 data")
	}
	return value, nil
}

func (b *BinaryData) PopInt32() (int32, error) {
	value, err := popNumber[int32](b)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read int32 data")
	}
	return value, nil
}

func (b *BinaryData) PopInt64() (int64, error) {
	value, err := popNumber[int64](b)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read int64 data")
	}
	return value, nil
}

func (b *BinaryData) PopByte() (byte, error) {
	value, err := popNumber[uint8](b)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read uint8 data")
	}
	return value, nil
}

func (b *BinaryData) PopUInt16() (uint16, error) {
	value, err := popNumber[uint16](b)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read uint16 data")
	}
	return value, nil
}

func (b *BinaryData) PopUInt32() (uint32, error) {
	value, err := popNumber[uint32](b)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read uint32 data")
	}
	return value, nil
}

func (b *BinaryData) PopUInt64() (uint64, error) {
	value, err := popNumber[uint64](b)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read uint64 data")
	}
	return value, nil
}

func (b *BinaryData) PopFloat32() (float32, error) {
	value, err := popNumber[float32](b)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read float32 data")
	}
	return value, nil
}

func (b *BinaryData) PopFloat64() (float64, error) {
	value, err := popNumber[float64](b)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read float64 data")
	}
	return value, nil
}

func (b *BinaryData) PopString() (string, error) {
	result, err := b.PopByteArray()
	if err != nil {
		return "", errors.Wrap(err, "Faield to read string data")
	}
	return string(result), nil
}

// ==================================================
// Pop Array
// ==================================================

func (b *BinaryData) PopInt32Array() ([]int32, error) {
	length, err := b.PopUInt32()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of int32 array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []int32{}, nil
	}
	result := make([]int32, length)
	for i := uint32(0); i < length; i++ {
		result[i], err = b.PopInt32()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read data of int32 array")
		}
	}
	return result, nil
}

func (b *BinaryData) PopInt64Array() ([]int64, error) {
	length, err := b.PopUInt32()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of int64 array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []int64{}, nil
	}
	result := make([]int64, length)
	for i := uint32(0); i < length; i++ {
		result[i], err = b.PopInt64()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read data of int64 array")
		}
	}
	return result, nil
}

func (b *BinaryData) PopByteArray() ([]byte, error) {
	length, err := b.PopUInt32()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of byte array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []byte{}, nil
	}
	result, err := b.FetchByteArray(length)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch byte array")
	}
	return result, nil
}

func (b *BinaryData) PopUInt32Array() ([]uint32, error) {
	length, err := b.PopUInt32()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of uint32 array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []uint32{}, nil
	}
	result := make([]uint32, length)
	for i := uint32(0); i < length; i++ {
		result[i], err = b.PopUInt32()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read data of uint32 array")
		}
	}
	return result, nil
}

func (b *BinaryData) PopUInt64Array() ([]uint64, error) {
	length, err := b.PopUInt32()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of uint64 array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []uint64{}, nil
	}
	result := make([]uint64, length)
	for i := uint32(0); i < length; i++ {
		result[i], err = b.PopUInt64()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read data of uint64 array")
		}
	}
	return result, nil
}

func (b *BinaryData) PopFloat32Array() ([]float32, error) {
	length, err := b.PopUInt32()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of float32 array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []float32{}, nil
	}
	result := make([]float32, length)
	for i := uint32(0); i < length; i++ {
		result[i], err = b.PopFloat32()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read data of float32 array")
		}
	}
	return result, nil
}


func (b *BinaryData) PopFloat64Array() ([]float64, error) {
	length, err := b.PopUInt32()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of float64 array")
	}
	// 明確處理空陣列的情況
	if length == 0 {
		return []float64{}, nil
	}
	result := make([]float64, length)
	for i := uint32(0); i < length; i++ {
		result[i], err = b.PopFloat64()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read data of float64 array")
		}
	}
	return result, nil
}

// ==================================================
// Pop Map
// ==================================================

func (b *BinaryData) PopMapInt64Int64Array() (map[int64][]int64, error) {
	result := map[int64][]int64{}
	length, err := b.PopInt32()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of map")
	}
	var key int64
	var value []int64
	for i := int32(0); i < length; i++ {
		key, err = b.PopInt64()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read key of map")
		}
		value, err = b.PopInt64Array()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read value of map")
		}
		result[key] = value
	}
	return result, nil
}

func (b *BinaryData) PopMapUInt32UInt32Array() (map[uint32][]uint32, error) {
	result := map[uint32][]uint32{}
	length, err := b.PopInt32()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of map")
	}
	var key uint32
	var value []uint32
	for i := int32(0); i < length; i++ {
		key, err = b.PopUInt32()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read key of map")
		}
		value, err = b.PopUInt32Array()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read value of map")
		}
		result[key] = value
	}
	return result, nil
}


func (b *BinaryData) PopMapStringString() (map[string]string, error) {
	result := map[string]string{}
	length, err := b.PopInt32()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of map")
	}
	var key, value string
	for i := int32(0); i < length; i++ {
		key, err = b.PopString()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read key of map")
		}
		value, err = b.PopString()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read value of map")
		}
		result[key] = value
	}
	return result, nil
}

func (b *BinaryData) PopMapStringByteArray() (map[string][]byte, error) {
	result := map[string][]byte{}
	length, err := b.PopInt32()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read length of map")
	}
	var key string
	var value []byte
	for i := int32(0); i < length; i++ {
		key, err = b.PopString()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read key of map")
		}
		value, err = b.PopByteArray()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read value of map")
		}
		result[key] = value
	}
	return result, nil
}