package prob

import (
	"fmt"
)

// PmfElement is a discrete element in a PMF
type PmfElement struct {
	Name string
	Prob float64
}

// NewPmfElement creates a new PmfElement
func NewPmfElement(name string, prob float64) *PmfElement {
	return &PmfElement{
		Name: name,
		Prob: prob,
	}
}

// Pmf is a probability mass function
type Pmf struct {
	prob map[string]float64
	sum  float64
}

// NewPmf creates a new Pmf
func NewPmf() *Pmf {
	return &Pmf{
		prob: make(map[string]float64),
	}
}

// Set sets the value of an element
func (p *Pmf) Set(elem string, val float64) {
	p.prob[elem] = val
	p.sum += val
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
func (p *Pmf) Mult(elem string, multVal float64) {
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
func (p *Pmf) Prob(elem string) float64 {
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
		fmt.Printf("%s: %0.2f\n", elem, prob)
	}
	fmt.Println(border)
	fmt.Println()
}
