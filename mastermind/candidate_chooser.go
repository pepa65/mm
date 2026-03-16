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
	rand.Seed(time.Now().Unix())
	r := rand.Intn(len(solutionSpace))
	return solutionSpace[r]
}

type premierCandidateChooser struct{}

func (chooser premierCandidateChooser) choose(solutionSpace []string) string {
	return solutionSpace[0]
}
