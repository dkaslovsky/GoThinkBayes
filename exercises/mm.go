package exercises

import (
	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// M&M problem:
// In 1995 M&Ms changed the distribution of colors in a bag.
// Given a bag from 1994 and a bag from 1996, we draw one M&M from each bag: one yellow, one green.
// What is the probability the yellow one came from the 1994 bag?

// define m&m bags by their distribution of colors
type mmBag map[string]float64

var (
	mmBag94 = mmBag{
		"brown":  .30,
		"yellow": .20,
		"red":    .20,
		"green":  .10,
		"orange": .10,
		"tan":    .10,
	}
	mmBag96 = mmBag{
		"blue":   .24,
		"green":  .20,
		"orange": .16,
		"yellow": .14,
		"red":    .13,
		"brown":  .13,
	}
)

type mmHypothesis struct {
	hypo *prob.NamedPmfElement
	bags map[string]mmBag
}

var (
	// hypothesis A: bag1 is from 1994, bag 2 is from 1996
	hypoA = mmHypothesis{
		hypo: prob.NewNamedPmfElement("hypo A", 1),
		bags: map[string]mmBag{
			"bag 1": mmBag94,
			"bag 2": mmBag96,
		},
	}
	// hypothesis B: bag1 is from 1996, bag 2 is from 1994
	hypoB = mmHypothesis{
		hypo: prob.NewNamedPmfElement("hypo B", 1),
		bags: map[string]mmBag{
			"bag 1": mmBag96,
			"bag 2": mmBag94,
		},
	}
)

var mmHypos = map[string]mmHypothesis{
	hypoA.hypo.Name: hypoA,
	hypoB.hypo.Name: hypoB,
}

// an observation corresponds to a color and the bag from which it was drawn
type mmObservation struct {
	bag   string
	color string
}

// Getlikelihood is the likelihood function for the M&M problem
func (o *mmObservation) GetLikelihood(hypoName string) float64 {
	hypo, ok := mmHypos[hypoName]
	if !ok {
		return 0
	}

	bag, ok := hypo.bags[o.bag]
	if !ok {
		return 0
	}

	like, ok := bag[o.color]
	if !ok {
		return 0
	}
	return like
}

// MMs runs the M&M problem
func MMs() {
	s := prob.NewNamedSuite(hypoA.hypo, hypoB.hypo)

	obs := []prob.NamedSuiteObservation{
		&mmObservation{bag: "bag 1", color: "yellow"},
		&mmObservation{bag: "bag 2", color: "green"},
	}
	s.UpdateSet(obs)

	s.Print()
}
