package uslice

import "testing"

func TestIsExist(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		v    int
		want bool
	}{
		{"exist", []int{1, 2}, 1, true},
		{"not exist", []int{1, 2}, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExist(tt.s, tt.v); got != tt.want {
				t.Errorf("IsExist(): %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
