package exercises

import (
	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// Dice Problem:
// We have 5 dice: 4-sided, 6-sided, 8-sided, 12-sided, and 20-sided
// Select a die at random, roll it, and record the number (potentially repeating the process).
// What is the probability that I rolled each die?

type diceHypothesis struct {
	hypo  *prob.PmfElement
	sides int16
}

var (
	sides4 = diceHypothesis{
		hypo:  prob.NewPmfElement("4-sided", 1),
		sides: 4,
	}
	sides6 = diceHypothesis{
		hypo:  prob.NewPmfElement("6-sided", 1),
		sides: 6,
	}
	sides8 = diceHypothesis{
		hypo:  prob.NewPmfElement("8-sided", 1),
		sides: 8,
	}
	sides12 = diceHypothesis{
		hypo:  prob.NewPmfElement("12-sided", 1),
		sides: 12,
	}
	sides20 = diceHypothesis{
		hypo:  prob.NewPmfElement("20-sided", 1),
		sides: 20,
	}
)

var diceHypos = map[string]diceHypothesis{
	sides4.hypo.Name:  sides4,
	sides6.hypo.Name:  sides6,
	sides8.hypo.Name:  sides8,
	sides12.hypo.Name: sides12,
	sides20.hypo.Name: sides20,
}

// an observation is the value rolled on the die
type diceObservation struct {
	val int16
}

// Getlikelihood is the likelihood function for the dice problem
func (o diceObservation) GetLikelihood(hypoName string) float64 {
	hypo, ok := diceHypos[hypoName]
	if !ok {
		return 0
	}
	if hypo.sides < o.val {
		// value of the roll is greater than the number of sides on hypothesis die; can't happen
		return 0
	}
	return 1 / float64(hypo.sides)
}

// Dice runs the dice problem
func Dice() {
	d := prob.NewSuite(sides4.hypo, sides6.hypo, sides8.hypo, sides12.hypo, sides20.hypo)

	obs := []diceObservation{
		diceObservation{6},
		diceObservation{6},
		diceObservation{8},
		diceObservation{7},
		diceObservation{7},
		diceObservation{5},
		diceObservation{4},
	}
	for _, ob := range obs {
		d.Update(ob)
	}

	d.Print()
}
