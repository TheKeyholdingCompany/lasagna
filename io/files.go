package io

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func lineHasContainsWords(line string, words []string) bool {
	for _, word := range words {
		if strings.Contains(line, word) {
			return true
		}
	}
	return false
}

func GrepFile(sourcePath string, token string) (string, error) {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return "", err
	}
	defer sourceFile.Close()

	scanner := bufio.NewScanner(sourceFile)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, token) {
			return line, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", errors.New("token not found")
}

func CopyFileExcludingLines(sourcePath string, targetPath string, excludes []string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	targetFile, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	scanner := bufio.NewScanner(sourceFile)
	writer := bufio.NewWriter(targetFile)

	for scanner.Scan() {
		line := scanner.Text()
		if !lineHasContainsWords(line, excludes) {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return writer.Flush()
}
