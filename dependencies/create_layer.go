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

func doInstall(dependencyFile string, directory string, excludes []string, isVerbose bool) ([]byte, error) {
	if strings.HasSuffix(dependencyFile, "requirements.txt") {
		if isVerbose {
			log.Printf("Installing dependencies from %s\n", dependencyFile)
		}

		defaultExcludes := []string{"botocore", "boto3"}
		newRequirementsFile := dependencyFile + ".tmp"
		err := lio.CopyFileExcludingLines(dependencyFile, newRequirementsFile, append(defaultExcludes, excludes...))
		defer os.Remove(newRequirementsFile)
		helpers.CheckError(err)
		result, installErr := exec.Command("pip3",
			"install",
			"-r", newRequirementsFile,
			"-t", directory+"/python").CombinedOutput()

		if isVerbose {
			log.Println(string(result))
		}

		filePath, fileErr := lio.FindPath(directory+"/python", "cryptography", 1, true)
		if filePath != "" && fileErr == nil {
			if isVerbose {
				log.Println("Re-fetching cryptography for manylinux2014_x86_64")
			}
			cryptoInstallResult, cryptoInstallErr := exec.Command("pip3",
				"install",
				"--platform", "manylinux2014_x86_64",
				"-t", directory+"/python",
				"--implementation", "cp",
				"--only-binary=:all:",
				"--upgrade",
				"cryptography==40.0.2").CombinedOutput()
			if isVerbose {
				log.Println(string(cryptoInstallResult))
			}
			helpers.CheckError(cryptoInstallErr)
		}

		return result, installErr
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

func FetchDependencies(dependencyFile string, directory string, excludes []string, isVerbose bool) {
	_, err := doInstall(dependencyFile, directory, excludes, isVerbose)
	helpers.CheckError(err)
}
