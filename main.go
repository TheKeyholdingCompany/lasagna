package main

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"lasagna/dependencies"
	"lasagna/helpers"
	"lasagna/io"
	"log"
	"os"
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
  -z --nix-zip        Use zip, rather than letting lasagna do it (this is faster).

Examples:
  lasagna --output=./my-layer.zip`

	arguments, _ := docopt.ParseDoc(usage)
	version, _ := arguments.Bool("--version")
	if version {
		fmt.Println("lasagna " + VERSION)
		os.Exit(0)
	}
	output, _ := arguments.String("--output")
	useSystemZip, _ := arguments.Bool("--nix-zip")

	cwd, err := os.Getwd()
	helpers.CheckError(err)
	file := io.FindDependencies(cwd)

	log.Println("Fetching dependencies...")
	dependencies.FetchDependencies(file, output+".tmp")
	log.Println("Zipping dependencies...")
	io.Zip(output+".tmp", output, useSystemZip)
	os.RemoveAll(output + ".tmp")
	log.Println("Done")
}
