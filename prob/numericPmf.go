package prob

import (
	"fmt"
)

// NumericPmfElement is a discrete element in a NumericPmf
type NumericPmfElement struct {
	Val  float64
	Prob float64
}

// NewNumericPmfElement creates a new NumericPmfElement
func NewNumericPmfElement(val float64, prob float64) *NumericPmfElement {
	return &NumericPmfElement{
		Val:  val,
		Prob: prob,
	}
}

// NumericPmf is a probability mass function
type NumericPmf struct {
	prob map[float64]float64
	sum  float64
}

// NewNumericPmf creates a new Pmf
func NewNumericPmf() *NumericPmf {
	return &NumericPmf{
		prob: make(map[float64]float64),
	}
}

// Set sets the value of an element
func (p *NumericPmf) Set(elem *NumericPmfElement) {
	p.prob[elem.Val] = elem.Prob
	p.sum += elem.Prob
}

// Normalize normalizes the values of the Pmf to sum to 1
func (p *NumericPmf) Normalize() {
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
func (p *NumericPmf) Mult(elem float64, multVal float64) {
	curVal, ok := p.prob[elem]
	if !ok {
		// TODO: log a warning, print for now
		fmt.Printf("Attempting to modify nonexisting element [%v]\n", elem)
		return
	}
	p.prob[elem] *= multVal
	p.sum += curVal * (multVal - 1) // maintain sum by subtracting curVal and adding curVal*multVal
}

// Prob returns the probability associated with an element
func (p *NumericPmf) Prob(elem float64) float64 {
	val, ok := p.prob[elem]
	if !ok {
		return 0
	}
	return val
}

// Print prints the Pmf
func (p *NumericPmf) Print() {
	border := "----------"
	fmt.Println(border)
	for elem, prob := range p.prob {
		fmt.Printf("%v: %0.2f\n", elem, prob)
	}
	fmt.Println(border)
	fmt.Println()
}

type likelihoodGetterByNumeric interface {
	GetLikelihood(float64) float64
}

// NumericSuite is a suite of hypotheses with associated probabilities (a Pmf)
type NumericSuite struct {
	*NumericPmf
}

// NewNumericSuite creates a new Suite
func NewNumericSuite(hypos ...*NumericPmfElement) *NumericSuite {
	s := &NumericSuite{NewNumericPmf()}
	for _, hypo := range hypos {
		s.Set(hypo)
	}
	s.Normalize()
	return s
}

// Update updates the probabilities based on an observation
func (s *NumericSuite) Update(obs likelihoodGetterByNumeric) {
	for hypoName := range s.prob {
		like := obs.GetLikelihood(hypoName)
		s.Mult(hypoName, like)
	}
	s.Normalize()
}
