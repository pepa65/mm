package mastermind

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateSecretEmptySecret(t *testing.T) {
	var game = Game{
		Pegs:   3,
		Colors: "0123",
		Secret: "",
	}
	var err = game.validateSecret()
	assert.Equal(t, err, fmt.Errorf("The length of the secret should be 3"))
}

func TestValidateSecretLongerThanPegs(t *testing.T) {
	var game = Game{
		Pegs:   3,
		Colors: "0123",
		Secret: "0123",
	}
	var err = game.validateSecret()
	assert.Equal(t, err, fmt.Errorf("The length of the secret should be 3"))
}

func TestValidateSecretContainsIllegalSymbol(t *testing.T) {
	var game = Game{
		Pegs:   4,
		Colors: "0123",
		Secret: "1234",
	}
	var err = game.validateSecret()
	assert.Equal(t, err, fmt.Errorf("The secret contains invalid symbols"))
}

func TestValidateSecretNoError(t *testing.T) {
	var game = Game{
		Pegs:   4,
		Colors: "0123",
		Secret: "0123",
	}
	var err = game.validateSecret()
	assert.Nil(t, err)
}

func TestValidateGuessAllCorrect(t *testing.T) {
	var game = Game{
		Pegs:   4,
		Colors: "0123",
		Secret: "0123",
	}
	assert.Equal(t, game.validateGuess("0123"), Result{4, 0})
}

func TestValidateGuessNoneCorrect(t *testing.T) {
	var game = Game{
		Pegs:   4,
		Colors: "0123456789",
		Secret: "0123",
	}
	assert.Equal(t, game.validateGuess("5678"), Result{0, 0})
}

func TestValidateGuessAllColorsCorrect(t *testing.T) {
	var game = Game{
		Pegs:   4,
		Colors: "0123456789",
		Secret: "0123",
	}
	assert.Equal(t, game.validateGuess("3210"), Result{0, 4})
}

func TestValidateGuessRepeatedColors(t *testing.T) {
	var game = Game{
		Pegs:   4,
		Colors: "0123456789",
		Secret: "0123",
	}
	assert.Equal(t, game.validateGuess("2233"), Result{1, 3})
}

func TestGenerateSolutionSpace(t *testing.T) {
	var game = Game{
		Pegs:   2,
		Colors: "012",
	}
	assert.Equal(
		t, game.generateSolutionSpace(),
		[]string{"00", "10", "20", "01", "11", "21", "02", "12", "22"},
	)
}

func TestGenerateInitialGuess(t *testing.T) {
	var game = Game{
		Pegs:   9,
		Colors: "0123456789",
	}
	assert.Equal(t, game.generateInitialGuess(), "001122334")
}

func TestGenerateInitialGuessFourPegs(t *testing.T) {
	var game = Game{
		Pegs:   4,
		Colors: "012",
	}
	assert.Equal(t, game.generateInitialGuess(), "0011")
}

func TestGenerateInitialGuessThreePegs(t *testing.T) {
	var game = Game{
		Pegs:   3,
		Colors: "012",
	}
	assert.Equal(t, game.generateInitialGuess(), "001")
}
