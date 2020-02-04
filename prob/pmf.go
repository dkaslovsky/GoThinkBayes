package prob

import (
	"fmt"
	"math"
)

const (
	// float64EqualTol is the tolerance at which we consider float64s equal
	float64EqualTol = 1e-5
	// renormalize after this many consecutive updates
	renormalizeEvery = 10
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
func (p *Pmf) Mult(elem float64, multVal float64) {
	if _, ok := p.prob[elem]; !ok {
		// TODO: log a warning, print for now
		fmt.Printf("attempting to modify nonexisting element [%v]\n", elem)
		return
	}
	p.prob[elem] *= multVal
}

// Prob returns the probability associated with an element
func (p *Pmf) Prob(elem float64) float64 {
	val, ok := p.prob[elem]
	if !ok {
		return 0
	}
	return val
}

// Print prints the Pmf
func (p *Pmf) Print() {
	border := "----------"
	fmt.Println(border)
	for elem, prob := range p.prob {
		fmt.Printf("%v: %f\n", elem, prob)
	}
	fmt.Println(border)
	fmt.Println()
}

// Mean computes the mean of the Pmf
func (p *Pmf) Mean() float64 {
	total := 0.0
	for elem, prob := range p.prob {
		total += elem * prob
	}
	return total
}

// Percentile computes the specified percentile of the distribution
func (p *Pmf) Percentile(percentile float64) (float64, error) {
	if percentile < 0 || percentile > 1 {
		return 0, fmt.Errorf("percentile [%f] is outside of required range [0, 1]", percentile)
	}
	if len(p.prob) == 0 {
		return 0, fmt.Errorf("cannot compute percentile of empty Pmf")
	}
	// if !almostEqual(p.sum, 1.0, float64EqualTol) {
	// 	return 0, fmt.Errorf(
	// 		"cannot compute percentile of unnormalized Pmf (sum of elements [%f])",
	// 		p.sum,
	// 	)
	// }

	total := 0.0
	for _, elem := range sortKeys(p.prob) {
		total += p.prob[elem]
		if total >= percentile {
			return elem, nil
		}
	}

	return 0, fmt.Errorf("unable to compute percentile, potentially unnormalized Pmf")
}

// MakeCdf transforms a Pmf to a Cdf
func (p *Pmf) MakeCdf() (*Cdf, error) {
	return NewCdf(p.prob)
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
	for i, ob := range obs {
		for hypoName := range s.prob {
			like := ob.GetLikelihood(hypoName)
			s.Mult(hypoName, like)
		}
		// renormalize every few iterations for numerical stability
		if shouldRenormalize(i) {
			s.Normalize()
		}
	}
	s.Normalize()
}

func shouldRenormalize(iter int) bool {
	return math.Mod(float64(iter), renormalizeEvery) == 0
}

func almostEqual(val1, val2, tol float64) bool {
	return math.Abs(val1-val2) < tol
}
