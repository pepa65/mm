package mastermind

import (
	"fmt"
	"strings"
)

// Result: pair of integers indicating: number of Well-placed and number of Misplaced pegs
type Result [2]int

func (r Result) ToString() string {
	//return fmt.Sprintf("(%d,%d)", r[0], r[1])
	return strings.Repeat("+", r[0]) + strings.Repeat("-", r[1])
}

// This is the structure representing a mastermind game
type Game struct {
	Pegs    int
	Colors  string
	Secret  string
	Chooser candidateChooser
}

func (game *Game) validateSecret() error {
	if len(game.Secret) != game.Pegs {
		return fmt.Errorf("the length of the secret should be %d", game.Pegs)
	}

	for _, s := range game.Secret {
		if !strings.ContainsRune(game.Colors, s) {
			return fmt.Errorf("the secret contains invalid symbols")
		}
	}
	return nil
}

func (game *Game) generateInitialGuess() string {
	var guess []rune
	for _, r := range game.Colors {
		guess = append(guess, r)
		if len(guess) == game.Pegs {
			break
		}
		guess = append(guess, r)
		if len(guess) == game.Pegs {
			break
		}
	}
	for range game.Pegs - len(guess) {
		guess = append(guess, rune(game.Colors[0]))
	}
	return string(guess)
}

func (game *Game) generateSolutionSpace() []string {
	sets := make([]string, game.Pegs)
	for i := 0; i < game.Pegs; i++ {
		sets[i] = game.Colors
	}
	return cartesianProduct(sets)
}

func (game *Game) scoreGuess(guess string) Result {
	return scoreGuess(game.Secret, guess)
}

func (game *Game) Solve() error {
	if err := game.validateSecret(); err != nil {
		return err
	}

	var (
		result     Result
		numGuesses int
	)
	fmt.Printf("   Secret: %s\n", game.Secret)
	solutionSpace := game.generateSolutionSpace()
	guess := game.generateInitialGuess()
	for {
		result = game.scoreGuess(guess)
		numGuesses++
		fmt.Printf("%2d: guess: %s eval: %-*s [from: %d]\n", numGuesses, guess, game.Pegs, result.ToString(), len(solutionSpace))
		if result[0] == game.Pegs {
			return nil
		}

		solutionSpace = eliminateSolutionSpace(solutionSpace, result, guess)
		if len(solutionSpace) > 0 {
			guess = game.Chooser.choose(solutionSpace)
		} else {
			panic("No candidate solution left.\n")
		}
	}
}

func scoreGuess(secret string, guess string) Result {
	correctPos := 0
	wrongPos := 0
	var counts [256]int
	for i, g := range guess {
		s := rune(secret[i])
		if g == s {
			correctPos++
		} else {
			if counts[s] < 0 {
				wrongPos++
			}
			if counts[g] > 0 {
				wrongPos++
			}
			counts[s]++
			counts[g]--
		}
	}
	return Result{correctPos, wrongPos}
}

func eliminateSolutionSpace(solutionSpace []string, result Result, guess string) []string {
	retval := []string{}
	for _, candidate := range solutionSpace {
		if scoreGuess(candidate, guess) == result {
			retval = append(retval, candidate)
		}
	}
	return retval
}
