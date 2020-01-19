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
	for elem := range p.prob {
		p.prob[elem] /= p.sum
	}
	p.sum = 1.0
}

// Prob ...
func (p *Pmf) Prob(elem string) (float64, bool) {
	val, ok := p.prob[elem]
	return val, ok
}

// Mult ...
func (p *Pmf) Mult(elem string, multVal float64) error {
	curVal, ok := p.Prob(elem)
	if !ok {
		return fmt.Errorf("%v does not exist", elem)
	}
	p.prob[elem] *= multVal
	p.sum += curVal * (multVal - 1) // maintain sum by subtracting curVal and adding curVal*multVal
	return nil
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
