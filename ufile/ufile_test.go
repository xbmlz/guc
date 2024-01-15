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
		{"not exist", "../testdata/test.txt", false},
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
	if len(paths) != 2 {
		t.Errorf("ListFiles(): %v", paths)
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
