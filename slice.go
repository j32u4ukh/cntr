package cntr

import (
	"bytes"
	"fmt"
)

// 二維 Slice 轉字串
func Slice2dToString[T Element](slice2d [][]T) string {
	var buffer bytes.Buffer
	length := len(slice2d)
	buffer.WriteString("{")
	if length > 0 {
		buffer.WriteString(SliceToString(slice2d[0]))
		for i := 1; i < length; i++ {
			buffer.WriteString(fmt.Sprintf(", %s", SliceToString(slice2d[0])))
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}

// Slice 轉字串
func SliceToString[T Element](slice []T) string {
	var buffer bytes.Buffer
	length := len(slice)
	buffer.WriteString("{")
	if length > 0 {
		buffer.WriteString(fmt.Sprintf("%v", slice[0]))
		for i := 1; i < length; i++ {
			buffer.WriteString(fmt.Sprintf(", %v", slice[i]))
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}

// 比較兩個二維 Slice 是否相同
func IsSlice2dEqual[T Element](a, b [][]T) bool {
	nA := len(a)
	nB := len(b)
	if nA != nB {
		return false
	}
	for i := 0; i < nA; i++ {
		if !IsSliceEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

// 比較兩個 Slice 是否相同
func IsSliceEqual[T Element](a, b []T) bool {
	nA := len(a)
	nB := len(b)
	if nA != nB {
		return false
	}
	for i := 0; i < nA; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
