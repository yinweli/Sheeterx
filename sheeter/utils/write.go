package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// FileName 取得檔案名稱
func FileName(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

// FileExist 檔案是否存在
func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// WriteFile 寫入檔案, 如果有需要會建立目錄
func WriteFile(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return fmt.Errorf("file write failed: %w", err)
	} // if

	if err := os.WriteFile(path, data, fs.ModePerm); err != nil {
		return fmt.Errorf("file write failed: %w", err)
	} // if

	return nil
}
