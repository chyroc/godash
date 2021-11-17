package godash

import (
	"math"
)

// Creates an array of elements splits into groups the length of size. If array
// can't be split evenly, the final chunk will be the remaining elements.

//go:generate make generate FILE=chunk FUNC=chunk TYPE=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,string,float32,float64,bool

func ChunkInt(array []int, size int) [][]int {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]int{}
	}

	index := 0
	resIndex := 0

	result := make([][]int, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}

func ChunkInt8(array []int8, size int) [][]int8 {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]int8{}
	}

	index := 0
	resIndex := 0

	result := make([][]int8, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}

func ChunkInt16(array []int16, size int) [][]int16 {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]int16{}
	}

	index := 0
	resIndex := 0

	result := make([][]int16, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}

func ChunkInt32(array []int32, size int) [][]int32 {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]int32{}
	}

	index := 0
	resIndex := 0

	result := make([][]int32, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}

func ChunkInt64(array []int64, size int) [][]int64 {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]int64{}
	}

	index := 0
	resIndex := 0

	result := make([][]int64, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}

func ChunkUint(array []uint, size int) [][]uint {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]uint{}
	}

	index := 0
	resIndex := 0

	result := make([][]uint, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}

func ChunkUint8(array []uint8, size int) [][]uint8 {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]uint8{}
	}

	index := 0
	resIndex := 0

	result := make([][]uint8, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}

func ChunkUint16(array []uint16, size int) [][]uint16 {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]uint16{}
	}

	index := 0
	resIndex := 0

	result := make([][]uint16, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}

func ChunkUint32(array []uint32, size int) [][]uint32 {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]uint32{}
	}

	index := 0
	resIndex := 0

	result := make([][]uint32, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}

func ChunkUint64(array []uint64, size int) [][]uint64 {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]uint64{}
	}

	index := 0
	resIndex := 0

	result := make([][]uint64, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}

func ChunkString(array []string, size int) [][]string {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]string{}
	}

	index := 0
	resIndex := 0

	result := make([][]string, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}

func ChunkFloat32(array []float32, size int) [][]float32 {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]float32{}
	}

	index := 0
	resIndex := 0

	result := make([][]float32, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}

func ChunkFloat64(array []float64, size int) [][]float64 {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]float64{}
	}

	index := 0
	resIndex := 0

	result := make([][]float64, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}

func ChunkBool(array []bool, size int) [][]bool {
	length := len(array)
	if length == 0 || size < 1 {
		return [][]bool{}
	}

	index := 0
	resIndex := 0

	result := make([][]bool, int(math.Ceil(float64(length)/float64(size))))
	for index < length {
		end := int(math.Min(float64(index+size), float64(length)))
		result[resIndex] = array[index:end]
		index = end
		resIndex++
	}
	return result
}