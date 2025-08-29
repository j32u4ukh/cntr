package cntr

import (
	"encoding/binary"

	"github.com/pkg/errors"
)

// ==================================================
// 加入數據
// ==================================================

func (b *BinaryData) AddRawData(data []byte) error {
	if data == nil {
		data = []byte{}
	}
	if len(data) > 0 {
		_, err := b.buffer.Write(data)
		if err != nil {
			return errors.Wrapf(err, "Failed to write data: %+v", data)
		}
	}
	return nil
}

func (b *BinaryData) AddBoolean(data bool) error {
	err := binary.Write(&b.buffer, b.order, data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %+v", data)
	}
	return nil
}

func (b *BinaryData) AddInt8(data int8) error {
	err := binary.Write(&b.buffer, b.order, data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddInt16(data int16) error {
	err := binary.Write(&b.buffer, b.order, data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddInt32(data int32) error {
	err := binary.Write(&b.buffer, b.order, data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddInt64(data int64) error {
	err := binary.Write(&b.buffer, b.order, data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddByte(data byte) error {
	err := b.buffer.WriteByte(data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddUInt16(data uint16) error {
	err := binary.Write(&b.buffer, b.order, data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddUInt32(data uint32) error {
	err := binary.Write(&b.buffer, b.order, data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddUInt64(data uint64) error {
	err := binary.Write(&b.buffer, b.order, data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddFloat32(data float32) error {
	err := binary.Write(&b.buffer, b.order, data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %f", data)
	}
	return nil
}

func (b *BinaryData) AddFloat64(data float64) error {
	err := binary.Write(&b.buffer, b.order, data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %f", data)
	}
	return nil
}

func (b *BinaryData) AddString(data string) error {
	err := b.AddByteArray([]byte(data))
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %s", data)
	}
	return nil
}

// ==================================================
// Add Array
// ==================================================

func (b *BinaryData) AddInt32Array(values []int32) error {
	// 明確處理 nil，將其視為空 slice
	if values == nil {
		values = []int32{}
	}
	length := uint32(len(values))
	err := b.AddUInt32(length)
	if err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	for _, value := range values {
		err = b.AddInt32(value)
		if err != nil {
			return errors.Wrapf(err, "Faield to write int32 data: %d", value)
		}
	}
	return nil
}

func (b *BinaryData) AddInt64Array(values []int64) error {
	// 明確處理 nil，將其視為空 slice
	if values == nil {
		values = []int64{}
	}
	length := uint32(len(values))
	err := b.AddUInt32(length)
	if err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	for _, value := range values {
		err = b.AddInt64(value)
		if err != nil {
			return errors.Wrapf(err, "Faield to write int64 data: %d", value)
		}
	}
	return nil
}

func (b *BinaryData) AddByteArray(data []byte) error {
	// 明確處理 nil，將其視為空 slice
	if data == nil {
		data = []byte{}
	}
	length := uint32(len(data))
	err := b.AddUInt32(length)
	if err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	// 只有在有數據時才進行寫入
	if length > 0 {
		_, err = b.buffer.Write(data)
		if err != nil {
			return errors.Wrapf(err, "Failed to write data: %+v", data)
		}
	}
	return nil
}

func (b *BinaryData) AddUInt32Array(values []uint32) error {
	// 明確處理 nil，將其視為空 slice
	if values == nil {
		values = []uint32{}
	}
	length := uint32(len(values))
	err := b.AddUInt32(length)
	if err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	for _, value := range values {
		err = b.AddUInt32(value)
		if err != nil {
			return errors.Wrapf(err, "Faield to write uint32 data: %d", value)
		}
	}
	return nil
}

func (b *BinaryData) AddUInt64Array(values []uint64) error {
	// 明確處理 nil，將其視為空 slice
	if values == nil {
		values = []uint64{}
	}
	length := uint32(len(values))
	err := b.AddUInt32(length)
	if err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	for _, value := range values {
		err = b.AddUInt64(value)
		if err != nil {
			return errors.Wrapf(err, "Faield to write uint64 data: %d", value)
		}
	}
	return nil
}

func (b *BinaryData) AddFloat32Array(values []float32) error {
	// 明確處理 nil，將其視為空 slice
	if values == nil {
		values = []float32{}
	}
	length := uint32(len(values))
	err := b.AddUInt32(length)
	if err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	for _, value := range values {
		err = b.AddFloat32(value)
		if err != nil {
			return errors.Wrapf(err, "Faield to write float32 data: %f", value)
		}
	}
	return nil
}

func (b *BinaryData) AddFloat64Array(values []float64) error {
	// 明確處理 nil，將其視為空 slice
	if values == nil {
		values = []float64{}
	}
	length := uint32(len(values))
	err := b.AddUInt32(length)
	if err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	for _, value := range values {
		err = b.AddFloat64(value)
		if err != nil {
			return errors.Wrapf(err, "Faield to write float64 data: %f", value)
		}
	}
	return nil
}

// ==================================================
// Add Map
// ==================================================

func (b *BinaryData) AddMapInt64Int64Array(data map[int64][]int64) error {
	if data == nil {
		data = make(map[int64][]int64)
	}
	length := uint32(len(data))
	err := b.AddUInt32(length)
	if err != nil {
		return errors.Wrapf(err, "Failed to write length of map: %d", length)
	}
	for k, v := range data {
		err = b.AddInt64(k)
		if err != nil {
			return errors.Wrapf(err, "Failed to write key of map: %d", k)
		}
		err = b.AddInt64Array(v)
		if err != nil {
			return errors.Wrapf(err, "Failed to write value of map: %+v", v)
		}
	}
	return nil
}

func (b *BinaryData) AddMapUInt32UInt32Array(data map[uint32][]uint32) error {
	if data == nil {
		data = make(map[uint32][]uint32)
	}
	length := uint32(len(data))
	err := b.AddUInt32(length)
	if err != nil {
		return errors.Wrapf(err, "Failed to write length of map: %d", length)
	}
	for k, v := range data {
		err = b.AddUInt32(k)
		if err != nil {
			return errors.Wrapf(err, "Failed to write key of map: %d", k)
		}
		err = b.AddUInt32Array(v)
		if err != nil {
			return errors.Wrapf(err, "Failed to write value of map: %+v", v)
		}
	}
	return nil
}

func (b *BinaryData) AddMapStringString(data map[string]string) error {
	if data == nil {
		data = make(map[string]string)
	}
	length := uint32(len(data))
	err := b.AddUInt32(length)
	if err != nil {
		return errors.Wrapf(err, "Failed to write length of map: %d", length)
	}
	for k, v := range data {
		err = b.AddString(k)
		if err != nil {
			return errors.Wrapf(err, "Failed to write key of map: %s", k)
		}
		err = b.AddString(v)
		if err != nil {
			return errors.Wrapf(err, "Failed to write value of map: %s", v)
		}
	}
	return nil
}

func (b *BinaryData) AddMapStringByteArray(data map[string][]byte) error {
	if data == nil {
		data = make(map[string][]byte)
	}
	length := uint32(len(data))
	err := b.AddUInt32(length)
	if err != nil {
		return errors.Wrapf(err, "Failed to write length of map: %d", length)
	}
	for k, v := range data {
		err = b.AddString(k)
		if err != nil {
			return errors.Wrapf(err, "Failed to write key of map: %s", k)
		}
		err = b.AddByteArray(v)
		if err != nil {
			return errors.Wrapf(err, "Failed to write value of map: %+v", v)
		}
	}
	return nil
}

// ==================================================
// 插入數據(目前只能插在最前面)
// ==================================================

func (b *BinaryData) InsertInt32(data int32) error {
	err := insertNumber(b, data)
	if err != nil {
		return errors.Wrapf(err, "Failed to insert int32 data: %d", data)
	}
	return nil
}

func (b *BinaryData) InsertUInt32(data uint32) error {
	err := insertNumber(b, data)
	if err != nil {
		return errors.Wrapf(err, "Failed to insert uint32 data: %d", data)
	}
	return nil
}
