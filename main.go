package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"

	"github.com/pepa65/mm/mastermind"
)

const (
	version    = "0.5.3"
	def_pegs   = 8
	def_colors = 10
)

func main() {
	usage := fmt.Sprintf("mm v%s - Mastermind\nUsage:  mm [PEGS COLORS]\n  PEGS:    Number of positions [default: 8]\n  COLORS:  String of all permissible characters [default: 0123456789]\n", version)
	if len(os.Args) == 2 {
		fmt.Fprintf(os.Stderr, "%s\nAbort: only 1 argument\n", usage)
		os.Exit(1)
	}

	if len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "%s\nAbort: more than 2 argument\n", usage)
		os.Exit(1)
	}

	pegs := def_pegs
	var p uint64
	colors := def_colors
	Colors := "0123456789ABCDEFGKLMNOPQRSTUVXXYZ"[:colors]
	var err error
	if len(os.Args) == 3 {
		p, err = strconv.ParseUint(os.Args[1], 10, 0)
		if err != nil || p < 1 {
			fmt.Fprintf(os.Stderr, "%s\nAbort: 1st argument not numeric or smaller than 1\n", usage)
			os.Exit(1)
		}
		pegs = int(p)
		Colors = os.Args[2]
		colors = len(Colors)
		if colors < 2 {
			fmt.Fprintf(os.Stderr, "%s\nAbort: need at least 2 colors\n", usage)
			os.Exit(1)
		}
		seen := make(map[rune]bool)
		for _, r := range Colors {
			if seen[r] {
			}
			seen[r] = true
    }
		if len(seen) != colors {
			fmt.Fprintf(os.Stderr, "%s\nAbort: duplicate color found\n", usage)
			os.Exit(1)
		}
	}

	var secret string
	// Generate a secret
	n := pegs
	for {
		secret += string(Colors[rand.IntN(colors)])
		n--
		if n < 1 {
			break
		}
	}
	game := mastermind.Game{
		Pegs:    pegs,
		Colors:  Colors,
		Secret:  secret,
		Chooser: &mastermind.RandomCandidateChooser{},
	}
	if err := game.Solve(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\nAbort: %s\n", usage, err)
		os.Exit(1)
	}
	os.Exit(0)
}
