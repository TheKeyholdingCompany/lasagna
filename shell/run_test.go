package shell

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	Run("echo 'hello' > ./example.txt")
	fileContent, err := os.ReadFile("./example.txt")
	if err != nil {
		t.Fatalf(`Run("echo 'hello' > ./example.txt") failed to create file.`)
	}
	if !strings.Contains(string(fileContent), "hello") {
		t.Fatalf(`Run("echo 'hello' > ./example.txt") didn't have the expected message.`)
	}
	Run("rm ./example.txt")
	_, err = os.Stat("./example.txt")
	if err == nil || !errors.Is(err, os.ErrNotExist) {
		os.Remove("./example.txt")
		t.Fatalf(`Run("rm ./example.txt") failed to remove the file'.`)
	}
}
