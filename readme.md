<div align="center"><img src="./icon.svg" /></div>
<h1 align="center">Lasagna</h1>

A CLI tool to build your lambda layer zip for you.

# Usage
```
Usage:
  lasagna [options]
  lasagna --version
  lasagna --help

Options:
  -h --help               Show this Help.
  -o --output=<path>      Path to output file [default: ./layer.zip].
  -e --exclude=<lib>      A library or a comma-separated list of libraries to exclude.
  -z --nix-zip            Use zip, rather than letting lasagna do it (this is faster).
  -v --verbose            Extra output for debugging.
  -c --command=<command>  Command to run before packaging the layer.
  -r --replace=<library>  Comma-separated list of libraries with a specific version and platform. These will replace the default libraries.
  -h --host=<host>        The package repository host URL (include username and password if needed).

Examples:
  lasagna --output=./my-layer.zip
  lasagna --output=./my-layer.zip -z --host=https://user:password@my.pip.host.com/respository/my-pypi-all/simple
  lasagna --output=./my-layer.zip -z --exclude=lib1,lib2
  lasagna --output=./my-layer.zip -z --replace=lib1:1.0.0:platform,lib2:2.0.0:platform
```

# Installation
## MacOS
```shell
brew tap TheKeyholdingCompany/tap
brew install lasagna
```

## Linux
```shell
snap install lasagna
```

## Other
Download the binary from [releases](https://github.com/TheKeyholdingCompany/lasagna/releases).
