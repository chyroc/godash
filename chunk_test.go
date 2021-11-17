package godash

import (
	"reflect"
	"testing"
)

func TestChunkInt(t *testing.T) {
	type args struct {
		array []int
		size  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// size=-1
		{"-1", args{[]int{}, -1}, [][]int{}},
		{"-1", args{[]int{1}, -1}, [][]int{}},
		{"-1", args{[]int{1, 2}, -1}, [][]int{}},
		// size=0
		{"0", args{[]int{}, 0}, [][]int{}},
		{"0", args{[]int{1}, 0}, [][]int{}},
		{"0", args{[]int{1, 2}, 0}, [][]int{}},
		// size=1
		{"1", args{[]int{}, 1}, [][]int{}},
		{"1", args{[]int{1}, 1}, [][]int{{1}}},
		{"1", args{[]int{1, 2}, 1}, [][]int{{1}, {2}}},
		// size=2
		{"2", args{[]int{}, 2}, [][]int{}},
		{"2", args{[]int{1}, 2}, [][]int{{1}}},
		{"2", args{[]int{1, 2}, 2}, [][]int{{1, 2}}},
		{"2", args{[]int{1, 2, 3}, 2}, [][]int{{1, 2}, {3}}},
		{"2", args{[]int{1, 2, 3, 4}, 2}, [][]int{{1, 2}, {3, 4}}},
		{"2", args{[]int{1, 2, 3, 4, 5}, 2}, [][]int{{1, 2}, {3, 4}, {5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChunkInt(tt.args.array, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChunkInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
