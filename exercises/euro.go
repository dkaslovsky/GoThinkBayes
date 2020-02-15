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

func generateObs(nHeads, nTails int) (obs []prob.SuiteObservation) {
	heads := &euroObservation{side: "H"}
	tails := &euroObservation{side: "T"}

	for i := 1; i <= nHeads; i++ {
		obs = append(obs, heads)
	}
	for i := 1; i <= nTails; i++ {
		obs = append(obs, tails)
	}
	return obs
}

// RunEuro runs the Euro problem for a given set of hypotheses and observations,
// where a hypothesis represents that the probability of a heads is x%
func RunEuro(hypos []*prob.PmfElement, obs []prob.SuiteObservation) {
	s := prob.NewSuite(hypos...)
	s.UpdateSet(obs)
	report(s)
}

func report(s *prob.Suite) {
	mle, err := s.MaximumLikelihood()
	if err != nil {
		fmt.Printf("Unable to compute maximum likelihood due to error [%v]", err)
		return
	}
	fmt.Printf("Posterior maximum likelihood estimate: %0.2f\n", mle)

	mean, err := s.Mean()
	if err != nil {
		fmt.Printf("Unable to compute mean due to error [%v]", err)
		return
	}
	fmt.Printf("Posterior mean: %0.2f\n", mean)

	median, err := s.Percentile(0.5)
	if err != nil {
		fmt.Printf("Unable to compute median due to error [%v]", err)
		return
	}
	fmt.Printf("Posterior median: %0.2f\n", median)

	ci := 90.0
	lower, upper, err := prob.CredibleInterval(s, ci)
	if err != nil {
		fmt.Printf("Unable to compute %0.2f%%-CI due to error [%v]", ci, err)
		return
	}
	fmt.Printf("%0.2f%%-CI: (%0.2f, %0.2f)\n", ci, lower, upper)
}

// Euro runs the Euro problem
func Euro() {

	uniformPrior := prob.Uniform(prob.NewBound(0, 100))
	trianglePrior := prob.Triangle(prob.NewBound(0, 100))

	nHeads, nTails := 140, 110
	obs := generateObs(nHeads, nTails)

	fmt.Println("Uniform Prior:")
	RunEuro(uniformPrior, obs)

	fmt.Println("Triangle Prior:")
	RunEuro(trianglePrior, obs)
}
