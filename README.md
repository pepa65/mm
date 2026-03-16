[![Go Report Card](https://goreportcard.com/badge/github.com/pepa65/mastermind)](https://goreportcard.com/report/github.com/pepa65/mastermind)
[![GoDoc](https://godoc.org/github.com/pepa65/mastermind?status.svg)](https://godoc.org/github.com/pepa65/mastermind)
[![GitHub](https://img.shields.io/github/license/pepa65/mastermind.svg)](LICENSE)
[![run-ci](https://github.com/pepa65/mastermind/actions/workflows/ci.yml/badge.svg)](https://github.com/pepa65/mastermind/actions/workflows/ci.yml)

# mm
**Mastermind**

A [Golang](http://golang.org/) implementation of the solution to the [MasterMind](http://en.wikipedia.org/wiki/Mastermind_%28board_game%29) game.

* Repo: https://github.com/pepa65/mastermind
* After: https://github.com/kevinjqiu/mastermind.git

## Usage
`mm [SECRET]`

The parameters `pegs` and `colors` can be adjusted in `main.go` for different game configurations.

## Install
### Build locally (needs Golang install)
```
git clone https://github.com/pepa65/mastermind
cd mastermind
CGO_ENABLED=0 go install -ldflags="-s -w" mm  # Flags for smaller binary
upx --best --lzma mm  # Compress the binary for smaller size
mv mm ~/bin/  # Assuming ~/bin is in PATH

### Install from remote repo
`go install github.com/pepa65/mastermind@latest`

## Build
`go build -a -o mm`

## Test
`go test -a -v ./...`
