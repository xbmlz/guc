package ustring

import "testing"

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"empty", "", true},
		{"space", " ", false},
		{"space and tab", " \t\n", false},
		{"not empty", "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.s); got != tt.want {
				t.Errorf("IsEmpty(): %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"empty", "", true},
		{"space", " ", true},
		{"space and tab", " \t\n", true},
		{"not empty", "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlank(tt.s); got != tt.want {
				t.Errorf("IsBlank(): %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
