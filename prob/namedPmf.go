package prob

import (
	"fmt"
)

// NamedPmfElement is a discrete element in a NamedPmf
type NamedPmfElement struct {
	Name string
	Prob float64
}

// NewNamedPmfElement creates a new NamedPmfElement
func NewNamedPmfElement(name string, prob float64) *NamedPmfElement {
	return &NamedPmfElement{
		Name: name,
		Prob: prob,
	}
}

// NamedPmf is a probability mass function
type NamedPmf struct {
	prob map[string]float64
	sum  float64
}

// NewNamedPmf creates a new Pmf
func NewNamedPmf() *NamedPmf {
	return &NamedPmf{
		prob: make(map[string]float64),
	}
}

// Set sets the value of an element
func (p *NamedPmf) Set(elem *NamedPmfElement) {
	p.prob[elem.Name] = elem.Prob
	p.sum += elem.Prob
}

// Normalize normalizes the values of the Pmf to sum to 1
func (p *NamedPmf) Normalize() {
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
func (p *NamedPmf) Mult(elem string, multVal float64) {
	curVal, ok := p.prob[elem]
	if !ok {
		// TODO: log a warning, print for now
		fmt.Printf("Attempting to modify nonexisting element [%s]\n", elem)
		return
	}
	p.prob[elem] *= multVal
	p.sum += curVal * (multVal - 1) // maintain sum by subtracting curVal and adding curVal*multVal
}

// Prob returns the probability associated with an element
func (p *NamedPmf) Prob(elem string) float64 {
	val, ok := p.prob[elem]
	if !ok {
		return 0
	}
	return val
}

// Print prints the Pmf
func (p *NamedPmf) Print() {
	border := "----------"
	fmt.Println(border)
	for elem, prob := range p.prob {
		fmt.Printf("%s: %0.2f\n", elem, prob)
	}
	fmt.Println(border)
	fmt.Println()
}

// NamedSuiteObservation is the interface that must be satisfied to update probabilities
type NamedSuiteObservation interface {
	GetLikelihood(string) float64
}

// NamedSuite is a suite of hypotheses with associated probabilities (a Pmf)
type NamedSuite struct {
	*NamedPmf
}

// NewNamedSuite creates a new Suite
func NewNamedSuite(hypos ...*NamedPmfElement) *NamedSuite {
	s := &NamedSuite{NewNamedPmf()}
	for _, hypo := range hypos {
		s.Set(hypo)
	}
	s.Normalize()
	return s
}

// Update updates the probabilities based on an observation
func (s *NamedSuite) Update(ob NamedSuiteObservation) {
	for hypoName := range s.prob {
		like := ob.GetLikelihood(hypoName)
		s.Mult(hypoName, like)
	}
	s.Normalize()
}

// MultiUpdate updates the probabilities based on multiple observations
func (s *NamedSuite) MultiUpdate(obs []NamedSuiteObservation) {
	for _, ob := range obs {
		for hypoName := range s.prob {
			like := ob.GetLikelihood(hypoName)
			s.Mult(hypoName, like)
		}
	}
	s.Normalize()
}
