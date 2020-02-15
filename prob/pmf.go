package prob

import (
	"fmt"
	"math/rand"
	"time"
)

// PmfElement is a discrete element in a NumericPmf
type PmfElement struct {
	Val  float64
	Prob float64
}

// NewPmfElement creates a new NumericPmfElement
func NewPmfElement(val float64, prob float64) *PmfElement {
	return &PmfElement{
		Val:  val,
		Prob: prob,
	}
}

// Pmf is a probability mass function
type Pmf struct {
	prob map[float64]float64
}

// NewPmf creates a new Pmf
func NewPmf() *Pmf {
	return &Pmf{
		prob: map[float64]float64{},
	}
}

// Set sets the value of an element
func (p *Pmf) Set(elem *PmfElement) {
	p.prob[elem.Val] = elem.Prob
}

// Normalize normalizes the values of the Pmf to sum to 1
func (p *Pmf) Normalize() {
	// recompute sum each time rather than maintain it for simplicity
	// and to match ThinkBayes implementation
	sum := 0.0
	for _, prob := range p.prob {
		sum += prob
	}

	if sum == 0 {
		return
	}

	for elem := range p.prob {
		p.prob[elem] /= sum
	}
}

// Mult multiplies the probability associated with an element by the specified value
func (p *Pmf) Mult(val float64, multFactor float64) {
	if _, ok := p.prob[val]; !ok {
		// TODO: log a warning, print for now
		fmt.Printf("attempting to modify nonexisting value [%v]\n", val)
		return
	}
	p.prob[val] *= multFactor
}

// Prob returns the probability associated with an element
func (p *Pmf) Prob(val float64) float64 {
	pr, ok := p.prob[val]
	if !ok {
		return 0
	}
	return pr
}

// Print prints the Pmf
func (p *Pmf) Print() {
	border := "----------"
	fmt.Println(border)
	for val, prob := range p.prob {
		fmt.Printf("%v: %f\n", val, prob)
	}
	fmt.Println(border)
	fmt.Println()
}

// Mean computes the mean of the Pmf
func (p *Pmf) Mean() (float64, error) {
	if len(p.prob) == 0 {
		return 0.0, fmt.Errorf("unable to compute mean of empty pmf")
	}

	total := 0.0
	for val, pr := range p.prob {
		total += val * pr
	}
	return total, nil
}

// Percentile computes the specified percentile of the distribution
func (p *Pmf) Percentile(percentile float64) (float64, error) {
	if percentile < 0 || percentile > 1 {
		return 0, fmt.Errorf("percentile [%f] is outside of required range [0, 1]", percentile)
	}
	if len(p.prob) == 0 {
		return 0, fmt.Errorf("cannot compute percentile of empty Pmf")
	}

	total := 0.0
	for _, val := range sortKeys(p.prob) {
		total += p.prob[val]
		if total >= percentile {
			return val, nil
		}
	}

	return 0, fmt.Errorf("unable to compute percentile, potentially unnormalized Pmf")
}

// MakeCdf transforms a Pmf to a Cdf
func (p *Pmf) MakeCdf() (*Cdf, error) {
	c, err := NewCdf(p.prob)
	return c, err
}

// SuiteObservation is the interface that must be satisfied to update probabilities
type SuiteObservation interface {
	GetLikelihood(float64) float64
}

// Suite is a suite of hypotheses with associated probabilities (a Pmf)
type Suite struct {
	*Pmf
}

// NewSuite creates a new Suite
func NewSuite(hypos ...*PmfElement) *Suite {
	s := &Suite{NewPmf()}
	for _, hypo := range hypos {
		s.Set(hypo)
	}
	s.Normalize()
	return s
}

// Update updates the probabilities based on an observation
func (s *Suite) Update(ob SuiteObservation) {
	for hypoName := range s.prob {
		like := ob.GetLikelihood(hypoName)
		s.Mult(hypoName, like)
	}
	s.Normalize()
}

// MultiUpdate updates the probabilities based on multiple observations
func (s *Suite) MultiUpdate(obs []SuiteObservation) {
	// iterate elements of obs in random order for numerical stability: avoids long runs
	// of one observation that push the probability of the others to values very close to zero
	rand.Seed(time.Now().UnixNano())
	for _, i := range rand.Perm(len(obs)) {
		ob := obs[i]
		for hypoName := range s.prob {
			like := ob.GetLikelihood(hypoName)
			s.Mult(hypoName, like)
		}
	}
	s.Normalize()
}
