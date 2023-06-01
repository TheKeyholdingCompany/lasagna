package dependencies

import (
	"errors"
	"io"
	"lasagna/helpers"
	lio "lasagna/io"
	"os"
	"os/exec"
	"strings"
)

func doInstall(dependencyFile string, directory string, platform string) ([]byte, error) {
	if strings.HasSuffix(dependencyFile, "requirements.txt") {
		_platform := platform
		if platform == "" {
			_platform = "manylinux1_x86_64"
		}
		err := lio.CopyFileExcludingLines(dependencyFile, dependencyFile+".tmp", []string{"botocore", "boto3"})
		defer os.Remove(dependencyFile + ".tmp")
		helpers.CheckError(err)
		return exec.Command("pip3", "install", "-r", dependencyFile+".tmp", "-t", directory+"/python", "--platform", _platform).Output()
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

		return exec.Command("npm", "--prefix", directory+"/nodejs", "install").Output()
	}
	return nil, errors.New("unknown project type")
}

func FetchDependencies(dependencyFile string, directory string, platform string) {
	_, err := doInstall(dependencyFile, directory, platform)
	//log.Println(string(output))
	helpers.CheckError(err)
}
