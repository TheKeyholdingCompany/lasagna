package dependencies

import (
	"errors"
	"io"
	"lasagna/helpers"
	lio "lasagna/io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func doInstall(dependencyFile string, directory string) ([]byte, error) {
	if strings.HasSuffix(dependencyFile, "requirements.txt") {
		err := lio.CopyFileExcludingLines(dependencyFile, dependencyFile+".tmp", []string{"botocore", "boto3"})
		defer os.Remove(dependencyFile + ".tmp")
		helpers.CheckError(err)
		return exec.Command("pip3", "install", "-r", dependencyFile+".tmp", "-t", directory+"/python").CombinedOutput()
	}
	if strings.HasSuffix(dependencyFile, "package.json") {
		source, err := os.Open(dependencyFile)
		defer source.Close()
		helpers.CheckError(err)

		err = os.MkdirAll(directory+"/nodejs", os.ModePerm)
		helpers.CheckError(err)

		destination, err := os.Create(directory + "/nodejs/package.json")
		defer destination.Close()
		helpers.CheckError(err)

		_, err = io.Copy(destination, source)
		helpers.CheckError(err)

		return exec.Command("npm", "--prefix", directory+"/nodejs", "install").CombinedOutput()
	}
	return nil, errors.New("unknown project type")
}

func FetchDependencies(dependencyFile string, directory string, isVerbose bool) {
	output, err := doInstall(dependencyFile, directory)
	if isVerbose {
		log.Println(string(output))
	}
	helpers.CheckError(err)
}
