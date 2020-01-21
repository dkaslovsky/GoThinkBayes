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

var mmBag94 = mmBag{
	"brown":  .30,
	"yellow": .20,
	"red":    .20,
	"green":  .10,
	"orange": .10,
	"tan":    .10,
}
var mmBag96 = mmBag{
	"blue":   .24,
	"green":  .20,
	"orange": .16,
	"yellow": .14,
	"red":    .13,
	"brown":  .13,
}

type mmHypothesis struct {
	Hypo *prob.PmfElement
	Bags map[string]mmBag
}

// hypothesis A: bag1 is from 1994, bag 2 is from 1996
var hypoA = mmHypothesis{
	Hypo: prob.NewPmfElement("hypo A", 1),
	Bags: map[string]mmBag{
		"bag 1": mmBag94,
		"bag 2": mmBag96,
	},
}

// hypothesis B: bag1 is from 1996, bag 2 is from 1994
var hypoB = mmHypothesis{
	Hypo: prob.NewPmfElement("hypo B", 1),
	Bags: map[string]mmBag{
		"bag 1": mmBag96,
		"bag 2": mmBag94,
	},
}

// an observation corresponds to a color and the bag from which it was drawn
type mmObservation struct {
	Bag   string
	Color string
}

// Getlikelihood is the likelihood function for the M&M problem
func (o mmObservation) GetLikelihood(hypoName string) float64 {
	var hypo mmHypothesis

	switch hypoName {
	case hypoA.Hypo.Name:
		hypo = hypoA
	case hypoB.Hypo.Name:
		hypo = hypoB
	default:
		return 0
	}

	bag, ok := hypo.Bags[o.Bag]
	if !ok {
		return 0
	}

	like, ok := bag[o.Color]
	if !ok {
		return 0
	}
	return like
}

// MMs runs the M&M problem
func MMs() {
	m := prob.NewSuite(hypoA.Hypo, hypoB.Hypo)

	obs1 := mmObservation{Bag: "bag 1", Color: "yellow"}
	obs2 := mmObservation{Bag: "bag 2", Color: "green"}
	m.Update(obs1)
	m.Update(obs2)

	m.Print()
}
