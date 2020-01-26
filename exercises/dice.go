package exercises

import (
	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// Dice Problem:
// We have 5 dice: 4-sided, 6-sided, 8-sided, 12-sided, and 20-sided
// Select a die at random, roll it, and record the number (potentially repeating the process).
// What is the probability that I rolled each die?

// an observation is the value rolled on a die
type diceObservation struct {
	val float64
}

// Getlikelihood is the likelihood function for the dice problem
func (o *diceObservation) GetLikelihood(hypo float64) float64 {
	if hypo < o.val {
		// value of the roll is greater than the number of sides on hypothesis die; can't happen
		return 0
	}
	return 1 / hypo
}

// Dice runs the dice problem
func Dice() {
	s := prob.NewSuite(
		prob.NewPmfElement(4, 1),
		prob.NewPmfElement(6, 1),
		prob.NewPmfElement(8, 1),
		prob.NewPmfElement(12, 1),
		prob.NewPmfElement(20, 1),
	)

	obs := []prob.SuiteObservation{
		&diceObservation{6},
		&diceObservation{6},
		&diceObservation{8},
		&diceObservation{7},
		&diceObservation{7},
		&diceObservation{5},
		&diceObservation{4},
	}
	s.MultiUpdate(obs)

	s.Print()
}
