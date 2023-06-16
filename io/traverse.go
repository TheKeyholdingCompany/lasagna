package io

import (
	"io"
	"io/fs"
	"lasagna/helpers"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func countDepth(baseDir string, currentPath string) int {
	return strings.Count(currentPath, string(os.PathSeparator)) - strings.Count(baseDir, string(os.PathSeparator))
}

func FindFile(dirPath string, pattern string, maxDepth int) (string, error) {
	reg, e := regexp.Compile(pattern)
	helpers.CheckError(e)
	var filePath string
	err := filepath.WalkDir(dirPath, func(path string, dir fs.DirEntry, err error) error {
		helpers.CheckError(err)
		if dir.IsDir() || countDepth(dirPath, path) > maxDepth {
			return nil
		}
		if reg.MatchString(path) {
			filePath = path
			return io.EOF
		}
		return nil
	})
	if err == io.EOF {
		return filePath, nil
	}
	return filePath, err
}

func FindDependencies(path string) string {
	log.Printf("Finding dependencies in %s\n", path)
	file, err := FindFile(path, "(requirements.txt|package.json)", 1)
	helpers.CheckError(err)
	return file
}
