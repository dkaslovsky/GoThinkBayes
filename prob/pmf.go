package prob

import (
	"fmt"
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
	sum  float64
}

// NewPmf creates a new Pmf
func NewPmf() *Pmf {
	return &Pmf{
		prob: make(map[float64]float64),
	}
}

// Set sets the value of an element
func (p *Pmf) Set(elem *PmfElement) {
	p.prob[elem.Val] = elem.Prob
	p.sum += elem.Prob
}

// Normalize normalizes the values of the Pmf to sum to 1
func (p *Pmf) Normalize() {
	if p.sum == 0 {
		for elem := range p.prob {
			p.prob[elem] = 0
		}
		return
	}

	for elem := range p.prob {
		p.prob[elem] /= p.sum
	}
	p.sum = 1.0
}

// Mult multiplies the probability associated with an element by the specified value
func (p *Pmf) Mult(elem float64, multVal float64) {
	curVal, ok := p.prob[elem]
	if !ok {
		// TODO: log a warning, print for now
		fmt.Printf("attempting to modify nonexisting element [%v]\n", elem)
		return
	}
	p.prob[elem] *= multVal
	p.sum += curVal * (multVal - 1) // maintain sum by subtracting curVal and adding curVal*multVal
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
		fmt.Printf("%v: %0.2f\n", elem, prob)
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
func (p *Pmf) Percentile(percentile float64) float64 {
	elems := sortKeys(p.prob)

	if percentile < 0 {
		return elems[0]
	}
	if percentile > 1 {
		return elems[len(elems)-1]
	}

	total := 0.0
	for _, elem := range elems {
		total += p.prob[elem]
		if total >= percentile {
			return elem
		}
	}
	return elems[len(elems)-1]
}

// MakeCdf transforms a Pmf to a Cdf
func (p *Pmf) MakeCdf() *Cdf {
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
	for _, ob := range obs {
		for hypoName := range s.prob {
			like := ob.GetLikelihood(hypoName)
			s.Mult(hypoName, like)
		}
	}
	s.Normalize()
}
