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

func doPythonLibraryReplacements(libraryReplacements LibrarySystemInformation, directory string, isVerbose bool) {
	if isVerbose {
		log.Printf("Checking for %s...\n", libraryReplacements.Name)
	}
	filePath, fileErr := lio.FindPath(directory+"/python", libraryReplacements.Name, 1, true)
	if filePath != "" && fileErr == nil {
		if isVerbose {
			log.Printf("Re-fetching %s for %s\n", libraryReplacements.Name, libraryReplacements.Platform)
		}
		installResult, installErr := exec.Command("pip3",
			"install",
			"--platform", libraryReplacements.Platform,
			"-t", directory+"/python",
			"--implementation", "cp",
			"--only-binary=:all:",
			"--upgrade",
			libraryReplacements.Name+"=="+libraryReplacements.Version).CombinedOutput()
		if isVerbose {
			log.Println(string(installResult))
		}
		helpers.CheckError(installErr)
	}
}

func doPythonInstall(dependencyFile string, directory string, libraryReplacements []LibrarySystemInformation, excludes []string, host string, isVerbose bool) ([]byte, error) {
	if isVerbose {
		ddd, _ := exec.Command("which",
			"pip3").CombinedOutput()
		log.Printf("Pip3 location: %s\n", string(ddd))
		log.Printf("Installing dependencies from %s\n", dependencyFile)
		if host != "" {
			log.Printf("Using package repo URL: %s\n", host)
		}
	}

	defaultExcludes := []string{"botocore", "boto3"}
	log.Println("Excluding: ", append(defaultExcludes, excludes...))
	err := lio.CopyFileExcludingLines(dependencyFile, dependencyFile+".tmp", append(defaultExcludes, excludes...))
	defer os.Remove(dependencyFile + ".tmp")
	helpers.CheckError(err)
	if host != "" {
		err = os.Setenv("PIP_INDEX_URL", host)
		helpers.CheckError(err)
	}
	result, installErr := exec.Command("pip3",
		"install",
		"-r", dependencyFile+".tmp",
		"-t", directory+"/python").CombinedOutput()
	if isVerbose {
		log.Println(string(result))
	}

	for _, libraryReplacement := range libraryReplacements {
		doPythonLibraryReplacements(libraryReplacement, directory, isVerbose)
	}

	if isVerbose {
		log.Println("Finished fetching dependencies")
	}

	return result, installErr
}

func doNodeInstall(dependencyFile string, directory string, libraryReplacements []LibrarySystemInformation, excludes []string, host string, isVerbose bool) ([]byte, error) {
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

func doInstall(dependencyFile string, directory string, libraryReplacements []LibrarySystemInformation, excludes []string, host string, isVerbose bool) ([]byte, error) {
	if strings.HasSuffix(dependencyFile, "requirements.txt") {
		return doPythonInstall(dependencyFile, directory, libraryReplacements, excludes, host, isVerbose)
	}
	if strings.HasSuffix(dependencyFile, "package.json") {
		return doNodeInstall(dependencyFile, directory, libraryReplacements, excludes, host, isVerbose)
	}
	return nil, errors.New("unknown project type")
}

func FetchDependencies(dependencyFile string, directory string, libraryReplacements []LibrarySystemInformation, excludes []string, host string, isVerbose bool) {
	_, err := doInstall(dependencyFile, directory, libraryReplacements, excludes, host, isVerbose)
	helpers.CheckError(err)
}
