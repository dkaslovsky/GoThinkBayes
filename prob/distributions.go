package prob

import "math"

// TODO: consider noninteger bounds (floating point precision problem for keys underlying map...)

// Bound contains the bounds of a distribution
type Bound struct {
	Low  int
	High int
}

// NewBound constructs a new Bound
func NewBound(low, high int) *Bound {
	return &Bound{
		Low:  low,
		High: high,
	}
}

// Uniform generates PmfElements representing a uniform distribution
func Uniform(b *Bound) (elems []*PmfElement) {
	for i := b.Low; i <= b.High; i++ {
		elems = append(elems, NewPmfElement(float64(i), 1))
	}
	return elems
}

// PowerLaw generates PmfElements representing a power law distribution
func PowerLaw(b *Bound, alpha float64) (elems []*PmfElement) {
	a := -alpha
	for i := b.Low; i <= b.High; i++ {
		n := float64(i)
		elems = append(elems, NewPmfElement(n, math.Pow(n, a)))
	}
	return elems
}
