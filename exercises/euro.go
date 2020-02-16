package exercises

import (
	"fmt"
	"math"

	"github.com/dkaslovsky/GoThinkBayes/prob"
)

// Euro Problem:
// When spun on edge 250 times, a Belgian one-euro coin came up heads 140 times and tails 110.
// First, we esti-mate the probability that the coin lands face up.
// Second we evaluate whether the data support the hypothesis that the coin is biased.

type euroObservation struct {
	side string
}

// Getlikelihood is the likelihood function for the Euro problem using euroObservation
func (o *euroObservation) GetLikelihood(hypo float64) float64 {
	pHeads := hypo / 100
	if o.side == "H" {
		return pHeads
	}
	if o.side == "T" {
		return 1.0 - pHeads
	}
	return 0
}

func generateObs(nHeads, nTails int64) (obs []prob.SuiteObservation) {
	heads := &euroObservation{side: "H"}
	tails := &euroObservation{side: "T"}

	var i int64
	for i = 1; i <= nHeads; i++ {
		obs = append(obs, heads)
	}
	for i = 1; i <= nTails; i++ {
		obs = append(obs, tails)
	}
	return obs
}

// runEuro runs the Euro problem for a given set of hypotheses and observations,
// where a hypothesis represents that the probability of a heads is x%
func runEuro(hypos []*prob.PmfElement, obs []prob.SuiteObservation) {
	s := prob.NewSuite(hypos...)
	s.UpdateSet(obs)
	report(s)
}

type euroMultiObservation struct {
	nHeads int64
	nTails int64
}

// Getlikelihood is the likelihood function for the Euro problem using euroMultiObservation
func (o *euroMultiObservation) GetLikelihood(hypo float64) float64 {
	pHeads := hypo / 100
	return math.Pow(pHeads, float64(o.nHeads)) * math.Pow(1-pHeads, float64(o.nTails))
}

func runEuroMultiObservation(hypos []*prob.PmfElement, ob *euroMultiObservation) {
	s := prob.NewSuite(hypos...)
	s.Update(ob)
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

	var nHeads int64 = 140
	var nTails int64 = 110

	uniformPrior := prob.Uniform(prob.NewBound(0, 100))
	trianglePrior := prob.Triangle(prob.NewBound(0, 100))

	// run Euro problem using an observation for each flip
	obs := generateObs(nHeads, nTails)
	fmt.Println("Uniform Prior:")
	runEuro(uniformPrior, obs)
	fmt.Println("Triangle Prior:")
	runEuro(trianglePrior, obs)

	// run Euro problem using a multiobservationn to capture the results of multiple flips
	ob := &euroMultiObservation{nHeads: nHeads, nTails: nTails}
	fmt.Println("Uniform Prior (multiObservation):")
	runEuroMultiObservation(uniformPrior, ob)
	fmt.Println("Triangle Prior (multiObservation):")
	runEuroMultiObservation(trianglePrior, ob)

	// run Euro problem using a (continuous) Beta prior
	fmt.Println("Beta distribuion:")
	b, _ := prob.NewBeta(1, 1) // ignore error since we are passing positive parameters
	b.Update(float64(nHeads), float64(nTails))
	fmt.Printf("Posterior mean: %0.2f\n", b.Mean())
}
