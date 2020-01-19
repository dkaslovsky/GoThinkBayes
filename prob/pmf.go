package prob

import "fmt"

// Pmf ...
type Pmf struct {
	prob map[string]float64
	sum  float64
}

// NewPmf ...
func NewPmf() *Pmf {
	return &Pmf{
		prob: make(map[string]float64),
	}
}

// Set ...
func (p *Pmf) Set(elem string, val float64) {
	p.prob[elem] = val
	p.sum += val
}

// Normalize ...
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

// Prob ...
func (p *Pmf) Prob(elem string) float64 {
	val, ok := p.prob[elem]
	if !ok {
		return 0
	}
	return val
}

// Mult ...
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

// Print is for debug - DELETE ME!
func (p *Pmf) Print() {
	fmt.Println("----------")
	for elem, prob := range p.prob {
		fmt.Printf("%s: %0.2f\n", elem, prob)
	}
	fmt.Printf("(sum: %0.2f)\n", p.sum)
	fmt.Println()
}
