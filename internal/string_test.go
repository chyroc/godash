package internal

import (
	"testing"
)

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"", "", ""},
		{"a", "a", "A"},
		{"A", "A", "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelCase(tt.args); got != tt.want {
				t.Errorf("ToCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
