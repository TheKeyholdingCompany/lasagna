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
  -h --help           Show this Help.
  -o --output=<path>  Path to output file [default: ./layer.zip].
  -z --nix-zip        Use zip, rather than letting lasagna do it (this is faster).

Examples:
  lasagna --output=./my-layer.zip
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
