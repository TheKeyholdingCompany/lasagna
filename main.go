package main

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"lasagna/dependencies"
	"lasagna/helpers"
	"lasagna/io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var VERSION = "development"

func main() {
	usage := "lasagna (version: " + VERSION + `)

Usage:
  lasagna [options]
  lasagna --version
  lasagna --help

Options:
  -h --help           Show this Help.
  -o --output=<path>  Path to output file [default: ./layer.zip].
  -e --exclude=<lib>  A library or a comma-separated list of libraries to exclude.
  -z --nix-zip        Use zip, rather than letting lasagna do it (this is faster).
  -v --verbose        Extra output for debugging.

Examples:
  lasagna --output=./my-layer.zip
  lasagna --output=./my-layer.zip -z --exclude=lib1,lib2`

	arguments, _ := docopt.ParseDoc(usage)
	version, _ := arguments.Bool("--version")
	if version {
		fmt.Println("lasagna " + VERSION)
		os.Exit(0)
	}
	output, _ := arguments.String("--output")
	exclude, _ := arguments.String("--exclude")
	useSystemZip, _ := arguments.Bool("--nix-zip")
	isVerbose, _ := arguments.Bool("--verbose")
	absoluteOutput, _ := filepath.Abs(output)
	excludes := strings.Split(exclude, ",")

	cwd, err := os.Getwd()
	helpers.CheckError(err)
	file := io.FindDependencies(cwd)
	if isVerbose {
		log.Println("Found dependencies file: " + file)
	}
	log.Println("Fetching dependencies...")
	dependencies.FetchDependencies(file, absoluteOutput+".tmp", excludes, isVerbose)
	log.Println("Zipping dependencies...")
	io.Zip(output+".tmp", absoluteOutput, useSystemZip)
	os.RemoveAll(absoluteOutput + ".tmp")
	log.Println("Done: " + absoluteOutput + " (" + output + ")")
}
