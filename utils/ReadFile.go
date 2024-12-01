package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func ReadFile(relPath string) (string, error) {
	_, callerFile, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("unable to determine caller")
	}

	absPath := filepath.Join(filepath.Dir(callerFile), relPath)
	content, err := os.ReadFile(absPath)
	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(content), "\n"), nil
}

func Dir() (string, error) {
	_, callerFile, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("unable to determine caller's directory")
	}
	return filepath.Dir(callerFile), nil
}
