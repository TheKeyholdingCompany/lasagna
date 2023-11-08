package shell

import (
	"bytes"
	"lasagna/helpers"
	"log"
	"os/exec"
)

func Run(command string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	log.Println(stdout.String())
	stdErr := stdout.String()
	if stdErr != "" {
		log.Println(stdErr)
	}
	helpers.CheckError(err)
}
