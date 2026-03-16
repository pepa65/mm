package mastermind

import (
	"math/rand"
	"time"
)

// This is the strategy that governs how the candidate is chosen given a solution space
type candidateChooser interface {
	choose(solutionSpace []string) string
}

type RandomCandidateChooser struct{}

func (chooser RandomCandidateChooser) choose(solutionSpace []string) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := rnd.Intn(len(solutionSpace))
	return solutionSpace[r]
}
