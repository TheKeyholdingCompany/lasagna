package io

import (
	"archive/zip"
	"io"
	"lasagna/helpers"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func SystemZip(input string, output string) {
	cmd := exec.Command("zip", "-r", output, input)
	_, err := cmd.Output()
	helpers.CheckError(err)
}

// GoZip See https://stackoverflow.com/questions/37869793/how-do-i-zip-a-directory-containing-sub-directories-or-files-in-golang
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

func Zip(input string, absoluteOutput string, useSystem bool) {
	originalDir, _ := os.Getwd()
	os.Chdir(input)
	defer os.Chdir(originalDir)

	if useSystem {
		SystemZip(".", absoluteOutput)
	} else {
		GoZip(".", absoluteOutput)
	}
}
