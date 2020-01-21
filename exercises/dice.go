package exercises

import (
	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// Dice Problem:
// We have 5 dice: 4-sided, 6-sided, 8-sided, 12-sided, and 20-sided
// Select a die at random, roll it, and record the number (potentially repeating the process).
// What is the probability that I rolled each die?

type diceHypothesis struct {
	Hypo  *prob.PmfElement
	Sides int16
}

var (
	sides4 = diceHypothesis{
		Hypo:  prob.NewPmfElement("4-sided", 1),
		Sides: 4,
	}
	sides6 = diceHypothesis{
		Hypo:  prob.NewPmfElement("6-sided", 1),
		Sides: 6,
	}
	sides8 = diceHypothesis{
		Hypo:  prob.NewPmfElement("8-sided", 1),
		Sides: 8,
	}
	sides12 = diceHypothesis{
		Hypo:  prob.NewPmfElement("12-sided", 1),
		Sides: 12,
	}
	sides20 = diceHypothesis{
		Hypo:  prob.NewPmfElement("20-sided", 1),
		Sides: 20,
	}
)

var diceHypos = map[string]diceHypothesis{
	sides4.Hypo.Name:  sides4,
	sides6.Hypo.Name:  sides6,
	sides8.Hypo.Name:  sides8,
	sides12.Hypo.Name: sides12,
	sides20.Hypo.Name: sides20,
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
	if hypo.Sides < o.val {
		// value of the roll is greater than the number of sides on hypothesis die; can't happen
		return 0
	}
	return 1 / float64(hypo.Sides)
}

// Dice runs the dice problem
func Dice() {
	d := prob.NewSuite(sides4.Hypo, sides6.Hypo, sides8.Hypo, sides12.Hypo, sides20.Hypo)

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
