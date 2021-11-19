package godash

import (
	"testing"
)

func TestCountPre(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty/substr", args{"", ""}, 0},
		{"empty/substr", args{"x", ""}, 1},
		{"empty/substr", args{"xx", ""}, 2},
		{"empty/substr", args{"xxx中国字", ""}, 7},

		{"empty/1-substr", args{"", "x"}, 0},
		{"empty/substr", args{"xy", "x"}, 1},
		{"empty/substr", args{"xxy", "x"}, 2},
		{"empty/substr", args{"中中中y", "中"}, 7},

		{"empty/1-substr", args{"", "x"}, 0},
		{"empty/substr", args{"xyxy", "x"}, 1},
		{"empty/substr", args{"xxyxy", "x"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountPre(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("CountPre() = %v, want %v", got, tt.want)
			}
		})
	}
}
