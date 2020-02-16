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
		nameToIdx: map[string]float64{},
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
func (p *NamedPmf) Mult(name string, multFactor float64) {
	idx, ok := p.nameToIdx[name]
	if !ok {
		// TODO: log a warning, print for now
		fmt.Printf("Attempting to modify nonexisting element [%s]\n", name)
		return
	}
	p.pmf.Mult(idx, multFactor)
}

// Prob returns the probability associated with a name
func (p *NamedPmf) Prob(name string) float64 {
	idx, ok := p.nameToIdx[name]
	if !ok {
		return 0
	}
	return p.pmf.Prob(idx)
}

// Print prints the Pmf
func (p *NamedPmf) Print() {
	border := "----------"
	fmt.Println(border)
	for name := range p.nameToIdx {
		pr := p.Prob(name)
		fmt.Printf("%s: %0.2f\n", name, pr)
	}
	fmt.Println(border)
	fmt.Println()
}

// MaximumLikelihood returns the value with the highest probability
func (p *NamedPmf) MaximumLikelihood() (maxVal string, err error) {
	maxProb := 0.0
	for name, idx := range p.nameToIdx {
		prob := p.pmf.prob[idx]
		if prob > maxProb {
			maxProb = prob
			maxVal = name
		}
	}

	if maxProb == 0 {
		return maxVal, fmt.Errorf(
			"unable to compute maximum likelihood from empty pmf or all zero probabilities",
		)
	}
	return maxVal, nil
}
