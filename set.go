package godash

import (
	"sort"
)

//go:generate make generate FILE=set FUNC=set TYPE=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,string,float32,float64

type SetInt map[int]struct{}

func NewSetInt() SetInt {
	return make(map[int]struct{})
}

// Diff return new set with all items in s but not in o
func (s SetInt) Diff(o SetInt) SetInt {
	result := make(SetInt, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s SetInt) Inter(o SetInt) SetInt {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(SetInt, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s SetInt) Union(o SetInt) SetInt {
	result := make(SetInt, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s SetInt) Slice() []int {
	result := make([]int, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s SetInt) SortedSlice() []int {
	slice := listInt(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s SetInt) ReversedSortedSlice() []int {
	slice := listInt(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s SetInt) Add(id int) SetInt {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s SetInt) Del(id int) SetInt {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s SetInt) IsExist(id int) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s SetInt) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s SetInt) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s SetInt) DelSet(o SetInt) SetInt {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s SetInt) AddSlice(ids []int) SetInt {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s SetInt) AddSet(o SetInt) SetInt {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type listInt []int

func (s listInt) Len() int           { return len(s) }
func (s listInt) Less(i, j int) bool { return s[i] < s[j] }
func (s listInt) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SetInt8 map[int8]struct{}

func NewSetInt8() SetInt8 {
	return make(map[int8]struct{})
}

// Diff return new set with all items in s but not in o
func (s SetInt8) Diff(o SetInt8) SetInt8 {
	result := make(SetInt8, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s SetInt8) Inter(o SetInt8) SetInt8 {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(SetInt8, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s SetInt8) Union(o SetInt8) SetInt8 {
	result := make(SetInt8, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s SetInt8) Slice() []int8 {
	result := make([]int8, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s SetInt8) SortedSlice() []int8 {
	slice := listInt8(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s SetInt8) ReversedSortedSlice() []int8 {
	slice := listInt8(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s SetInt8) Add(id int8) SetInt8 {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s SetInt8) Del(id int8) SetInt8 {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s SetInt8) IsExist(id int8) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s SetInt8) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s SetInt8) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s SetInt8) DelSet(o SetInt8) SetInt8 {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s SetInt8) AddSlice(ids []int8) SetInt8 {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s SetInt8) AddSet(o SetInt8) SetInt8 {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type listInt8 []int8

func (s listInt8) Len() int           { return len(s) }
func (s listInt8) Less(i, j int) bool { return s[i] < s[j] }
func (s listInt8) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SetInt16 map[int16]struct{}

func NewSetInt16() SetInt16 {
	return make(map[int16]struct{})
}

// Diff return new set with all items in s but not in o
func (s SetInt16) Diff(o SetInt16) SetInt16 {
	result := make(SetInt16, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s SetInt16) Inter(o SetInt16) SetInt16 {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(SetInt16, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s SetInt16) Union(o SetInt16) SetInt16 {
	result := make(SetInt16, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s SetInt16) Slice() []int16 {
	result := make([]int16, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s SetInt16) SortedSlice() []int16 {
	slice := listInt16(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s SetInt16) ReversedSortedSlice() []int16 {
	slice := listInt16(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s SetInt16) Add(id int16) SetInt16 {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s SetInt16) Del(id int16) SetInt16 {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s SetInt16) IsExist(id int16) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s SetInt16) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s SetInt16) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s SetInt16) DelSet(o SetInt16) SetInt16 {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s SetInt16) AddSlice(ids []int16) SetInt16 {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s SetInt16) AddSet(o SetInt16) SetInt16 {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type listInt16 []int16

func (s listInt16) Len() int           { return len(s) }
func (s listInt16) Less(i, j int) bool { return s[i] < s[j] }
func (s listInt16) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SetInt32 map[int32]struct{}

func NewSetInt32() SetInt32 {
	return make(map[int32]struct{})
}

// Diff return new set with all items in s but not in o
func (s SetInt32) Diff(o SetInt32) SetInt32 {
	result := make(SetInt32, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s SetInt32) Inter(o SetInt32) SetInt32 {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(SetInt32, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s SetInt32) Union(o SetInt32) SetInt32 {
	result := make(SetInt32, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s SetInt32) Slice() []int32 {
	result := make([]int32, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s SetInt32) SortedSlice() []int32 {
	slice := listInt32(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s SetInt32) ReversedSortedSlice() []int32 {
	slice := listInt32(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s SetInt32) Add(id int32) SetInt32 {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s SetInt32) Del(id int32) SetInt32 {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s SetInt32) IsExist(id int32) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s SetInt32) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s SetInt32) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s SetInt32) DelSet(o SetInt32) SetInt32 {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s SetInt32) AddSlice(ids []int32) SetInt32 {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s SetInt32) AddSet(o SetInt32) SetInt32 {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type listInt32 []int32

func (s listInt32) Len() int           { return len(s) }
func (s listInt32) Less(i, j int) bool { return s[i] < s[j] }
func (s listInt32) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SetInt64 map[int64]struct{}

func NewSetInt64() SetInt64 {
	return make(map[int64]struct{})
}

// Diff return new set with all items in s but not in o
func (s SetInt64) Diff(o SetInt64) SetInt64 {
	result := make(SetInt64, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s SetInt64) Inter(o SetInt64) SetInt64 {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(SetInt64, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s SetInt64) Union(o SetInt64) SetInt64 {
	result := make(SetInt64, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s SetInt64) Slice() []int64 {
	result := make([]int64, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s SetInt64) SortedSlice() []int64 {
	slice := listInt64(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s SetInt64) ReversedSortedSlice() []int64 {
	slice := listInt64(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s SetInt64) Add(id int64) SetInt64 {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s SetInt64) Del(id int64) SetInt64 {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s SetInt64) IsExist(id int64) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s SetInt64) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s SetInt64) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s SetInt64) DelSet(o SetInt64) SetInt64 {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s SetInt64) AddSlice(ids []int64) SetInt64 {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s SetInt64) AddSet(o SetInt64) SetInt64 {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type listInt64 []int64

func (s listInt64) Len() int           { return len(s) }
func (s listInt64) Less(i, j int) bool { return s[i] < s[j] }
func (s listInt64) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SetUint map[uint]struct{}

func NewSetUint() SetUint {
	return make(map[uint]struct{})
}

// Diff return new set with all items in s but not in o
func (s SetUint) Diff(o SetUint) SetUint {
	result := make(SetUint, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s SetUint) Inter(o SetUint) SetUint {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(SetUint, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s SetUint) Union(o SetUint) SetUint {
	result := make(SetUint, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s SetUint) Slice() []uint {
	result := make([]uint, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s SetUint) SortedSlice() []uint {
	slice := listUint(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s SetUint) ReversedSortedSlice() []uint {
	slice := listUint(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s SetUint) Add(id uint) SetUint {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s SetUint) Del(id uint) SetUint {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s SetUint) IsExist(id uint) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s SetUint) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s SetUint) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s SetUint) DelSet(o SetUint) SetUint {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s SetUint) AddSlice(ids []uint) SetUint {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s SetUint) AddSet(o SetUint) SetUint {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type listUint []uint

func (s listUint) Len() int           { return len(s) }
func (s listUint) Less(i, j int) bool { return s[i] < s[j] }
func (s listUint) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SetUint8 map[uint8]struct{}

func NewSetUint8() SetUint8 {
	return make(map[uint8]struct{})
}

// Diff return new set with all items in s but not in o
func (s SetUint8) Diff(o SetUint8) SetUint8 {
	result := make(SetUint8, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s SetUint8) Inter(o SetUint8) SetUint8 {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(SetUint8, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s SetUint8) Union(o SetUint8) SetUint8 {
	result := make(SetUint8, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s SetUint8) Slice() []uint8 {
	result := make([]uint8, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s SetUint8) SortedSlice() []uint8 {
	slice := listUint8(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s SetUint8) ReversedSortedSlice() []uint8 {
	slice := listUint8(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s SetUint8) Add(id uint8) SetUint8 {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s SetUint8) Del(id uint8) SetUint8 {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s SetUint8) IsExist(id uint8) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s SetUint8) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s SetUint8) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s SetUint8) DelSet(o SetUint8) SetUint8 {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s SetUint8) AddSlice(ids []uint8) SetUint8 {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s SetUint8) AddSet(o SetUint8) SetUint8 {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type listUint8 []uint8

func (s listUint8) Len() int           { return len(s) }
func (s listUint8) Less(i, j int) bool { return s[i] < s[j] }
func (s listUint8) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SetUint16 map[uint16]struct{}

func NewSetUint16() SetUint16 {
	return make(map[uint16]struct{})
}

// Diff return new set with all items in s but not in o
func (s SetUint16) Diff(o SetUint16) SetUint16 {
	result := make(SetUint16, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s SetUint16) Inter(o SetUint16) SetUint16 {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(SetUint16, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s SetUint16) Union(o SetUint16) SetUint16 {
	result := make(SetUint16, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s SetUint16) Slice() []uint16 {
	result := make([]uint16, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s SetUint16) SortedSlice() []uint16 {
	slice := listUint16(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s SetUint16) ReversedSortedSlice() []uint16 {
	slice := listUint16(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s SetUint16) Add(id uint16) SetUint16 {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s SetUint16) Del(id uint16) SetUint16 {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s SetUint16) IsExist(id uint16) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s SetUint16) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s SetUint16) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s SetUint16) DelSet(o SetUint16) SetUint16 {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s SetUint16) AddSlice(ids []uint16) SetUint16 {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s SetUint16) AddSet(o SetUint16) SetUint16 {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type listUint16 []uint16

func (s listUint16) Len() int           { return len(s) }
func (s listUint16) Less(i, j int) bool { return s[i] < s[j] }
func (s listUint16) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SetUint32 map[uint32]struct{}

func NewSetUint32() SetUint32 {
	return make(map[uint32]struct{})
}

// Diff return new set with all items in s but not in o
func (s SetUint32) Diff(o SetUint32) SetUint32 {
	result := make(SetUint32, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s SetUint32) Inter(o SetUint32) SetUint32 {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(SetUint32, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s SetUint32) Union(o SetUint32) SetUint32 {
	result := make(SetUint32, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s SetUint32) Slice() []uint32 {
	result := make([]uint32, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s SetUint32) SortedSlice() []uint32 {
	slice := listUint32(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s SetUint32) ReversedSortedSlice() []uint32 {
	slice := listUint32(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s SetUint32) Add(id uint32) SetUint32 {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s SetUint32) Del(id uint32) SetUint32 {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s SetUint32) IsExist(id uint32) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s SetUint32) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s SetUint32) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s SetUint32) DelSet(o SetUint32) SetUint32 {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s SetUint32) AddSlice(ids []uint32) SetUint32 {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s SetUint32) AddSet(o SetUint32) SetUint32 {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type listUint32 []uint32

func (s listUint32) Len() int           { return len(s) }
func (s listUint32) Less(i, j int) bool { return s[i] < s[j] }
func (s listUint32) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SetUint64 map[uint64]struct{}

func NewSetUint64() SetUint64 {
	return make(map[uint64]struct{})
}

// Diff return new set with all items in s but not in o
func (s SetUint64) Diff(o SetUint64) SetUint64 {
	result := make(SetUint64, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s SetUint64) Inter(o SetUint64) SetUint64 {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(SetUint64, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s SetUint64) Union(o SetUint64) SetUint64 {
	result := make(SetUint64, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s SetUint64) Slice() []uint64 {
	result := make([]uint64, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s SetUint64) SortedSlice() []uint64 {
	slice := listUint64(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s SetUint64) ReversedSortedSlice() []uint64 {
	slice := listUint64(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s SetUint64) Add(id uint64) SetUint64 {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s SetUint64) Del(id uint64) SetUint64 {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s SetUint64) IsExist(id uint64) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s SetUint64) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s SetUint64) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s SetUint64) DelSet(o SetUint64) SetUint64 {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s SetUint64) AddSlice(ids []uint64) SetUint64 {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s SetUint64) AddSet(o SetUint64) SetUint64 {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type listUint64 []uint64

func (s listUint64) Len() int           { return len(s) }
func (s listUint64) Less(i, j int) bool { return s[i] < s[j] }
func (s listUint64) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SetString map[string]struct{}

func NewSetString() SetString {
	return make(map[string]struct{})
}

// Diff return new set with all items in s but not in o
func (s SetString) Diff(o SetString) SetString {
	result := make(SetString, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s SetString) Inter(o SetString) SetString {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(SetString, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s SetString) Union(o SetString) SetString {
	result := make(SetString, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s SetString) Slice() []string {
	result := make([]string, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s SetString) SortedSlice() []string {
	slice := listString(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s SetString) ReversedSortedSlice() []string {
	slice := listString(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s SetString) Add(id string) SetString {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s SetString) Del(id string) SetString {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s SetString) IsExist(id string) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s SetString) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s SetString) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s SetString) DelSet(o SetString) SetString {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s SetString) AddSlice(ids []string) SetString {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s SetString) AddSet(o SetString) SetString {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type listString []string

func (s listString) Len() int           { return len(s) }
func (s listString) Less(i, j int) bool { return s[i] < s[j] }
func (s listString) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SetFloat32 map[float32]struct{}

func NewSetFloat32() SetFloat32 {
	return make(map[float32]struct{})
}

// Diff return new set with all items in s but not in o
func (s SetFloat32) Diff(o SetFloat32) SetFloat32 {
	result := make(SetFloat32, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s SetFloat32) Inter(o SetFloat32) SetFloat32 {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(SetFloat32, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s SetFloat32) Union(o SetFloat32) SetFloat32 {
	result := make(SetFloat32, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s SetFloat32) Slice() []float32 {
	result := make([]float32, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s SetFloat32) SortedSlice() []float32 {
	slice := listFloat32(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s SetFloat32) ReversedSortedSlice() []float32 {
	slice := listFloat32(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s SetFloat32) Add(id float32) SetFloat32 {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s SetFloat32) Del(id float32) SetFloat32 {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s SetFloat32) IsExist(id float32) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s SetFloat32) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s SetFloat32) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s SetFloat32) DelSet(o SetFloat32) SetFloat32 {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s SetFloat32) AddSlice(ids []float32) SetFloat32 {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s SetFloat32) AddSet(o SetFloat32) SetFloat32 {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type listFloat32 []float32

func (s listFloat32) Len() int           { return len(s) }
func (s listFloat32) Less(i, j int) bool { return s[i] < s[j] }
func (s listFloat32) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SetFloat64 map[float64]struct{}

func NewSetFloat64() SetFloat64 {
	return make(map[float64]struct{})
}

// Diff return new set with all items in s but not in o
func (s SetFloat64) Diff(o SetFloat64) SetFloat64 {
	result := make(SetFloat64, len(s)-len(o))
	for e := range s {
		if _, in := o[e]; !in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Inter return new set with all items in s and in o
func (s SetFloat64) Inter(o SetFloat64) SetFloat64 {
	small := s
	big := o
	if len(s) < len(o) {
		small = o
		big = s
	}
	result := make(SetFloat64, len(small))
	for e := range small {
		if _, in := big[e]; in {
			result[e] = struct{}{}
		}
	}
	return result
}

// Union return new set with all items in s or in o
func (s SetFloat64) Union(o SetFloat64) SetFloat64 {
	result := make(SetFloat64, len(s)+len(o))
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range o {
		result[e] = struct{}{}
	}
	return result
}

// Slice return slice of all items in s
func (s SetFloat64) Slice() []float64 {
	result := make([]float64, len(s))
	idx := 0
	for id := range s {
		result[idx] = id
		idx++
	}
	return result
}

// SortedSlice return sorted slice of all items in s
func (s SetFloat64) SortedSlice() []float64 {
	slice := listFloat64(s.Slice())
	sort.Sort(slice)
	return slice
}

// ReversedSortedSlice return reversed sorted slice of all items in s
func (s SetFloat64) ReversedSortedSlice() []float64 {
	slice := listFloat64(s.Slice())
	sort.Sort(sort.Reverse(slice))
	return slice
}

// Add add item to set
func (s SetFloat64) Add(id float64) SetFloat64 {
	s[id] = struct{}{}
	return s
}

// Del remove item from set
func (s SetFloat64) Del(id float64) SetFloat64 {
	delete(s, id)
	return s
}

// IsExist check if item exist in set
func (s SetFloat64) IsExist(id float64) bool {
	_, ok := s[id]
	return ok
}

// Length return length of set
func (s SetFloat64) Length() int {
	return len(s)
}

// IsEmpty check if set is empty
func (s SetFloat64) IsEmpty() bool {
	return s.Length() == 0
}

// DelSet remove all items in o from set
func (s SetFloat64) DelSet(o SetFloat64) SetFloat64 {
	for e := range o {
		delete(s, e)
	}
	return s
}

// AddSlice add slice of items to set
func (s SetFloat64) AddSlice(ids []float64) SetFloat64 {
	for _, id := range ids {
		s[id] = struct{}{}
	}
	return s
}

// AddSet add all items in o to set
func (s SetFloat64) AddSet(o SetFloat64) SetFloat64 {
	for e := range o {
		s[e] = struct{}{}
	}
	return s
}

type listFloat64 []float64

func (s listFloat64) Len() int           { return len(s) }
func (s listFloat64) Less(i, j int) bool { return s[i] < s[j] }
func (s listFloat64) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
