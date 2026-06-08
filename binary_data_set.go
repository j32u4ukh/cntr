package cntr

import (
	"github.com/pkg/errors"
)

// ==================================================
// 統一標量寫入接頭
// ==================================================

// AddNumber 統一處理所有單一數字寫入，內部依型態自動選擇 Varint / Uvarint / IEEE 754 編碼。
// 新程式碼建議直接使用此接頭；下方 AddInt32、AddFloat64 等為向後相容的薄包裝。
func (b *BinaryData) AddNumber(v any) error {
	err := b.InsertNumber(v)
	if err != nil {
		return errors.Wrapf(err, "Failed to write scalar number data: %+v", v)
	}
	return nil
}

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

// AddBoolean 以變長 uint8 寫入（0 或 1），佔用 1 Byte。
func (b *BinaryData) AddBoolean(data bool) error {
	if data {
		return b.AddNumber(uint8(1))
	}
	return b.AddNumber(uint8(0))
}

// 以下為各型態的相容寫入接頭，內部皆委派至 AddNumber。

func (b *BinaryData) AddInt8(data int8) error {
	err := b.AddNumber(data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddInt16(data int16) error {
	err := b.AddNumber(data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddInt32(data int32) error {
	err := b.AddNumber(data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddInt64(data int64) error {
	err := b.AddNumber(data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddByte(data byte) error {
	err := b.AddNumber(data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddUInt16(data uint16) error {
	err := b.AddNumber(data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddUInt32(data uint32) error {
	err := b.AddNumber(data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddUInt64(data uint64) error {
	err := b.AddNumber(data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %d", data)
	}
	return nil
}

func (b *BinaryData) AddFloat32(data float32) error {
	err := b.AddNumber(data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %f", data)
	}
	return nil
}

func (b *BinaryData) AddFloat64(data float64) error {
	err := b.AddNumber(data)
	if err != nil {
		return errors.Wrapf(err, "Failed to write data: %f", data)
	}
	return nil
}

// AddString 將字串壓為 byte 切片，交給陣列協議寫入「長度 + 數據」。
func (b *BinaryData) AddString(data string) error {
	return b.AddByteArray([]byte(data))
}

// ==================================================
// Add Array
// 協議：先寫入變長長度前綴，再依序寫入各元素（元素本身亦為變長編碼）。
// ==================================================

func (b *BinaryData) AddInt32Array(values []int32) error {
	// 明確處理 nil，將其視為空 slice
	if values == nil {
		values = []int32{}
	}
	length := uint32(len(values))
	if err := b.AddNumber(length); err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	for _, value := range values {
		if err := b.AddNumber(value); err != nil {
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
	if err := b.AddNumber(length); err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	for _, value := range values {
		if err := b.AddNumber(value); err != nil {
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
	if err := b.AddNumber(length); err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	// 只有在有數據時才進行寫入
	if length > 0 {
		_, err := b.buffer.Write(data)
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
	if err := b.AddNumber(length); err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	for _, value := range values {
		if err := b.AddNumber(value); err != nil {
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
	if err := b.AddNumber(length); err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	for _, value := range values {
		if err := b.AddNumber(value); err != nil {
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
	if err := b.AddNumber(length); err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	for _, value := range values {
		if err := b.AddNumber(value); err != nil {
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
	if err := b.AddNumber(length); err != nil {
		return errors.Wrapf(err, "Faield to write length of byte array: %d", length)
	}
	for _, value := range values {
		if err := b.AddNumber(value); err != nil {
			return errors.Wrapf(err, "Faield to write float64 data: %f", value)
		}
	}
	return nil
}

// ==================================================
// Add Map
// 協議：先寫入條目數，再依序寫入 key + value（value 可為陣列或字串等複合型態）。
// ==================================================

func (b *BinaryData) AddMapInt64Int64Array(data map[int64][]int64) error {
	if data == nil {
		data = make(map[int64][]int64)
	}
	length := uint32(len(data))
	if err := b.AddNumber(length); err != nil {
		return errors.Wrapf(err, "Failed to write length of map: %d", length)
	}
	for k, v := range data {
		if err := b.AddNumber(k); err != nil {
			return errors.Wrapf(err, "Failed to write key of map: %d", k)
		}
		if err := b.AddInt64Array(v); err != nil {
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
	if err := b.AddNumber(length); err != nil {
		return errors.Wrapf(err, "Failed to write length of map: %d", length)
	}
	for k, v := range data {
		if err := b.AddNumber(k); err != nil {
			return errors.Wrapf(err, "Failed to write key of map: %d", k)
		}
		if err := b.AddUInt32Array(v); err != nil {
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
	if err := b.AddNumber(length); err != nil {
		return errors.Wrapf(err, "Failed to write length of map: %d", length)
	}
	for k, v := range data {
		if err := b.AddString(k); err != nil {
			return errors.Wrapf(err, "Failed to write key of map: %s", k)
		}
		if err := b.AddString(v); err != nil {
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
	if err := b.AddNumber(length); err != nil {
		return errors.Wrapf(err, "Failed to write length of map: %d", length)
	}
	for k, v := range data {
		if err := b.AddString(k); err != nil {
			return errors.Wrapf(err, "Failed to write key of map: %s", k)
		}
		if err := b.AddByteArray(v); err != nil {
			return errors.Wrapf(err, "Failed to write value of map: %+v", v)
		}
	}
	return nil
}

// ==================================================
// 插入數據(目前只能插在最前面)
// ==================================================

func (b *BinaryData) InsertInt32(data int32) error {
	err := b.InsertNumber(data)
	if err != nil {
		return errors.Wrapf(err, "Failed to insert int32 data: %d", data)
	}
	return nil
}

func (b *BinaryData) InsertUInt32(data uint32) error {
	err := b.InsertNumber(data)
	if err != nil {
		return errors.Wrapf(err, "Failed to insert uint32 data: %d", data)
	}
	return nil
}
