package ufile

import (
	"fmt"
	"testing"
)

func TestCopyDir(t *testing.T) {
	src := "../testdata"
	dst := "../testdata_copy"
	err := CopyDir(src, dst)
	if err != nil {
		t.Errorf("CopyDir(): %v", err)
	}
}

func TestCopyFile(t *testing.T) {
	src := "../testdata/test.txt"
	dst := "../testdata/test_copy.txt"
	err := CopyFile(src, dst)
	if err != nil {
		t.Errorf("CopyFile(): %v", err)
	}
}

func TestIsDir(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{"dir", "../testdata", true},
		{"file", "../testdata/test.txt", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDir(tt.path); got != tt.want {
				t.Errorf("IsDir(): %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestIsExist(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{"exist", "../testdata", true},
		{"not exist", "../testdata/test1.txt", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExist(tt.path); got != tt.want {
				t.Errorf("IsExist(): %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestLisFiles(t *testing.T) {
	paths, err := ListFiles("../testdata/", []string{".txt"}, true)
	if err != nil {
		t.Errorf("ListFiles(): %v", err)
	}
	fmt.Println(paths)
}

func TestGetMimeType(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"txt", "../testdata/test.jpeg", "image/jpeg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := GetMimeType(tt.path); got != tt.want || err != nil {
				t.Errorf("GetMimeType(): %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name string
		size int64
		want string
	}{
		{"B", 1, "1.00 B"},
		{"KB", 1024, "1.00 KB"},
		{"MB", 1024 * 1024, "1.00 MB"},
		{"GB", 1024 * 1024 * 1024, "1.00 GB"},
		{"TB", 1024 * 1024 * 1024 * 1024, "1.00 TB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatSize(tt.size); got != tt.want {
				t.Errorf("FormatSize(): %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestWrite(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"txt", "../testdata/test.txt", "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(tt.path, tt.want, true); err != nil {
				t.Errorf("Write(): %v = %v, want %v", tt.name, err, nil)
			}
		})
	}
}

func TestRead(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"txt", "../testdata/test.txt", "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Read(tt.path); got != tt.want || err != nil {
				t.Errorf("Read(): %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
