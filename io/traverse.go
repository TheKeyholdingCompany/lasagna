package io

import (
	"archive/zip"
	"io"
	"io/fs"
	"lasagna/helpers"
	"log"
	"os"
	"os/exec"
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
	var file string
	err := filepath.WalkDir(dirPath, func(path string, dir fs.DirEntry, err error) error {
		helpers.CheckError(err)
		if dir.IsDir() || countDepth(dirPath, path) > maxDepth {
			return nil
		}
		if reg.MatchString(path) {
			file = path
			return io.EOF
		}
		return nil
	})
	if err == io.EOF {
		return file, nil
	}
	return file, err
}

func FindDependencies(path string) string {
	log.Printf("Finding dependencies in %s\n", path)
	files, err := FindFile(path, "(requirements.txt|package.json)", 1)
	helpers.CheckError(err)
	return files
}

func SystemZip(input string, output string) {
	exec.Command("zip", "-r", output, input+"/*").Output()
}

// See https://stackoverflow.com/questions/37869793/how-do-i-zip-a-directory-containing-sub-directories-or-files-in-golang
func GoZip(input string, output string) {
	zipFile, err := os.Create(output)
	helpers.CheckError(err)
	defer zipFile.Close()

	z := zip.NewWriter(zipFile)
	defer z.Close()

	err = filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
		helpers.CheckError(err)
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		helpers.CheckError(err)
		defer file.Close()

		zipPath := strings.Replace(path, input, "", 1)
		zipPath = strings.TrimPrefix(zipPath, "/")

		f, err := z.Create(zipPath)
		helpers.CheckError(err)

		_, err = io.Copy(f, file)
		helpers.CheckError(err)

		return nil
	})
	helpers.CheckError(err)
}

func Zip(input string, output string, useSystem bool) {
	if useSystem {
		SystemZip(input, output)
	} else {
		GoZip(input, output)
	}
}
