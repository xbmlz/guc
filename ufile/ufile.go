package ufile

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/xbmlz/guc/uslice"
)

// CopyDir copies a directory from src to dst.
func CopyDir(src, dst string) error {
	files, err := ListFiles(src, nil, true)
	if err != nil {
		return err
	}
	for _, file := range files {
		dstFile := path.Join(dst, file[len(src):])
		dstDir := path.Dir(dstFile)
		if !IsExist(dstDir) {
			if err := os.MkdirAll(dstDir, 0755); err != nil {
				return err
			}
		}
		if err := CopyFile(file, dstFile); err != nil {
			return err
		}
	}
	return nil
}

// CopyFile copies a file from src to dst.
func CopyFile(src, dst string) error {
	// Open original file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create new file, if dir not exist, create it
	if !IsExist(path.Dir(dst)) {
		if err := os.MkdirAll(path.Dir(dst), 0755); err != nil {
			return err
		}
	}
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)

	return err
}

// IsDir checks if a path is a directory.
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsExist checks if a path exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ListFiles lists all files in a directory.
func ListFiles(dir string, exts []string, recursive bool) ([]string, error) {
	var paths []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		fullpath := path.Join(dir, entry.Name())
		if entry.IsDir() && recursive {
			p, _ := ListFiles(fullpath, exts, recursive)
			paths = append(paths, p...)
		} else {
			if len(exts) == 0 {
				paths = append(paths, fullpath)
			} else {
				if uslice.IsExist[string](exts, filepath.Ext(fullpath)) {
					paths = append(paths, fullpath)
				}
			}
		}
	}
	return paths, nil
}

// GetMimeType returns the mime type of a file.
func GetMimeType(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	mimeType := http.DetectContentType(buffer)
	return mimeType, nil
}

// FormatSize formats a file size to a human readable string.
func FormatSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2f B", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2f KB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f MB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f GB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f TB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2f EB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

// Write content to file
func Write(path, content string, append bool) (err error) {
	mode := os.O_WRONLY | os.O_CREATE
	if append {
		mode = mode | os.O_APPEND
	} else {
		mode = mode | os.O_TRUNC
	}
	file, err := os.OpenFile(path, mode, 0o666)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(content); err != nil {
		return err
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	return nil
}

// Read content from file
func Read(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
