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
	pmf       *Pmf
	nameToIdx map[string]float64
	nextIdx   float64
}

// NewNamedPmf creates a new Pmf
func NewNamedPmf() *NamedPmf {
	return &NamedPmf{
		pmf:       NewPmf(),
		nameToIdx: make(map[string]float64),
		nextIdx:   0.0,
	}
}

// Set sets the value of an element
func (p *NamedPmf) Set(elem *NamedPmfElement) {
	p.pmf.Set(NewPmfElement(p.nextIdx, elem.Prob))
	p.nameToIdx[elem.Name] = p.nextIdx
	p.nextIdx++
}

// Normalize normalizes the values of the Pmf to sum to 1
func (p *NamedPmf) Normalize() {
	p.pmf.Normalize()
}

// Mult multiplies the probability associated with an element by the specified value
func (p *NamedPmf) Mult(elem string, multVal float64) {
	idx, ok := p.nameToIdx[elem]
	if !ok {
		// TODO: log a warning, print for now
		fmt.Printf("Attempting to modify nonexisting element [%s]\n", elem)
		return
	}
	p.pmf.Mult(idx, multVal)
}

// Prob returns the probability associated with an element
func (p *NamedPmf) Prob(elem string) float64 {
	idx, ok := p.nameToIdx[elem]
	if !ok {
		return 0
	}
	return p.pmf.Prob(idx)
}

// Print prints the Pmf
func (p *NamedPmf) Print() {
	border := "----------"
	fmt.Println(border)
	for elem := range p.nameToIdx {
		prob := p.Prob(elem)
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
	for hypoName := range s.nameToIdx {
		like := ob.GetLikelihood(hypoName)
		s.Mult(hypoName, like)
	}
	s.Normalize()
}

// MultiUpdate updates the probabilities based on multiple observations
func (s *NamedSuite) MultiUpdate(obs []NamedSuiteObservation) {
	for _, ob := range obs {
		for hypoName := range s.nameToIdx {
			like := ob.GetLikelihood(hypoName)
			s.Mult(hypoName, like)
		}
	}
	s.Normalize()
}
