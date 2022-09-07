package utils

import (
	"path/filepath"
	"strings"
)

// Removes filename extension and returns only filename
func RemoveExtension(file string) string {
	ext := filepath.Ext(file)
	return strings.TrimSuffix(file, ext)
}
