package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/pepa65/mastermind/mastermind"
)

const (
	version = "0.2.1"
	pegs   = 8
	colors = 10
)

func main() {
	var secret string
	if len(os.Args) > 2 {
		fmt.Printf("Usage: mastermind [secret]\n")
		os.Exit(-1)
	} else if len(os.Args) == 1 {
		rand.Seed(time.Now().UnixNano())
		s := 0
		n := pegs
		for {
			s += rand.Intn(colors)
			if n < 1 {
				break
			}
			n--
			s *= 10
		}
		secret = strconv.Itoa(s)
	} else {
		secret = os.Args[1]
	}
	fmt.Printf("Secret:   %s\n", secret)
	game := mastermind.Game{
		Pegs:    pegs,
		Colors:  "0123456789ABCDEFGKLMNOPQRSTUVXXYZ"[:colors],
		Secret:  secret,
		Chooser: &mastermind.RandomCandidateChooser{},
	}
	if err := game.Solve(); err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}
	os.Exit(0)
}
