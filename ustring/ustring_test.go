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
				t.Errorf("IsEmpty(): name %v , got %v, want %v", tt.name, got, tt.want)
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
				t.Errorf("IsBlank(): name %v , got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestHasEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		want bool
	}{
		{"not empty and empty", []string{"abc", ""}, true},
		{"not empty", []string{"abc", "123"}, false},
		{"empty and tab", []string{" ", "\t", "\n"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasEmpty(tt.s...); got != tt.want {
				t.Errorf("HasEmpty(): name %v , got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestHasBlank(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		want bool
	}{
		{"not empty and empty", []string{"abc", ""}, true},
		{"not empty", []string{"abc", "123"}, false},
		{"empty and tab", []string{" ", "\t", "\n"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasBlank(tt.s...); got != tt.want {
				t.Errorf("HasBlank(): name %v , got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestRemovePrefix(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want string
	}{
		{"empty", "", "", ""},
		{"empty prefix", "abc", "", "abc"},
		{"empty string", "", "abc", ""},
		{"prefix not match", "abc", "xyz", "abc"},
		{"prefix match", "abc", "ab", "c"},
		{"prefix match", "abc", "c", "abc"},
		{"prefix match", "1.jpg", "1.", "jpg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemovePrefix(tt.s, tt.p); got != tt.want {
				t.Errorf("RemovePrefix(): name %v , got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestRemoveSuffix(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want string
	}{
		{"empty", "", "", ""},
		{"empty suffix", "abc", "", "abc"},
		{"empty string", "", "abc", ""},
		{"suffix not match", "abc", "xyz", "abc"},
		{"suffix match", "abc", "bc", "a"},
		{"suffix match", "abc", "a", "abc"},
		{"suffix match", "1.jpg", ".jpg", "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveSuffix(tt.s, tt.p); got != tt.want {
				t.Errorf("RemoveSuffix(): name %v , got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		name string
		s    string
		args []interface{}
		want string
	}{
		{"empty", "", nil, ""},
		{"empty args", "abc", nil, "abc"},
		{"args not match", "abc", []interface{}{"xyz"}, "abc"},
		{"args match", "abc{}", []interface{}{"d"}, "abcd"},
		{"args match", "abc{}{}fg", []interface{}{"d", "e"}, "abcdefg"},
		{"args match", "abc{}123{}ABC{}", []interface{}{"d", "4", "D"}, "abcd1234ABCD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.s, tt.args...); got != tt.want {
				t.Errorf("Format(): name %v , got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"lower case", "cat", "Cat"},
		{"upper case", "cAt", "CAt"},
		{"mixed case", "'cat'", "'cat'"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Capitalize(tt.s); got != tt.want {
				t.Errorf("Capitalize(): name %v , got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestChomp(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"new line", "abc \r", "abc "},
		{"new line", "abc\n", "abc"},
		{"new line", "abc\r\n", "abc"},
		{"new line", "abc\r\n\r\n", "abc\r\n"},
		{"new line", "abc\n\r", "abc\n"},
		{"new line", "你好\n\r", "你好\n"},
		{"new line", "\r", ""},
		{"new line", "\n", ""},
		{"new line", "\r\n", ""},
	}
	for ti, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chomp(tt.s); got != tt.want {
				t.Errorf("Chomp(): %v-%d = %v, want %v", tt.name, ti, got, tt.want)
			}
		})
	}
}

func TestChop(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"new line", "abc \r", "abc "},
		{"new line", "abc\n", "abc"},
		{"new line", "abc\r\n", "abc"},
		{"new line", "abc", "ab"},
		{"new line", "你好", "你"},
		{"new line", "abc\nabc", "abc\nab"},
		{"new line", "a", ""},
		{"new line", "\r", ""},
		{"new line", "\n", ""},
		{"new line", "\r\n", ""},
	}
	for ti, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chop(tt.s); got != tt.want {
				t.Errorf("Chop(): %v-%d = %v, want %v", tt.name, ti, got, tt.want)
			}
		})
	}
}
