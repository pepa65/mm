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
	version = "0.2.3"
	pegs    = 8
	colors  = 10
)

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var secret string
	if len(os.Args) > 2 {
		fmt.Printf("mm v%s - Mastermind\nUsage:  mm [secret]\n", version)
		os.Exit(-1)
	} else if len(os.Args) == 1 {
		s := 0
		n := pegs
		for {
			s += rnd.Intn(colors)
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
