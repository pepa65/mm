[![Go Report Card](https://goreportcard.com/badge/github.com/pepa65/mm)](https://goreportcard.com/report/github.com/pepa65/mm)
[![GoDoc](https://godoc.org/github.com/pepa65/mm?status.svg)](https://godoc.org/github.com/pepa65/mm)
[![GitHub](https://img.shields.io/github/license/pepa65/mm.svg)](LICENSE)
[![run-ci](https://github.com/pepa65/mm/actions/workflows/ci.yml/badge.svg)](https://github.com/pepa65/mm/actions/workflows/ci.yml)

# mm v0.5.3
**Mastermind**

Solving [MasterMind](http://en.wikipedia.org/wiki/Mastermind_%28board_game%29) games, a [Golang](http://golang.org/) implementation.

* Repo: https://github.com/pepa65/mm
* After: https://github.com/kevinjqiu/mastermind

## Usage
```
mm v0.5.3 - Mastermind
Usage:  mm [PEGS COLORS]
  PEGS:    Number of positions [default: 8]
  COLORS:  String of all permissible characters [default: 0123456789]
```

## Install
### Build locally (needs Golang install)
```
git clone https://github.com/pepa65/mm
cd mm
CGO_ENABLED=0 go install -ldflags="-s -w"  # Flags for smaller binary
upx --best --lzma mm  # Compress the binary for smaller size
mv mm ~/bin/  # Assuming ~/bin is in PATH

### Install from remote repo
`go install github.com/pepa65/mm@latest`

### Download
```
wget -O mm github.com/pepa65/mm/releases/download/v0.5.3/mm_0.5.3_linux_amd64
chmod +x mm
mv mm ~/bin/  # Provided ~/bin is in $PATH
```

## Build
`go build -a`

## Test
`go test -a -v ./...`
