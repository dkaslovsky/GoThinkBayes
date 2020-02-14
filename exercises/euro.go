package exercises

import (
	"fmt"

	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// Euro Problem:
// When spun on edge 250 times, a Belgian one-euro coin came up heads 140 times and tails 110.
// First, we esti-mate the probability that the coin lands face up.
// Second we evaluate whether the data support the hypothesis that the coin is biased.

type euroObservation struct {
	side string
}

// Getlikelihood is the likelihood function for the Euro problem
func (o *euroObservation) GetLikelihood(hypo float64) float64 {
	hProb := hypo / 100
	if o.side == "H" {
		return hProb
	}
	if o.side == "T" {
		return 1.0 - hProb
	}
	return 0
}

var (
	heads = &euroObservation{side: "H"}
	tails = &euroObservation{side: "T"}
)

// Euro runs the Euro problem
func Euro() {

	// a hypothesis represents that the probability of a heads is x%
	hypos := prob.Uniform(prob.NewBound(0, 100))
	s := prob.NewSuite(hypos...)

	nHeads := 140
	nTails := 110
	obs := []prob.SuiteObservation{}
	for i := 1; i <= nHeads; i++ {
		obs = append(obs, heads)
	}
	for i := 1; i <= nTails; i++ {
		obs = append(obs, tails)
	}

	s.MultiUpdate(obs)

	fmt.Printf("Posterior mean: %0.2f\n", s.Mean())
	median, err := s.Percentile(0.5)
	if err != nil {
		fmt.Printf("Unable to compute median due to error [%v]", err)
		return
	}
	fmt.Printf("Posterior median: %0.2f\n", median)
}
