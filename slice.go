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

// Slice3DArray 三維子陣列：先按高度（Y 軸），再按寬度（X 軸），最後按深度（Z 軸）
func SliceArray3D[T Element](array [][][]T, zStart, zEnd, yStart, yEnd, xStart, xEnd int) [][][]T {
	// 檢查索引是否合法
	if yStart < 0 || yEnd > len(array) || xStart < 0 || xEnd > len(array[0]) ||
		zStart < 0 || zEnd > len(array[0][0]) || yStart > yEnd || xStart > xEnd || zStart > zEnd {
		return nil
	}
	// 初始化三維子陣列
	subArray := make([][][]T, zEnd-zStart)
	// 初始化每個切片的寬度部分
	for i := zStart; i < zEnd; i++ {
		subArray[i-zStart] = SliceArray2D(array[i], yStart, yEnd, xStart, xEnd)
	}
	return subArray
}

// SliceArray2D 函數傳入二維陣列以及 X、Y 軸的開始與結束索引，返回子二維陣列（不包含結束索引值）
func SliceArray2D[T Element](array [][]T, yStart, yEnd, xStart, xEnd int) [][]T {
	// 檢查索引是否合法
	if yStart < 0 || yEnd > len(array) || xStart < 0 || xEnd > len(array[0]) || xStart > xEnd || yStart > yEnd {
		return nil
	}
	// 初始化二子維陣列
	subArray := make([][]T, yEnd-yStart)
	for i := yStart; i < yEnd; i++ {
		// 擷取每一行的子切片
		subArray[i-yStart] = array[i][xStart:xEnd]
	}
	return subArray
}
