package mastermind

import (
	"fmt"
	"strings"
)

// A Result is a pair of integers indicating:
// - the number of correct symbols and positions
// - the number of correct symbols (but wrong position)
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
		return fmt.Errorf("The length of the secret should be %d", game.Pegs)
	}

	for _, s := range game.Secret {
		if !strings.ContainsRune(game.Colors, s) {
			return fmt.Errorf("The secret contains invalid symbols")
		}
	}
	return nil
}

func (game *Game) generateInitialGuess() string {
	/*var guess []rune
	for i := 0; i < (game.Pegs+1)/2; i++ {
		guess = append(guess, rune(game.Colors[0]))
	}
	for i := 0; i < game.Pegs/2; i++ {
		guess = append(guess, rune(game.Colors[1]))
	}
	return string(guess)*/
	return "00112233445566778899AABBCCDDEEFFGGHH"[:game.Pegs]
}

func (game *Game) generateSolutionSpace() []string {
	sets := make([]string, game.Pegs)
	for i := 0; i < game.Pegs; i++ {
		sets[i] = game.Colors
	}
	return cartesianProduct(sets)
}

func (game *Game) validateGuess(guess string) Result {
	return validateGuess(game.Secret, guess)
}

func (game *Game) Solve() error {
	if err := game.validateSecret(); err != nil {
		return err
	}

	var (
		result     Result
		numGuesses int
	)
	solutionSpace := game.generateSolutionSpace()
	guess := game.generateInitialGuess()
	for {
		result = game.validateGuess(guess)
		numGuesses++
		fmt.Printf("%2d: guess %s %-*s [from %d]\n", numGuesses, guess, game.Pegs, result.ToString(), len(solutionSpace))
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

func validateGuess(secret string, guess string) Result {
	var (
		correctPositions int
		correctColors    int
	)
	for i, g := range guess {
		s := rune(secret[i])
		if g == s {
			correctPositions++
		} else {
			if strings.ContainsRune(secret, g) {
				correctColors++
			}
		}
	}
	return Result{correctPositions, correctColors}
}

func eliminateSolutionSpace(solutionSpace []string, result Result, guess string) []string {
	retval := []string{}
	for _, candidate := range solutionSpace {
		if validateGuess(candidate, guess) == result {
			retval = append(retval, candidate)
		}
	}
	return retval
}
