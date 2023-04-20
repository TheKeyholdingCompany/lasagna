package dependencies

import (
	"errors"
	"io"
	"lasagna/helpers"
	"os"
	"os/exec"
	"strings"
)

func doInstall(dependencyFile string, directory string) ([]byte, error) {
	if strings.HasSuffix(dependencyFile, "requirements.txt") {
		return exec.Command("pip3", "install", "-r", "requirements.txt", "-t", directory+"/python").Output()
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

func FetchDependencies(dependencyFile string, directory string) {
	_, err := doInstall(dependencyFile, directory)
	//log.Println(string(output))
	helpers.CheckError(err)
}
